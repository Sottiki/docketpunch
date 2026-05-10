/*
Copyright © 2025 Sottiki
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Sottiki/docketpunch/internal/storage"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit <id> <description>",
	Short: "タスクの説明文を編集する",
	Long: `指定したIDのタスクの説明文を編集します。
$ docket edit 1 新しい説明文`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("Invalid task ID: %s", args[0])
		}

		newDescription := strings.Join(args[1:], " ")

		docket, err := storage.Load()
		if err != nil {
			log.Fatalf("Failed to load data: %v", err)
		}

		editedTask, ok := docket.EditTask(taskID, newDescription)
		if !ok {
			fmt.Printf("Task #%d not found\n", taskID)
			return
		}

		if err := storage.Save(docket); err != nil {
			log.Fatalf("Failed to save data: %v", err)
		}

		fmt.Printf("Edited task #%d: %s\n", editedTask.ID, editedTask.Description)
		for _, t := range docket.Tasks {
			fmt.Println(formatTaskAsTicket(t))
		}
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
