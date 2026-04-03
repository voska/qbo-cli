package cmd

import (
	"encoding/json"
	"fmt"
	"mime"
	"os"
	"path/filepath"

	"github.com/voska/qbo-cli/internal/api"
	"github.com/voska/qbo-cli/internal/errfmt"
	"github.com/voska/qbo-cli/internal/output"
)

type AttachCmd struct {
	Args          []string `arg:"" optional:"" help:"[entity-type entity-id] file-path"`
	Note          string   `name:"note" help:"Note text to include with the attachment."`
	IncludeOnSend bool     `name:"include-on-send" help:"Include attachment when sending entity."`
}

const maxUploadSize = 100 * 1024 * 1024 // 100 MB

func (c *AttachCmd) Run(g *Globals) error {
	var entityType, entityID, filePath string

	nargs := len(c.Args)
	if nargs != 1 && nargs != 3 {
		return errfmt.Usage("usage: qbo attach [entity-type entity-id] <file>")
	}
	if nargs == 3 {
		entityType, entityID, filePath = c.Args[0], c.Args[1], c.Args[2]
	} else {
		filePath = c.Args[0]
	}

	var entityName string
	if entityType != "" {
		entity, ok := api.LookupEntity(entityType)
		if !ok {
			return errfmt.Usage("unknown entity type: " + entityType)
		}
		entityName = entity.Name
	}

	fi, err := os.Stat(filePath)
	if err != nil {
		return errfmt.Wrap(errfmt.ExitError, "cannot access file", err)
	}
	if fi.Size() > maxUploadSize {
		return errfmt.New(errfmt.ExitError, fmt.Sprintf("file too large (%d bytes, max %d)", fi.Size(), maxUploadSize))
	}

	fileName := fi.Name()
	contentType := mime.TypeByExtension(filepath.Ext(filePath))

	metadata := buildMetadata(fileName, contentType, c.Note, entityName, entityID, c.IncludeOnSend)
	metaJSON, err := json.Marshal(metadata)
	if err != nil {
		return errfmt.Wrap(errfmt.ExitError, "cannot marshal metadata", err)
	}

	if g.CLI.DryRun {
		output.Hint("[dry-run] POST /v3/company/{id}/upload")
		output.Hint("File: %s (%d bytes)", fileName, fi.Size())
		DryRunLog(g.CLI, "POST", "upload", metaJSON)
		return nil
	}

	client, _, err := g.NewAPIClient()
	if err != nil {
		return err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return errfmt.Wrap(errfmt.ExitError, "cannot open file", err)
	}
	defer func() { _ = file.Close() }()

	result, err := client.Upload(g.Ctx, metaJSON, fileName, contentType, file)
	if err != nil {
		return err
	}
	return WriteOutput(g.Ctx, result)
}

func buildMetadata(fileName, contentType, note, entityName, entityID string, includeOnSend bool) map[string]any {
	meta := map[string]any{
		"FileName": fileName,
	}
	if contentType != "" {
		meta["ContentType"] = contentType
	}
	if note != "" {
		meta["Note"] = note
	}
	if entityName != "" && entityID != "" {
		meta["AttachableRef"] = []map[string]any{
			{
				"EntityRef": map[string]any{
					"type":  entityName,
					"value": entityID,
				},
				"IncludeOnSend": includeOnSend,
			},
		}
	}
	return meta
}
