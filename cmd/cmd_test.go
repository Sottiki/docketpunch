package cmd

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

// captureOutput は os.Stdout を pipe に差し替えて f() の出力を文字列で返す
func captureOutput(f func()) string {
	r, w, _ := os.Pipe()
	origStdout := os.Stdout
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = origStdout

	var buf bytes.Buffer
	buf.ReadFrom(r)
	return buf.String()
}

// withTempHome はテスト中 HOME を一時ディレクトリに差し替え、終了後に戻す
func withTempHome(t *testing.T) {
	t.Helper()
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)
}

// run はコマンド引数を渡して rootCmd を実行し、標準出力を返す
func run(args ...string) string {
	return captureOutput(func() {
		rootCmd.SetArgs(args)
		rootCmd.Execute() //nolint:errcheck
	})
}

// --- add ---

func TestAddCmd_AddsTask(t *testing.T) {
	withTempHome(t)

	out := run("add", "テストタスク")

	if !strings.Contains(out, "Added task") {
		t.Errorf("add output should contain 'Added task', got: %s", out)
	}
	if !strings.Contains(out, "テストタスク") {
		t.Errorf("add output should contain task description, got: %s", out)
	}
}

func TestAddCmd_ShowsListAfterAdd(t *testing.T) {
	withTempHome(t)

	out := run("add", "タスクA")

	if !strings.Contains(out, "#1") {
		t.Errorf("add output should show ticket list with #1, got: %s", out)
	}
}

func TestAddCmd_MultipleTasksIncrementID(t *testing.T) {
	withTempHome(t)

	run("add", "タスクA")
	out := run("add", "タスクB")

	if !strings.Contains(out, "#2") {
		t.Errorf("second task should have ID #2, got: %s", out)
	}
}

// --- list ---

func TestListCmd_EmptyShowsMessage(t *testing.T) {
	withTempHome(t)

	out := run("list")

	if !strings.Contains(out, "No tasks found") {
		t.Errorf("list on empty docket should show 'No tasks found', got: %s", out)
	}
}

func TestListCmd_ShowsTicketFormat(t *testing.T) {
	withTempHome(t)

	run("add", "テストタスク")
	out := run("list")

	if !strings.Contains(out, "[") || !strings.Contains(out, "#1") {
		t.Errorf("list should show ticket format, got: %s", out)
	}
}

func TestListCmd_DoneFlag_ShowsOnlyCompleted(t *testing.T) {
	withTempHome(t)

	run("add", "未完了タスク")
	run("add", "完了タスク")
	run("punch", "2")

	out := run("list", "--done")

	if !strings.Contains(out, "#2") {
		t.Errorf("--done should show completed task #2, got: %s", out)
	}
	if strings.Contains(out, "#1") {
		t.Errorf("--done should not show pending task #1, got: %s", out)
	}
}

func TestListCmd_PendingFlag_ShowsOnlyPending(t *testing.T) {
	withTempHome(t)

	run("add", "未完了タスク")
	run("add", "完了タスク")
	run("punch", "2")

	out := run("list", "--pending")

	if !strings.Contains(out, "#1") {
		t.Errorf("--pending should show pending task #1, got: %s", out)
	}
	if strings.Contains(out, "#2") {
		t.Errorf("--pending should not show completed task #2, got: %s", out)
	}
}

func TestListCmd_DoneFlag_EmptyResult(t *testing.T) {
	withTempHome(t)

	run("add", "未完了のみ")

	out := run("list", "--done")

	if !strings.Contains(out, "No tasks found") {
		t.Errorf("--done with no completed tasks should show 'No tasks found', got: %s", out)
	}
}

func TestListCmd_PendingFlag_EmptyResult(t *testing.T) {
	withTempHome(t)

	run("add", "完了のみ")
	run("punch", "1")

	out := run("list", "--pending")

	if !strings.Contains(out, "No tasks found") {
		t.Errorf("--pending with all tasks done should show 'No tasks found', got: %s", out)
	}
}

func TestListCmd_BothFlags_Error(t *testing.T) {
	withTempHome(t)

	run("add", "タスク")

	out := run("list", "--done", "--pending")

	if !strings.Contains(out, "cannot use --done and --pending together") {
		t.Errorf("using both flags should show error message, got: %s", out)
	}
}

// --- punch ---

func TestPunchCmd_NoArgs_PunchesLatest(t *testing.T) {
	withTempHome(t)

	run("add", "タスクA")
	run("add", "タスクB")
	out := run("punch")

	if !strings.Contains(out, "Punched latest task #2") {
		t.Errorf("punch without args should punch latest task #2, got: %s", out)
	}
}

func TestPunchCmd_WithID_PunchesSpecified(t *testing.T) {
	withTempHome(t)

	run("add", "タスクA")
	run("add", "タスクB")
	out := run("punch", "1")

	if !strings.Contains(out, "Punched task #1") {
		t.Errorf("punch 1 should punch task #1, got: %s", out)
	}
}

func TestPunchCmd_ShowsListAfterPunch(t *testing.T) {
	withTempHome(t)

	run("add", "タスクA")
	out := run("punch")

	if !strings.Contains(out, "o") {
		t.Errorf("punch output should show completed mark o, got: %s", out)
	}
}

func TestPunchCmd_NoIncompleteTasks(t *testing.T) {
	withTempHome(t)

	out := run("punch")

	if !strings.Contains(out, "No incomplete tasks found") {
		t.Errorf("punch on empty docket should show no tasks message, got: %s", out)
	}
}

func TestPunchCmd_AlreadyDone(t *testing.T) {
	withTempHome(t)

	run("add", "タスクA")
	run("punch", "1")
	out := run("punch", "1")

	if !strings.Contains(out, "not found or already completed") {
		t.Errorf("punching already done task should show error, got: %s", out)
	}
}

// --- delete ---

func TestDeleteCmd_DeletesTask(t *testing.T) {
	withTempHome(t)

	run("add", "タスクA")
	out := run("delete", "1")

	if !strings.Contains(out, "Deleted task #1") {
		t.Errorf("delete 1 should show deleted message, got: %s", out)
	}
}

func TestDeleteCmd_NotFound(t *testing.T) {
	withTempHome(t)

	out := run("delete", "999")

	if !strings.Contains(out, "not found") {
		t.Errorf("delete non-existent ID should show 'not found', got: %s", out)
	}
}

func TestDeleteCmd_All_DeletesAllTasks(t *testing.T) {
	withTempHome(t)

	run("add", "タスクA")
	run("add", "タスクB")
	out := run("delete", "--all")

	if !strings.Contains(out, "Deleted all 2 tasks") {
		t.Errorf("delete --all should show deleted count, got: %s", out)
	}
}

func TestDeleteCmd_All_Empty(t *testing.T) {
	withTempHome(t)

	out := run("delete", "--all")

	if !strings.Contains(out, "No tasks to delete") {
		t.Errorf("delete --all on empty docket should show empty message, got: %s", out)
	}
}

// --- clear ---

func TestClearCmd_ResetsAll(t *testing.T) {
	withTempHome(t)

	run("add", "タスクA")
	run("add", "タスクB")
	out := run("clear")

	if !strings.Contains(out, "Cleared all tasks") {
		t.Errorf("clear should show reset message, got: %s", out)
	}
}

func TestClearCmd_ResetsNextID(t *testing.T) {
	withTempHome(t)

	run("add", "タスクA")
	run("add", "タスクB")
	run("clear")
	out := run("add", "新タスク")

	if !strings.Contains(out, "#1") {
		t.Errorf("after clear, next task should be #1, got: %s", out)
	}
}

// --- formatter ---

func TestFormatTaskAsTicket_Incomplete(t *testing.T) {
	withTempHome(t)

	run("add", "テスト")
	out := run("list")

	if strings.Contains(out, "→)") {
		t.Errorf("incomplete task should not have trailing arrow, got: %s", out)
	}
	if !strings.Contains(out, "[ ") {
		t.Errorf("incomplete task should have space status mark, got: %s", out)
	}
}

func TestFormatTaskAsTicket_Completed(t *testing.T) {
	withTempHome(t)

	run("add", "テスト")
	run("punch", "1")
	out := run("list")

	if !strings.Contains(out, "o") {
		t.Errorf("completed task should show o mark, got: %s", out)
	}
	if !strings.Contains(out, "→") {
		t.Errorf("completed task should show created→completed dates, got: %s", out)
	}
}

// --- edit ---

func TestEditCmd_EditsDescription(t *testing.T) {
	withTempHome(t)

	run("add", "元のタスク")
	out := run("edit", "1", "編集後のタスク")

	if !strings.Contains(out, "Edited task #1") {
		t.Errorf("edit output should contain 'Edited task #1', got: %s", out)
	}
	if !strings.Contains(out, "編集後のタスク") {
		t.Errorf("edit output should contain new description, got: %s", out)
	}

	listOut := run("list")
	if !strings.Contains(listOut, "編集後のタスク") {
		t.Errorf("after edit, list should show new description, got: %s", listOut)
	}
	if strings.Contains(listOut, "元のタスク") {
		t.Errorf("after edit, list should not show old description, got: %s", listOut)
	}
}

func TestEditCmd_NotFound(t *testing.T) {
	withTempHome(t)

	out := run("edit", "999", "存在しない")

	if !strings.Contains(out, "Task #999 not found") {
		t.Errorf("edit on non-existent ID should show 'Task #999 not found', got: %s", out)
	}
}

func TestEditCmd_ShowsListAfterEdit(t *testing.T) {
	withTempHome(t)

	run("add", "タスクA")
	run("add", "タスクB")
	out := run("edit", "1", "更新済みタスク")

	if !strings.Contains(out, "#1") || !strings.Contains(out, "#2") {
		t.Errorf("edit output should show full ticket list with #1 and #2, got: %s", out)
	}
	if !strings.Contains(out, "[") {
		t.Errorf("edit output should show ticket format, got: %s", out)
	}
}

func TestEditCmd_MultiWordDescription(t *testing.T) {
	withTempHome(t)

	run("add", "old")
	out := run("edit", "1", "new", "multi", "word", "description")

	if !strings.Contains(out, "new multi word description") {
		t.Errorf("edit should join multi-word args with spaces, got: %s", out)
	}
}

// --- priority ---

func TestAddCmd_WithPriority_ShowsInList(t *testing.T) {
	withTempHome(t)

	out := run("add", "--priority", "high", "重要タスク")

	if !strings.Contains(out, "[high]") {
		t.Errorf("add --priority high should show [high] in output, got: %s", out)
	}
}

func TestAddCmd_WithPriorityMedium_ShowsInList(t *testing.T) {
	withTempHome(t)

	out := run("add", "--priority", "medium", "中優先タスク")

	if !strings.Contains(out, "[medium]") {
		t.Errorf("add --priority medium should show [medium] in output, got: %s", out)
	}
}

func TestAddCmd_WithPriorityLow_ShowsInList(t *testing.T) {
	withTempHome(t)

	out := run("add", "--priority", "low", "低優先タスク")

	if !strings.Contains(out, "[low]") {
		t.Errorf("add --priority low should show [low] in output, got: %s", out)
	}
}

func TestAddCmd_WithPriority_InvalidValue(t *testing.T) {
	withTempHome(t)

	out := run("add", "--priority", "urgent", "タスク")

	if !strings.Contains(out, "Invalid priority") {
		t.Errorf("invalid priority should show error, got: %s", out)
	}
}

func TestListCmd_PriorityFilter(t *testing.T) {
	withTempHome(t)

	run("add", "--priority", "high", "高優先タスク")
	run("add", "--priority", "low", "低優先タスク")
	run("add", "優先度なし")

	out := run("list", "--priority", "high")

	if !strings.Contains(out, "高優先タスク") {
		t.Errorf("--priority high should show high priority task, got: %s", out)
	}
	if strings.Contains(out, "低優先タスク") {
		t.Errorf("--priority high should not show low priority task, got: %s", out)
	}
	if strings.Contains(out, "優先度なし") {
		t.Errorf("--priority high should not show task with no priority, got: %s", out)
	}
}

func TestListCmd_PriorityFilter_NoMatch(t *testing.T) {
	withTempHome(t)

	run("add", "優先度なし")

	out := run("list", "--priority", "high")

	if !strings.Contains(out, "No tasks found") {
		t.Errorf("--priority filter with no matches should show 'No tasks found', got: %s", out)
	}
}

// --- stderr を捨てる helper（log.Fatalf による os.Exit を回避するためテストは正常系のみ） ---

func init() {
	// テスト実行時に log の出力先を捨てる
	rootCmd.SetErr(io.Discard)
}
