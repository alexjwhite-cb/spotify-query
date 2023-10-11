package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	clientId     string
	clientSecret string
)

func init() {
	rootCmd.PersistentFlags().StringVar(&clientId, "client-id", os.Getenv("SPOTIFY_CLIENT_ID"), "")
	if err := rootCmd.MarkPersistentFlagRequired("client-id"); err != nil {
		log.Fatal(err)
	}
	rootCmd.PersistentFlags().StringVar(&clientSecret, "client-secret", os.Getenv("SPOTIFY_CLIENT_SECRET"), "")
	if err := rootCmd.MarkPersistentFlagRequired("client-secret"); err != nil {
		log.Fatal(err)
	}
}

var rootCmd = &cobra.Command{
	Use:   "spotify",
	Short: "",
	Long:  ``,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
