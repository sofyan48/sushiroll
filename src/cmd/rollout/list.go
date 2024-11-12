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
			namespace, _ := c.Flags().GetString("namespace")
			data, err := i.roll.SetNamepace(namespace).GetList()
			if err != nil {
				log.Fatalf("Error get list: ", err)
			}

			headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
			columnFmt := color.New(color.FgYellow).SprintfFunc()

			tbl := table.New("No", "Name", "Revision", "Status", "Replica", "Generation", "Strategy")
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
			for k, v := range data.Rollouts {
				k += 1
				tbl.AddRow(k,
					v.ObjectMeta.Name,
					v.ReplicaSet[0].Revision,
					v.Status,
					v.ReplicaSet[0].Replicas,
					v.Generation,
					v.Strategy,
				)
			}
			tbl.Print()

		},
	}
	command.Flags().StringP("namespace", "n", "", "Your namespace")
	return command
}
