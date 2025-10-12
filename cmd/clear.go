/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/Sottiki/docketpunch/internal/storage"
	"github.com/spf13/cobra"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "clear all dockets",
	Long: `Delete all completed tasks from the docket permanently.
		Incomplete tasks will remain.
		$ docket clear`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		docket, err := storage.Load()
		if err != nil {
			log.Fatalf("Failed to load data: %v\n", err)
		}

		deletedTasks := docket.ClearCompletedTasks()
		if len(deletedTasks) == 0 {
			fmt.Println("No completed tasks to clear.")
			return
		}
		if err := storage.Save(docket); err != nil {
			log.Fatalf("Failed to save data: %v", err)
		}
		fmt.Printf("✓ Cleared %d completed tasks:\n", len(deletedTasks))
		for _, t := range docket.Tasks {
			fmt.Println(formatTaskAsTicket(t))
		}
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clearCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clearCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
