package docket

import (
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
