package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/Sottiki/docketpunch/internal/docket"
	"github.com/Sottiki/docketpunch/internal/task"
)

// テスト用に一時ディレクトリをホームディレクトリとして差し替える
func withTempHome(t *testing.T) func() {
	t.Helper()
	tmpDir := t.TempDir()
	original := os.Getenv("HOME")
	t.Setenv("HOME", tmpDir)
	return func() {
		os.Setenv("HOME", original)
	}
}

func TestSaveAndLoad(t *testing.T) {
	defer withTempHome(t)()

	d := docket.NewDocket()
	d.AddTask("タスクA", "")
	d.AddTask("タスクB", "")
	d.PunchTask(1)

	if err := Init(); err != nil {
		t.Fatalf("Init() error: %v", err)
	}
	if err := Save(d); err != nil {
		t.Fatalf("Save() error: %v", err)
	}

	loaded, err := Load()
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}

	if len(loaded.Tasks) != len(d.Tasks) {
		t.Errorf("Tasks length = %d, want %d", len(loaded.Tasks), len(d.Tasks))
	}

	if loaded.NextID != d.NextID {
		t.Errorf("NextID = %d, want %d", loaded.NextID, d.NextID)
	}

	for i, want := range d.Tasks {
		got := loaded.Tasks[i]
		if got.ID != want.ID {
			t.Errorf("Tasks[%d].ID = %d, want %d", i, got.ID, want.ID)
		}
		if got.Description != want.Description {
			t.Errorf("Tasks[%d].Description = %s, want %s", i, got.Description, want.Description)
		}
		if got.Done != want.Done {
			t.Errorf("Tasks[%d].Done = %v, want %v", i, got.Done, want.Done)
		}
	}
}

func TestInit_CreatesFileOnFirstRun(t *testing.T) {
	defer withTempHome(t)()

	if err := Init(); err != nil {
		t.Fatalf("Init() error: %v", err)
	}

	path, _ := getDocketPath()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Error("Init() should create tasks.json but it does not exist")
	}
}

func TestInit_Idempotent(t *testing.T) {
	defer withTempHome(t)()

	// 1回目: ファイル作成
	if err := Init(); err != nil {
		t.Fatalf("First Init() error: %v", err)
	}

	// ファイルを書き換えて中身を確認
	d := docket.NewDocket()
	d.AddTask("永続タスク", "")
	if err := Save(d); err != nil {
		t.Fatalf("Save() error: %v", err)
	}

	// 2回目: 既存ファイルを上書きしないことを確認
	if err := Init(); err != nil {
		t.Fatalf("Second Init() error: %v", err)
	}

	loaded, err := Load()
	if err != nil {
		t.Fatalf("Load() after second Init() error: %v", err)
	}

	if len(loaded.Tasks) != 1 || loaded.Tasks[0].Description != "永続タスク" {
		t.Error("Second Init() should not overwrite existing tasks.json")
	}
}

func TestLoad_InvalidJSON(t *testing.T) {
	defer withTempHome(t)()

	// 先にディレクトリとファイルを作成
	path, _ := getDocketPath()
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatalf("MkdirAll error: %v", err)
	}
	if err := os.WriteFile(path, []byte("不正なJSON{{{"), 0644); err != nil {
		t.Fatalf("WriteFile error: %v", err)
	}

	_, err := Load()
	if err == nil {
		t.Error("Load() with invalid JSON should return an error")
	}
}

func TestSave_PersistsCompletedAt(t *testing.T) {
	defer withTempHome(t)()

	d := docket.NewDocket()
	d.AddTask("完了タスク", "")
	d.PunchTask(1)

	if err := Init(); err != nil {
		t.Fatalf("Init() error: %v", err)
	}
	if err := Save(d); err != nil {
		t.Fatalf("Save() error: %v", err)
	}

	loaded, err := Load()
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}

	if loaded.Tasks[0].CompletedAt == nil {
		t.Error("CompletedAt should be persisted after Save/Load")
	}
}

func TestSave_JSONFormat(t *testing.T) {
	defer withTempHome(t)()

	d := &docket.Docket{
		Tasks:  []*task.Task{},
		NextID: 1,
	}

	if err := Init(); err != nil {
		t.Fatalf("Init() error: %v", err)
	}
	if err := Save(d); err != nil {
		t.Fatalf("Save() error: %v", err)
	}

	path, _ := getDocketPath()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile error: %v", err)
	}

	// 有効な JSON であることを確認
	var parsed map[string]interface{}
	if err := json.Unmarshal(data, &parsed); err != nil {
		t.Errorf("Saved file is not valid JSON: %v", err)
	}
}
