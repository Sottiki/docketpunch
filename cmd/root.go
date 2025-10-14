/*
Copyright © 2025 Sottiki
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const version = "0.1.0"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "docket",
	Short: "タスクにパンチ！",
	Long: `docketpunch - レトロなパンチカードをイメージしたタスク管理CLIツール
	
command & Enterで、タスクを完了するたびに、まるでチケットにパンチするように`,
	Run: func(cmd *cobra.Command, args []string) {
		showWelcome()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// --version フラグを追加
	rootCmd.Flags().BoolP("version", "v", false, "バージョンを表示")

	// --version が指定された場合の処理
	rootCmd.PreRun = func(cmd *cobra.Command, args []string) {
		v, _ := cmd.Flags().GetBool("version")
		if v {
			showWelcome()
			os.Exit(0)
		}
	}
}

// showWelcome はウェルカメッセージを表示
func showWelcome() {
	fmt.Print(`
  ╔═══════════════════╗
  ○   DOCKET PUNCH    ║
  ○ ───────────────── ║
  ○ Punch your tasks! ║
  ╚═══════════════════╝
       v` + version + `

`)
	fmt.Println("使い方: docket [command]")
	fmt.Println("コマンド一覧: docket --help")
}
