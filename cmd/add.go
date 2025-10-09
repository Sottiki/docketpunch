/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
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

		docket, err := storage.Load()
		if err != nil {
			log.Fatalf("Failed to load data: %v", err)
		}

		newTask := docket.AddTask(description)

		if err := storage.Save(docket); err != nil {
			log.Fatalf("Error saving storage:", err)
		}

		fmt.Printf("Added task: #%d : %s\n", newTask.ID, newTask.Description)
		for _, t := range docket.Tasks {
			fmt.Println(formatTaskAsTicket(t))
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
