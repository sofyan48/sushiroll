package rollout

import (
	"context"
	"log"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

func (i *cmdRollout) list(ctx context.Context) *cobra.Command {
	command := &cobra.Command{
		Use:   "list",
		Short: "Service list",
		Run: func(c *cobra.Command, args []string) {
			data, err := i.roll.GetList()
			if err != nil {
				log.Fatalf("Error get list: ", err)
			}

			headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
			columnFmt := color.New(color.FgYellow).SprintfFunc()

			tbl := table.New("No", "Name", "Revision", "Status", "Generation")
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
			for k, v := range data.Rollouts {
				k += 1
				tbl.AddRow(k, v.ObjectMeta.Name, v.ReplicaSet[0].Revision, v.Status, v.Generation)
			}
			tbl.Print()

		},
	}
	return command
}
