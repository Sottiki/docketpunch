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
		switch t.Priority {
		case "high":
			priorityTag = "[HIG] "
		case "medium":
			priorityTag = "[MED] "
		case "low":
			priorityTag = "[LOW] "
		}
		return fmt.Sprintf("[ o |#%d| %s %s%s]", t.ID, t.Description, priorityTag, dateInfo)
	}

	var priorityTag string
	switch t.Priority {
	case "high":
		priorityTag = priorityHigh("[HIG]") + " "
	case "medium":
		priorityTag = priorityMedium("[MED]") + " "
	case "low":
		priorityTag = priorityLow("[LOW]") + " "
	}
	dateInfo := fmt.Sprintf("(%s)", createDate)
	return fmt.Sprintf("[   |#%d| %s %s%s]", t.ID, t.Description, priorityTag, dateInfo)
}
