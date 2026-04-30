package docket

import (
	"testing"
)

func TestNewDocket(t *testing.T) {
	d := NewDocket()

	if d == nil {
		t.Fatal("NewDocket() returned nil")
	}

	if len(d.Tasks) != 0 {
		t.Errorf("NewDocket().Tasks length = %d, want 0", len(d.Tasks))
	}

	if d.NextID != 1 {
		t.Errorf("NewDocket().NextID = %d, want 1", d.NextID)
	}
}

func TestDocket_AddTask(t *testing.T) {
	d := NewDocket()
	description := "Test task"

	task := d.AddTask(description)

	// タスクが追加されたことを確認
	if len(d.Tasks) != 1 {
		t.Errorf("After AddTask, Tasks length = %d, want 1", len(d.Tasks))
	}

	// 正しいIDが割り当てられたことを確認
	if task.ID != 1 {
		t.Errorf("First task ID = %d, want 1", task.ID)
	}

	// NextIDがインクリメントされたことを確認
	if d.NextID != 2 {
		t.Errorf("After AddTask, NextID = %d, want 2", d.NextID)
	}

	// 説明が正しく設定されたことを確認
	if task.Description != description {
		t.Errorf("Task description = %s, want %s", task.Description, description)
	}

	// 2つ目のタスクを追加
	task2 := d.AddTask("Second task")
	if task2.ID != 2 {
		t.Errorf("Second task ID = %d, want 2", task2.ID)
	}

	if len(d.Tasks) != 2 {
		t.Errorf("After adding second task, Tasks length = %d, want 2", len(d.Tasks))
	}
}

func TestDocket_PunchTask(t *testing.T) {
	d := NewDocket()
	d.AddTask("Task 1")
	d.AddTask("Task 2")

	// タスクを完了させる
	punchedTask, ok := d.PunchTask(1)

	if !ok {
		t.Fatal("PunchTask(1) returned false, want true")
	}

	if punchedTask == nil {
		t.Fatal("PunchTask(1) returned nil task")
	}

	if !punchedTask.Done {
		t.Error("Punched task should be marked as done")
	}

	if punchedTask.CompletedAt == nil {
		t.Error("Punched task should have CompletedAt set")
	}

	// すでに完了したタスクを再度パンチ
	_, ok = d.PunchTask(1)
	if ok {
		t.Error("PunchTask on already done task should return false")
	}

	// 存在しないIDをパンチ
	_, ok = d.PunchTask(999)
	if ok {
		t.Error("PunchTask on non-existent ID should return false")
	}
}

func TestDocket_GetLatestIncompleteTask(t *testing.T) {
	d := NewDocket()

	// タスクがない場合
	task := d.GetLatestIncompleteTask()
	if task != nil {
		t.Error("GetLatestIncompleteTask() with no tasks should return nil")
	}

	// タスクを追加
	d.AddTask("Task 1")
	d.AddTask("Task 2")
	d.AddTask("Task 3")

	// 最新の未完了タスクを取得
	latest := d.GetLatestIncompleteTask()
	if latest == nil {
		t.Fatal("GetLatestIncompleteTask() returned nil")
	}

	if latest.ID != 3 {
		t.Errorf("Latest incomplete task ID = %d, want 3", latest.ID)
	}

	// Task 3を完了させる
	d.PunchTask(3)

	// 次の未完了タスクを取得
	latest = d.GetLatestIncompleteTask()
	if latest == nil {
		t.Fatal("GetLatestIncompleteTask() returned nil after punching task 3")
	}

	if latest.ID != 2 {
		t.Errorf("Latest incomplete task ID = %d, want 2", latest.ID)
	}

	// すべてのタスクを完了させる
	d.PunchTask(1)
	d.PunchTask(2)

	// 未完了タスクがない場合
	latest = d.GetLatestIncompleteTask()
	if latest != nil {
		t.Error("GetLatestIncompleteTask() with all tasks done should return nil")
	}
}

func TestDocket_DeleteTask(t *testing.T) {
	d := NewDocket()
	d.AddTask("Task 1")
	d.AddTask("Task 2")
	d.AddTask("Task 3")

	// タスクを削除
	deleted, ok := d.DeleteTask(2)

	if !ok {
		t.Fatal("DeleteTask(2) returned false, want true")
	}

	if deleted == nil {
		t.Fatal("DeleteTask(2) returned nil task")
	}

	if deleted.ID != 2 {
		t.Errorf("Deleted task ID = %d, want 2", deleted.ID)
	}

	// タスク数が減ったことを確認
	if len(d.Tasks) != 2 {
		t.Errorf("After DeleteTask, Tasks length = %d, want 2", len(d.Tasks))
	}

	// 削除されたタスクが存在しないことを確認
	for _, task := range d.Tasks {
		if task.ID == 2 {
			t.Error("Deleted task still exists in Tasks")
		}
	}

	// 存在しないIDを削除
	_, ok = d.DeleteTask(999)
	if ok {
		t.Error("DeleteTask on non-existent ID should return false")
	}
}

func TestDocket_ResetDocket(t *testing.T) {
	d := NewDocket()
	d.AddTask("Task 1")
	d.AddTask("Task 2")
	d.AddTask("Task 3")

	d.ResetDocket()

	if len(d.Tasks) != 0 {
		t.Errorf("After ResetDocket, Tasks length = %d, want 0", len(d.Tasks))
	}

	if d.NextID != 1 {
		t.Errorf("After ResetDocket, NextID = %d, want 1", d.NextID)
	}

	// リセット後に追加したタスクが ID=1 から再採番されることを確認
	newTask := d.AddTask("新タスク")
	if newTask.ID != 1 {
		t.Errorf("After ResetDocket, first new task ID = %d, want 1", newTask.ID)
	}
}

func TestDocket_ResetDocket_Empty(t *testing.T) {
	d := NewDocket()

	d.ResetDocket()

	if len(d.Tasks) != 0 {
		t.Errorf("ResetDocket on empty docket: Tasks length = %d, want 0", len(d.Tasks))
	}

	if d.NextID != 1 {
		t.Errorf("ResetDocket on empty docket: NextID = %d, want 1", d.NextID)
	}
}

func TestDocket_DeleteAllTasks(t *testing.T) {
	d := NewDocket()
	d.AddTask("Task 1")
	d.AddTask("Task 2")
	d.AddTask("Task 3")

	expectedNextID := d.NextID // 4

	d.DeleteAllTasks()

	if len(d.Tasks) != 0 {
		t.Errorf("After DeleteAllTasks, Tasks length = %d, want 0", len(d.Tasks))
	}

	if d.NextID != expectedNextID {
		t.Errorf("After DeleteAllTasks, NextID = %d, want %d (should not reset)", d.NextID, expectedNextID)
	}

	// 削除後に追加したタスクが連番継続されることを確認
	newTask := d.AddTask("継続タスク")
	if newTask.ID != expectedNextID {
		t.Errorf("After DeleteAllTasks, next task ID = %d, want %d", newTask.ID, expectedNextID)
	}
}

func TestDocket_DeleteAllTasks_Empty(t *testing.T) {
	d := NewDocket()

	d.DeleteAllTasks()

	if len(d.Tasks) != 0 {
		t.Errorf("DeleteAllTasks on empty docket: Tasks length = %d, want 0", len(d.Tasks))
	}
}
