package pods

import (
	"context"
	"log"
	"strings"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

func (i *cmdPods) list(ctx context.Context) *cobra.Command {
	command := &cobra.Command{
		Use:   "list",
		Short: "List pods",
		Run: func(c *cobra.Command, args []string) {
			namespace, _ := c.Flags().GetString("namespace")
			service, _ := c.Flags().GetString("service")
			data, err := i.roll.SetNamepace(namespace).Detail(service)
			if err != nil {
				log.Fatalf("Error get list: ", err)
			}
			headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
			columnFmt := color.New(color.FgYellow).SprintfFunc()
			tblAll := table.New("No", "Pods", "Status", "Ready", "Created", "Version")
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
				dataStructured = append(dataStructured, detailStructured)
			}
			for k, v := range dataStructured {
				k += 1
				tblAll.AddRow(k,
					v.Pod,
					v.Status,
					v.Ready,
					v.Created,
					v.Images,
				)
			}
			tblAll.Print()

		},
	}
	command.Flags().StringP("namespace", "n", "", "Your namespace")
	command.Flags().StringP("service", "s", "", "Your service")
	return command
}
