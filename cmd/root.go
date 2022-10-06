package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Application environment variables
var (
	applicationAddress             string
	applicationHomeDirectoryPath   string
	applicationChromeHeadlessWSURL string
)

func init() {
	// Load up our `environment variables` from our operating system.
	appAddress := os.Getenv("HTMLTOIMAGE_SERVER_ADDRESS")
	if appAddress == "" {
		appAddress = "127.0.0.1:8001" // Set default IP address to localhost with port 8001.
	}
	appHomePath := os.Getenv("HTMLTOIMAGE_SERVER_HOME_DIRECTORY_PATH")
	if appHomePath == "" {
		path, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		appHomePath = path + "/db" // Set `db folder in the current location of the app.`
	}

	chWSURL := os.Getenv("HTMLTOIMAGE_SERVER_CHROME_HEADLESS_WS_URL")

	// Attach environment variables to system.
	rootCmd.PersistentFlags().StringVar(&applicationAddress, "appAddress", appAddress, "The applications address.")
	rootCmd.PersistentFlags().StringVar(&applicationHomeDirectoryPath, "appHomePath", appHomePath, "The path to the directory where this application saves the local files.")
	rootCmd.PersistentFlags().StringVar(&applicationChromeHeadlessWSURL, "chWSURL", chWSURL, "-")

	viper.BindPFlag("appAddress", rootCmd.PersistentFlags().Lookup("appAddress"))
	viper.BindPFlag("appHomePath", rootCmd.PersistentFlags().Lookup("appHomePath"))
	viper.BindPFlag("chWSURL", rootCmd.PersistentFlags().Lookup("chWSURL"))

	viper.SetDefault("appAddress", appAddress)
	viper.SetDefault("appHomePath", appHomePath)
	viper.SetDefault("chWSURL", chWSURL)
}

var rootCmd = &cobra.Command{
	Use:   "htmltoimage-server",
	Short: "RPC server over a single running HTML to IMAGE instance",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Do nothing.
	},
}

// Execute is the main entry into the application from the command line terminal.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
