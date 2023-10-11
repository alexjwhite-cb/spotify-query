package cmd

import (
	"github.com/spf13/cobra"
)

var (
	artist string
)

func init() {
	rootCmd.AddCommand(artistCmd)

	artistCmd.Flags().StringVar(&artist, "artist", "", "")
}

var artistCmd = &cobra.Command{
	Use:   "artist",
	Short: "artist",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
