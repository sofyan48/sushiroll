package pods

import (
	"context"
	"log"
	"strings"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

func (i *cmdPods) container(ctx context.Context) *cobra.Command {
	command := &cobra.Command{
		Use:   "container",
		Short: "Pods container list",
		Run: func(c *cobra.Command, args []string) {
			namespace, _ := c.Flags().GetString("namespace")
			service, _ := c.Flags().GetString("service")
			data, err := i.roll.SetNamepace(namespace).Detail(service)
			if err != nil {
				log.Fatalf("Error get list: ", err)
			}
			headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
			columnFmt := color.New(color.FgYellow).SprintfFunc()
			tblAll := table.New("No", "Container", "Images", "version")
			tblAll.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
			for k, v := range data.Containers {
				k += 1
				images := strings.Split(v.Image, ":")
				tblAll.AddRow(k,
					v.Name,
					images[0],
					images[1],
				)
			}
			tblAll.Print()

		},
	}
	command.Flags().StringP("namespace", "n", "", "Your namespace")
	command.Flags().StringP("service", "s", "", "Your service")
	return command
}
