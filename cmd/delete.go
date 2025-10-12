/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Sottiki/docketpunch/internal/storage"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [task ID]",
	Short: "Delete a docket from the list",
	Long: `Delete a docket from the list
		For example:
		$ docketpunch delete "1"`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("Invalid task ID: %s\n", args[0])
		}

		docket, err := storage.Load()
		if err != nil {
			log.Fatalf("Failed to load data: %v\n", err)
		}

		deletedTask, success := docket.DeleteTask(taskID)
		if !success {
			fmt.Printf("Task #%d not found\n", taskID)
			return
		}

		if err := storage.Save(docket); err != nil {
			log.Fatal("Failed to save data:", err)
		}
		fmt.Printf("✓ Deleted task #%d: %s\n", deletedTask.ID, deletedTask.Description)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
