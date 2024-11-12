package app

import (
	"context"

	"github.com/spf13/cobra"
)

type Command interface {
	Command(ctx context.Context) *cobra.Command
}
