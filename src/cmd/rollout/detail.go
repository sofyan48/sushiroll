package rollout

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

type Detail struct {
	Name    string
	Pod     string
	Status  string
	Ready   string
	Created *time.Time
	Images  string
}

func (i *cmdRollout) detail(ctx context.Context) *cobra.Command {
	command := &cobra.Command{
		Use:   "detail",
		Short: "Service detail",
		Run: func(c *cobra.Command, args []string) {
			namespace, _ := c.Flags().GetString("namespace")
			service, _ := c.Flags().GetString("service")

			data, err := i.roll.SetNamepace(namespace).Detail(service)
			if err != nil {
				log.Fatalf("Error get list: ", err)
			}

			headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
			columnFmt := color.New(color.FgYellow).SprintfFunc()

			tbl := table.New("Name", "Pods", "Status", "Available", "Desired", "Current", "Ready", "Version")
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
			tbl.AddRow(
				data.ObjectMeta.Name,
				data.ReplicaSets[0].Pods[0].ObjectMeta.Name,
				data.ReplicaSets[0].Pods[0].Status,
				data.ReplicaSets[0].Available,
				data.Desired,
				data.Current,
				data.ReplicaSets[0].Pods[0].Ready,
				strings.Split(data.ReplicaSets[0].Images[0], ":")[1],
			)
			tbl.Print()
		},
	}
	command.Flags().StringP("namespace", "n", "", "Your namespace")
	command.Flags().StringP("service", "s", "", "Your service")
	return command
}
