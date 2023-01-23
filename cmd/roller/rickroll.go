/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package roller

import (
	"github.com/d1msk1y/go-rickroll/rickroll"
	"github.com/spf13/cobra"
)

// rickrollCmd represents the rickroll command
var rickrollCmd = &cobra.Command{
	Use:   "rickroll",
	Short: "This is palette that actually contains the rickroll commands",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		rickroll.RickRoll()
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rickrollCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rickrollCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
