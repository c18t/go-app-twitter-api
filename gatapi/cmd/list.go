package cmd

import (
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Manage user list",
	Long: `Gatapi - Go Application for Twitter API. (for study)
Manage user list`,
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
