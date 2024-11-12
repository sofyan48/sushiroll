package rollout

import (
	"context"

	"github.com/sofyan48/sushiroll/src/cmd/app"
	"github.com/spf13/cobra"
)

type cmdRollout struct {
}

func NewRolloutCommand() app.Command {
	return &cmdRollout{}
}

func (i *cmdRollout) Command(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rollout",
		Short: "Rollout model",
	}
	// cmd.AddCommand(i.packagesImporter(ctx))
	// cmd.AddCommand(i.memberImporter(ctx))
	// cmd.AddCommand(i.instantPackagesImporter(ctx))
	// cmd.AddCommand(i.gmvImporter(ctx))
	return cmd
}
