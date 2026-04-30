/*
Copyright © 2025 Sottiki
*/
package cmd

import (
	"fmt"

	"github.com/Sottiki/docketpunch/internal/task"
)

func formatTaskAsTicket(t *task.Task) string {
	statusMark := " "
	if t.Done {
		statusMark = "o"
	}

	createDate := t.CreatedAt.Format("01/02")

	var dateInfo string
	if t.Done && t.CompletedAt != nil {
		completeDate := t.CompletedAt.Format("01/02")
		dateInfo = fmt.Sprintf("(%s→%s)", createDate, completeDate)
	} else {
		dateInfo = fmt.Sprintf("(%s)", createDate)
	}
	return fmt.Sprintf("[ %s|#%d|%s %s]", statusMark, t.ID, t.Description, dateInfo)
}
