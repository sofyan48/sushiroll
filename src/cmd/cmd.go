package cmd

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"

	"github.com/sofyan48/sushiroll/src/cmd/actions"
	"github.com/sofyan48/sushiroll/src/cmd/auth"
	"github.com/sofyan48/sushiroll/src/cmd/rollout"
	"github.com/sofyan48/sushiroll/src/pkg/argo"
	"github.com/sofyan48/sushiroll/src/pkg/requester"
)

// Start handler registering service command
func Start() {

	rootCmd := &cobra.Command{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	url := os.Getenv("ARGO_ROLLOUT_URL")
	user := os.Getenv("ARGO_ROLLOUT_USERNAME")
	pass := os.Getenv("ARGO_ROLLOUT_PASSWORD")
	argoLib := argo.NewArgoRolloutLibrary(requester.New(), url, user, pass, true)
	cmd := []*cobra.Command{
		auth.NewAuthCommand().Command(ctx),
		rollout.NewRolloutCommand(argoLib).Command(ctx),
		actions.NewActionRolloutCommand().Command(ctx),
	}

	rootCmd.AddCommand(cmd...)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
