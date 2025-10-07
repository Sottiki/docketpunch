/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// punchCmd represents the punch command
var punchCmd = &cobra.Command{
	Use:   "punch [task ID]",
	Short: "done a task",
	Long: `punch a task by its ID
		For example:
		$ docket punch "#1"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("punch called")
	},
}

func init() {
	rootCmd.AddCommand(punchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// punchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// punchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
