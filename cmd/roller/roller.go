/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package roller

import (
	"github.com/spf13/cobra"
)

// RollerCmd represents the roller command
var RollerCmd = &cobra.Command{
	Use:   "roller",
	Short: "Roller is a palette that contains roll based commands",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	RollerCmd.AddCommand(rickrollCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rollerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rollerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
