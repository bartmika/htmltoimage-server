package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/spf13/cobra"

	"github.com/bartmika/htmltoimage-server/internal/config"
	"github.com/bartmika/htmltoimage-server/pkg/rpc"
)

func init() {
	rootCmd.AddCommand(sampleCmd)
}

var sampleCmd = &cobra.Command{
	Use:   "sample",
	Short: "Run a sample RPC on an active server.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Load up all the environment variables.
		appConf := config.AppConfig()

		// Connect to a running client.
		applicationAddress := fmt.Sprintf("%s:%s", appConf.Server.IP, appConf.Server.Port)
		client, err := rpc.NewClient(applicationAddress, 3, 15*time.Second)
		if err != nil {
			log.Fatal(err)
		}

		// Execute the remote call.
		imgBin, err := client.Screenshot("https://brank.as/")
		if err != nil {
			log.Fatal("Sample command failed generating image with error:", err)
		}

		// Save our file.
		if err := ioutil.WriteFile("data/screenshot.png", imgBin, 0o644); err != nil {
			log.Fatal("Sample command failed writing file with error:", err)
		}
		log.Println("Saved file: screenshot.png")
	},
}
