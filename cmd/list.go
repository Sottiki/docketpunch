/*
Copyright © 2025 Sottiki
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/Sottiki/docketpunch/internal/storage"
	"github.com/Sottiki/docketpunch/internal/task"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long: `List tasks in ticket format.

Example:
  docket list
  docket list --priority high
  docket list --done
  docket list --pending`,
	Run: func(cmd *cobra.Command, args []string) {
		priorityFilter, _ := cmd.Flags().GetString("priority")
		doneOnly, _ := cmd.Flags().GetBool("done")
		pendingOnly, _ := cmd.Flags().GetBool("pending")
		defer func() {
			_ = cmd.Flags().Set("done", "false")
			_ = cmd.Flags().Set("pending", "false")
		}()

		if doneOnly && pendingOnly {
			fmt.Println("cannot use --done and --pending together")
			return
		}

		if priorityFilter != "" && priorityFilter != "high" && priorityFilter != "medium" && priorityFilter != "low" {
			fmt.Printf("Invalid priority: %s. Use: high, medium, low\n", priorityFilter)
			return
		}

		docket, err := storage.Load()
		if err != nil {
			log.Fatalf("Failed to load data: %v", err)
		}

		tasks := filterTasks(docket.Tasks, doneOnly, pendingOnly, priorityFilter)
		if len(tasks) == 0 {
			fmt.Println("No tasks found")
			return
		}

		for _, t := range tasks {
			fmt.Println(formatTaskAsTicket(t))
		}
	},
}

func filterTasks(tasks []*task.Task, doneOnly, pendingOnly bool, priority string) []*task.Task {
	filtered := make([]*task.Task, 0, len(tasks))
	for _, t := range tasks {
		if doneOnly && !t.Done {
			continue
		}
		if pendingOnly && t.Done {
			continue
		}
		if priority != "" && t.Priority != priority {
			continue
		}
		filtered = append(filtered, t)
	}
	return filtered
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().Bool("done", false, "完了したタスクのみ表示")
	listCmd.Flags().Bool("pending", false, "未完了のタスクのみ表示")
	listCmd.Flags().StringP("priority", "p", "", "優先度でフィルタ (high, medium, low)")
}
