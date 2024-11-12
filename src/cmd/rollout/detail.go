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

			dataStructured := []Detail{}
			for _, i := range data.ReplicaSets {
				detailStructured := Detail{}
				for _, p := range i.Pods {
					detailStructured.Ready = p.Ready
					detailStructured.Pod = p.ObjectMeta.Name
					detailStructured.Status = p.Status
					detailStructured.Created = &p.ObjectMeta.CreationTimestamp
				}
				for _, c := range i.Images {
					detailStructured.Images = strings.Split(c, ":")[1]
				}
				detailStructured.Name = i.ObjectMeta.Name
				dataStructured = append(dataStructured, detailStructured)
			}

			headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
			columnFmt := color.New(color.FgYellow).SprintfFunc()

			tbl := table.New("No", "Name", "Pods", "Ready", "Status", "Created", "Version")
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
			for k, v := range dataStructured {
				k += 1
				tbl.AddRow(k,
					v.Name,
					v.Pod,
					v.Ready,
					v.Status,
					v.Created,
					v.Images,
				)
			}
			tbl.Print()
		},
	}
	command.Flags().StringP("namespace", "n", "", "Your namespace")
	command.Flags().StringP("service", "s", "", "Your service")
	return command
}
