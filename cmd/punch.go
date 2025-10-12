/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Sottiki/docketpunch/internal/storage"
	"github.com/Sottiki/docketpunch/internal/task"
	"github.com/spf13/cobra"
)

// punchCmd represents the punch command
var punchCmd = &cobra.Command{
	Use:   "punch [task ID]",
	Short: "done a task",
	Long: `punch a task by its ID
		For example:
		$ docket punch "#1"
		or
		$ docket punch`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		docket, err := storage.Load()
		if err != nil {
			log.Fatalf("Failed to load data: %v", err)
		}

		var punchedTask *task.Task
		var success bool

		if len(args) == 0 {
			punchedTask = docket.GetLatestIncompleteTask()
			if punchedTask == nil {
				fmt.Println("No incomplete tasks found")
				return
			}
			_, success = docket.PunchTask(punchedTask.ID)
		} else {
			taskID, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatalf("Invalid task ID: %s", args[0])
			}
			punchedTask, success = docket.PunchTask(taskID)
		}

		if !success {
			if len(args) == 0 {
				fmt.Println("Failed to punch task")
			} else {
				fmt.Printf("Task #%s not found or already completed\n", args[0])
			}
			return
		}

		if err := storage.Save(docket); err != nil {
			log.Fatalf("Failed to save data: %v", err)
		}

		if len(args) == 0 {
			fmt.Printf("✓ Punched latest task #%d: %s\n", punchedTask.ID, punchedTask.Description)
		} else {
			fmt.Printf("✓ Punched task #%d: %s\n", punchedTask.ID, punchedTask.Description)
		}
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
