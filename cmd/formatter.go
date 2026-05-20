/*
Copyright © 2025 Sottiki
*/
package cmd

import (
	"fmt"

	"github.com/Sottiki/docketpunch/internal/task"
	"github.com/fatih/color"
)

var (
	priorityHigh   = color.New(color.FgRed, color.Bold).SprintFunc()
	priorityMedium = color.New(color.FgHiGreen).SprintFunc()
	priorityLow    = color.New(color.FgCyan).SprintFunc()
)

func formatTaskAsTicket(t *task.Task) string {
	createDate := t.CreatedAt.Format("01/02")

	if t.Done {
		var dateInfo string
		if t.CompletedAt != nil {
			completeDate := t.CompletedAt.Format("01/02")
			dateInfo = fmt.Sprintf("(%s→%s)", createDate, completeDate)
		} else {
			dateInfo = fmt.Sprintf("(%s)", createDate)
		}
		var priorityTag string
		if t.Priority != "" {
			priorityTag = fmt.Sprintf("[%s] ", t.Priority)
		}
		return fmt.Sprintf("[ o|#%d|%s %s%s]", t.ID, t.Description, priorityTag, dateInfo)
	}

	var priorityTag string
	switch t.Priority {
	case "high":
		priorityTag = priorityHigh("[high]") + " "
	case "medium":
		priorityTag = priorityMedium("[medium]") + " "
	case "low":
		priorityTag = priorityLow("[low]") + " "
	}
	dateInfo := fmt.Sprintf("(%s)", createDate)
	return fmt.Sprintf("[  |#%d|%s %s%s]", t.ID, t.Description, priorityTag, dateInfo)
}
