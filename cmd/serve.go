package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/bartmika/htmltoimage-server/internal/app"
	"github.com/bartmika/htmltoimage-server/internal/config"
	"github.com/bartmika/htmltoimage-server/internal/inputports/rpc"
	"github.com/bartmika/htmltoimage-server/pkg/uuid"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Print the serve number",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serveing...")
		runServe()
	},
}

func runServe() {
	// Load up all the environment variables.
	appConf := config.AppConfig()

	uuidp := uuid.NewUUIDProvider()

	// Setup the application that handles our apps logic.
	a, err := app.New(appConf, uuidp)
	if err != nil {
		log.Fatal(err)
	}

	// Setup the RPC server to serve our app.
	srv := rpc.NewServer(appConf, a)

	// Run in the forground the RPC server. When the server gets termination
	// signal then the server will terminate.
	if err := srv.RunMainRuntimeLoop(); err != nil {
		log.Fatal(err)
	}
}
