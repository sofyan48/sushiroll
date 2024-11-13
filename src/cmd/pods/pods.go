package pods

import (
	"context"
	"time"

	"github.com/sofyan48/sushiroll/src/cmd/contract"
	"github.com/sofyan48/sushiroll/src/pkg/argo"
	"github.com/spf13/cobra"
)

type cmdPods struct {
	roll argo.ArgoRolloutLibrary
}

func NewPodRolloutCommand(roll argo.ArgoRolloutLibrary) contract.Command {
	return &cmdPods{
		roll: roll,
	}
}

type Detail struct {
	Pod     string
	Status  string
	Ready   string
	Created *time.Time
	Images  string
}

func (i *cmdPods) Command(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pods",
		Short: "Manage your pods",
	}
	cmd.AddCommand(i.list(ctx))
	cmd.AddCommand(i.container(ctx))
	return cmd
}
