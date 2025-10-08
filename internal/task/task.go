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

// 新しいタスクを作成するコンストラクタ関数
func NewTask(id int, description string) *Task {
	return &Task{
		ID:          id,
		Description: description,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
}
