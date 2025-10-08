package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"github.com/Sottiki/docketpunch/internal/docket"
	"github.com/Sottiki/docketpunch/internal/task"
)

// ~/.docket/tasks.json のパスを返す
func getDocketPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, ".docket", "tasks.json"), nil
}

// Docketディレクトリとファイルを初期化
func Init() error {
	path, err := getDocketPath()
	if err != nil {
		return err
	}

	// ディレクトリを作成（存在しない場合）
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// ファイルが存在しない場合、サンプルデータを作成
	if _, err := os.Stat(path); os.IsNotExist(err) {
		sampleData := createSampleData()
		return Save(sampleData)
	}

	return nil
}

// サンプル作成用
// TODO: 後で削除する
func createSampleData() *docket.Docket {
	d := docket.NewDocket()

	task1 := &task.Task{
		ID:          1,
		Description: "Go言語の勉強",
		Done:        false,
		CreatedAt:   time.Now().Add(-48 * time.Hour), // 2日前
		CompletedAt: nil,
	}

	task2 := &task.Task{
		ID:          2,
		Description: "ドキュメントを書く",
		Done:        false,
		CreatedAt:   time.Now().Add(-24 * time.Hour), // 1日前
		CompletedAt: nil,
	}

	completedTime := time.Now().Add(-12 * time.Hour) // 12時間前
	task3 := &task.Task{
		ID:          3,
		Description: "コードレビュー",
		Done:        true,
		CreatedAt:   time.Now().Add(-72 * time.Hour), // 3日前
		CompletedAt: &completedTime,
	}

	d.Tasks = []*task.Task{task1, task2, task3}
	d.NextID = 4

	return d
}

// タスクデータをファイルに保存する
func Save(d *docket.Docket) error {
	path, err := getDocketPath()
	if err != nil {
		return err
	}

	// JSONに変換（インデント付き）
	data, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return err
	}

	// ファイルに書き込み
	return os.WriteFile(path, data, 0644)
}

// ファイルからデータを読み込む
func Load() (*docket.Docket, error) {
	// 初期化を確認
	if err := Init(); err != nil {
		return nil, err
	}

	path, err := getDocketPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// JSONをパース
	var d docket.Docket
	if err := json.Unmarshal(data, &d); err != nil {
		return nil, err
	}

	return &d, nil
}
