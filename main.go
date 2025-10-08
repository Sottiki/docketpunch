/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"log"

	"github.com/Sottiki/docketpunch/cmd"
	"github.com/Sottiki/docketpunch/internal/storage"
)

func main() {
	cmd.Execute()

	// TODO: 以下は動作確認用。後で削除する
	d, err := storage.Load()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("タスク数: %d\n", len(d.Tasks))
	fmt.Printf("Next ID: %d\n", d.NextID)
}
