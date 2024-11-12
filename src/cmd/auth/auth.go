package auth

import (
	"context"

	"github.com/sofyan48/sushiroll/src/cmd/contract"
	"github.com/spf13/cobra"
)

type cmdAuth struct {
}

func NewAuthCommand() contract.Command {
	return &cmdAuth{}
}

func (i *cmdAuth) Command(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Setup authentication",
	}
	cmd.AddCommand(i.configSet(ctx))
	return cmd
}
