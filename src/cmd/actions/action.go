package actions

import (
	"context"

	"github.com/sofyan48/sushiroll/src/cmd/contract"
	"github.com/spf13/cobra"
)

type cmdRolloutAction struct {
}

func NewActionRolloutCommand() contract.Command {
	return &cmdRolloutAction{}
}

func (i *cmdRolloutAction) Command(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "action",
		Short: "All action to setup service",
	}
	cmd.AddCommand(i.promote(ctx))
	return cmd
}
