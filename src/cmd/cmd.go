package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"

	"github.com/sofyan48/sushiroll/src/cmd/rollout"
)

// Start handler registering service command
func Start() {

	rootCmd := &cobra.Command{}
	ctx, cancel := context.WithCancel(context.Background())

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-quit
		cancel()
	}()

	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	cmd := []*cobra.Command{
		rollout.NewRolloutCommand().Command(ctx),
	}

	rootCmd.AddCommand(cmd...)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
