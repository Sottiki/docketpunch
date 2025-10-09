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

		for _, t := range docket.Tasks {
			fmt.Println(formatTaskAsTicket(t))
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

// タスクをチケット形式で表示する
func formatTaskAsTicket(t *task.Task) string {
	statusMark := " "
	if t.Done {
		statusMark = "◯"
	}

	createDate := t.CreatedAt.Format("01/02")

	var dateInfo string
	if t.Done && t.CompletedAt != nil {
		completeDate := t.CompletedAt.Format("01/02")
		dateInfo = fmt.Sprintf("(%s→%s)", createDate, completeDate)
	} else {
		dateInfo = fmt.Sprintf("(%s→)", createDate)
	}
	return fmt.Sprintf("[ %s|#%d|%s %s]", statusMark, t.ID, t.Description, dateInfo)

}
