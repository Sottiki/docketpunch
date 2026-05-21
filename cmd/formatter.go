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

func priorityTag(priority string) string {
	switch priority {
	case "high":
		return priorityHigh("[HIG]")
	case "medium":
		return priorityMedium("[MED]")
	case "low":
		return priorityLow("[LOW]")
	default:
		return "[ - ]"
	}
}

func formatTaskAsTicket(t *task.Task) string {
	createDate := t.CreatedAt.Format("01/02")
	tag := priorityTag(t.Priority)

	if t.Done {
		var dateInfo string
		if t.CompletedAt != nil {
			completeDate := t.CompletedAt.Format("01/02")
			dateInfo = fmt.Sprintf("(%s→%s)", createDate, completeDate)
		} else {
			dateInfo = fmt.Sprintf("(%s)", createDate)
		}
		return fmt.Sprintf("[ o |#%d|%s| %s %s]", t.ID, tag, t.Description, dateInfo)
	}

	dateInfo := fmt.Sprintf("(%s)", createDate)
	return fmt.Sprintf("[   |#%d|%s| %s %s]", t.ID, tag, t.Description, dateInfo)
}
