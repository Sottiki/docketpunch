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

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "全タスクを削除して連番をリセットする",
	Long: `全タスクを削除し、IDの連番を1からリセットします。
		$ docket clear`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		docket, err := storage.Load()
		if err != nil {
			log.Fatalf("Failed to load data: %v\n", err)
		}

		docket.ResetDocket()

		if err := storage.Save(docket); err != nil {
			log.Fatalf("Failed to save data: %v", err)
		}
		fmt.Println("✓ Cleared all tasks and reset ID counter.")
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
