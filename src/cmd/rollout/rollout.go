package rollout

import (
	"context"

	"github.com/sofyan48/sushiroll/src/cmd/contract"
	"github.com/sofyan48/sushiroll/src/pkg/argo"
	"github.com/spf13/cobra"
)

type cmdRollout struct {
	roll argo.ArgoRolloutLibrary
}

func NewRolloutCommand(roll argo.ArgoRolloutLibrary) contract.Command {
	return &cmdRollout{
		roll: roll,
	}
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
