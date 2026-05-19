/*
Copyright © 2025 Sottiki
*/
package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/Sottiki/docketpunch/internal/storage"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [docket description]",
	Short: "add new docket to the list",
	Long: `add new docket to the list
		For example:
		$ docketpunch add "My new docket"`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		description := strings.Join(args, " ")
		priority, _ := cmd.Flags().GetString("priority")

		if priority != "" && priority != "high" && priority != "medium" && priority != "low" {
			fmt.Printf("Invalid priority: %s. Use: high, medium, low\n", priority)
			return
		}

		docket, err := storage.Load()
		if err != nil {
			log.Fatalf("Failed to load data: %v", err)
		}

		newTask := docket.AddTask(description, priority)

		if err := storage.Save(docket); err != nil {
			log.Fatalf("Error saving storage: %v", err)
		}

		fmt.Printf("Added task: #%d : %s\n", newTask.ID, newTask.Description)
		for _, t := range docket.Tasks {
			fmt.Println(formatTaskAsTicket(t))
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("priority", "p", "", "優先度を設定 (high, medium, low)")
}
