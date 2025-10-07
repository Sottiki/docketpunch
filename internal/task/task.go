package task

import "time"

// 個別のタスク項目を表す構造体
type Task struct {
	ID          int        `json:"id"`
	Description string     `json:"description"`
	Done        bool       `json:"done"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"` // 🆕 完了時刻
}

// タスクリスト全体を管理する構造体
type Docket struct {
	Tasks []*Task `json:"tasks"`
}
