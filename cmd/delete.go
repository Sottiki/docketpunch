/*
Copyright © 2025 Sottiki
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Sottiki/docketpunch/internal/storage"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [task ID]",
	Short: "指定したタスクを削除する",
	Long: `指定したIDのタスクを削除します。
--all フラグを指定すると全タスクを削除します（連番はリセットしません）。
$ docket delete 1
$ docket delete --all`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		all, _ := cmd.Flags().GetBool("all")

		docket, err := storage.Load()
		if err != nil {
			log.Fatalf("Failed to load data: %v\n", err)
		}

		if all {
			count := len(docket.Tasks)
			if count == 0 {
				fmt.Println("No tasks to delete.")
				return
			}
			docket.DeleteAllTasks()
			if err := storage.Save(docket); err != nil {
				log.Fatal("Failed to save data:", err)
			}
			fmt.Printf("✓ Deleted all %d tasks.\n", count)
			return
		}

		if len(args) == 0 {
			fmt.Println("task ID を指定してください。全削除は --all を使用してください。")
			return
		}

		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("Invalid task ID: %s\n", args[0])
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
	deleteCmd.Flags().Bool("all", false, "全タスクを削除する（連番はリセットしない）")
}
