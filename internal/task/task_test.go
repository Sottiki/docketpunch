package task

import (
	"testing"
	"time"
)

func TestNewTask(t *testing.T) {
	// テストケース
	id := 1
	description := "Test task"

	// 実行
	task := NewTask(id, description)

	// 検証
	if task.ID != id {
		t.Errorf("NewTask().ID = %d, want %d", task.ID, id)
	}

	if task.Description != description {
		t.Errorf("NewTask().Description = %s, want %s", task.Description, description)
	}

	if task.Done != false {
		t.Errorf("NewTask().Done = %v, want false", task.Done)
	}

	if task.CompletedAt != nil {
		t.Errorf("NewTask().CompletedAt = %v, want nil", task.CompletedAt)
	}

	// CreatedAtが現在時刻に近いことを確認
	now := time.Now()
	diff := now.Sub(task.CreatedAt)
	if diff > time.Second {
		t.Errorf("NewTask().CreatedAt is too old: %v", diff)
	}
}

func TestTask_InitialState(t *testing.T) {
	task := NewTask(1, "Initial task")

	if task.Done {
		t.Error("New task should not be done")
	}

	if task.CompletedAt != nil {
		t.Error("New task should not have CompletedAt set")
	}
}
