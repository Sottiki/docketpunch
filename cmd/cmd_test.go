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

// --- stderr を捨てる helper（log.Fatalf による os.Exit を回避するためテストは正常系のみ） ---

func init() {
	// テスト実行時に log の出力先を捨てる
	rootCmd.SetErr(io.Discard)
}
