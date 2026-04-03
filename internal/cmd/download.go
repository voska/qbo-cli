package cmd

import (
	"io"
	"os"
	"path/filepath"

	"github.com/voska/qbo-cli/internal/errfmt"
	"github.com/voska/qbo-cli/internal/output"
)

type DownloadCmd struct {
	ID     string `arg:"" help:"Attachable ID to download."`
	Output string `name:"output" short:"o" help:"Save to this file path instead of original filename."`
	URL    bool   `name:"url" help:"Print the temporary download URL instead of saving the file."`
}

func (c *DownloadCmd) Run(g *Globals) error {
	if g.CLI.DryRun {
		output.Hint("[dry-run] GET /v3/company/{id}/download/%s", c.ID)
		return nil
	}

	client, _, err := g.NewAPIClient()
	if err != nil {
		return err
	}

	if c.URL {
		url, err := client.FetchDownloadURL(g.Ctx, c.ID)
		if err != nil {
			return err
		}
		return WriteOutput(g.Ctx, map[string]any{"url": url})
	}

	savePath := c.Output
	if savePath == "" {
		meta, err := client.Read(g.Ctx, "attachable", c.ID)
		if err != nil {
			return err
		}
		savePath = extractFileName(meta)
	}

	body, err := client.Download(g.Ctx, c.ID)
	if err != nil {
		return err
	}
	defer func() { _ = body.Close() }()

	f, err := os.Create(savePath)
	if err != nil {
		return errfmt.Wrap(errfmt.ExitError, "cannot create file", err)
	}

	_, err = io.Copy(f, body)
	_ = f.Close()
	if err != nil {
		_ = os.Remove(savePath)
		return errfmt.Wrap(errfmt.ExitError, "cannot write file", err)
	}

	output.Hint("saved %s", savePath)
	return nil
}

func extractFileName(meta map[string]any) string {
	att, _ := meta["Attachable"].(map[string]any)
	name, _ := att["FileName"].(string)
	return filepath.Base(name)
}
