package docket

import (
	"time"

	"github.com/Sottiki/docketpunch/internal/task"
)

// タスクの集合体を表す構造体
type Docket struct {
	Tasks  []*task.Task `json:"tasks"`
	NextID int          `json:"next_id"`
}

// 新しいDocketを初期化する関数
func NewDocket() *Docket {
	return &Docket{
		Tasks:  []*task.Task{},
		NextID: 1,
	}
}

func (d *Docket) AddTask(description string) *task.Task {
	newTask := task.NewTask(d.NextID, description)
	d.Tasks = append(d.Tasks, newTask)
	d.NextID++
	return newTask
}

func (d *Docket) PunchTask(id int) (*task.Task, bool) {
	for _, t := range d.Tasks {
		if t.ID == id && !t.Done {
			t.Done = true
			now := time.Now()
			t.CompletedAt = &now
			return t, true
		}
	}
	return nil, false
}

func (d *Docket) GetLatestIncompleteTask() *task.Task {
	for i := len(d.Tasks) - 1; i >= 0; i-- {
		if !d.Tasks[i].Done {
			return d.Tasks[i]
		}
	}
	return nil
}

func (d *Docket) DeleteTask(id int) (*task.Task, bool) {
	for i, t := range d.Tasks {
		if t.ID == id {
			// タスクを削除
			deletedTask := t
			d.Tasks = append(d.Tasks[:i], d.Tasks[i+1:]...)
			return deletedTask, true
		}
	}
	return nil, false
}

func (d *Docket) ClearCompletedTasks() []*task.Task {
	var deletedTasks []*task.Task
	var remainingTasks []*task.Task

	for _, t := range d.Tasks {
		if t.Done {
			deletedTasks = append(deletedTasks, t)
		} else {
			remainingTasks = append(remainingTasks, t)
		}
	}

	d.Tasks = remainingTasks
	return deletedTasks
}
