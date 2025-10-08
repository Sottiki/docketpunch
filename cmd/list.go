/*
Copyright © 2025 Sottiki
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/Sottiki/docketpunch/internal/storage"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long: `List all tasks in ticket format.
	
Example:
  docket list`,
	Run: func(cmd *cobra.Command, args []string) {
		docket, err := storage.Load()
		if err != nil {
			log.Fatalf("Failed to load data: %v", err)
		}

		if len(docket.Tasks) == 0 {
			fmt.Println("No tasks found")
			return
		}

		// TODOタスクを表示する
		fmt.Printf("タスク数: %d\n", len(docket.Tasks))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
