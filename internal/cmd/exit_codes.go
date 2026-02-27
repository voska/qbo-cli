package cmd

import (
	"github.com/voska/qbo-cli/internal/errfmt"
)

type ExitCodesCmd struct{}

func (c *ExitCodesCmd) Run(g *Globals) error {
	return WriteOutput(g.Ctx, errfmt.ExitCodeTable())
}
