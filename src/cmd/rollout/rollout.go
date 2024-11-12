package rollout

import (
	"context"

	"github.com/sofyan48/sushiroll/src/cmd/contract"
	"github.com/spf13/cobra"
)

type cmdRollout struct {
}

func NewRolloutCommand() contract.Command {
	return &cmdRollout{}
}

func (i *cmdRollout) Command(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rollout",
		Short: "Rollout model",
	}
	cmd.AddCommand(i.list(ctx))
	cmd.AddCommand(i.detail(ctx))
	return cmd
}
