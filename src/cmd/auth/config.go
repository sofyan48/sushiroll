package auth

import (
	"context"

	"github.com/spf13/cobra"
)

func (i *cmdAuth) configSet(ctx context.Context) *cobra.Command {
	command := &cobra.Command{
		Use:   "set",
		Short: "Configuration setup",
		Run: func(c *cobra.Command, args []string) {
			// start, _ := c.Flags().GetString("start")
			// end, _ := c.Flags().GetString("end")

		},
	}
	command.Flags().StringP("start", "s", "", "Add start date")
	command.Flags().StringP("end", "e", "", "Add end date")
	return command
}
