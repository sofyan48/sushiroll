package rollout

import (
	"context"
	"fmt"
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
	var allOption bool
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
			if allOption {
				fmt.Println("\nShow All resource")
				tblAll := table.New("No", "Name", "Pods", "Status", "Ready", "Created", "Version")
				tblAll.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
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
				for k, v := range dataStructured {
					k += 1
					tblAll.AddRow(k,
						v.Name,
						v.Pod,
						v.Status,
						v.Ready,
						v.Created,
						v.Images,
					)
				}
				tblAll.Print()
			}
		},
	}
	command.Flags().StringP("namespace", "n", "", "Your namespace")
	command.Flags().StringP("service", "s", "", "Your service")
	command.Flags().BoolVarP(&allOption, "all", "a", false, "Show all items")
	return command
}
