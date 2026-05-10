# docketpunch 実装要件・TODO

> 最終更新: 2026-05-10

## 完了済み TODO

- [x] `internal/storage` テスト追加 (PR #16)
- [x] `internal/docket` に `ResetDocket` / `DeleteAllTasks` のメソッドとテスト追加 (PR #17)
- [x] `clear` コマンドを JSON 完全リセットに変更（全タスク削除 + NextID=1）(PR #18)
- [x] `delete --all` オプションを追加（全タスク削除・NextID はリセットしない）(PR #19)
- [x] `createSampleData()` を削除・初回起動時は空の Docket で開始 (PR #20)
- [x] 著作権ヘッダを修正 (PR #20)
- [x] `list` コマンドの未完了タスク日付表示を修正: `(01/02→)` → `(01/02)` (PR #21)
- [x] `punch` 実行後にタスク一覧を表示 (PR #22)
- [x] `formatTaskAsTicket` を `cmd/formatter.go` に分離 (PR #23)
- [x] `cmd/` 統合テスト追加（18ケース） (PR #24)
- [x] GitHub Actions CI ワークフロー追加 (PR #25)
- [x] README を最新実装に合わせて更新
- [x] `ClearCompletedTasks` デッドコード削除 (PR #26)
- [x] 完了マークを `◯` → `o` に変更（Windows 文字化け対応）(PR #27)
- [x] `list --done` / `list --pending` フィルタフラグを追加 (PR #28)
- [x] `edit` コマンドを追加（タスクの説明文を編集）(PR #29)

---

## コマンド設計まとめ（現在）

| コマンド | 動作 |
|---------|------|
| `add <description>` | 新規タスクを追加（追加後に一覧表示） |
| `list` | 全タスクをチケット形式で表示 |
| `punch [id]` | 引数なし: 最新未完了を完了 / 引数あり: 指定 ID を完了（完了後に一覧表示） |
| `delete <id>` | 指定 ID のタスクを削除 |
| `delete --all` | 全タスク削除（NextID はリセットしない） |
| `clear` | JSON を完全リセット（全タスク削除 + NextID=1） |
| `edit <id> <新しい説明>` | 指定 ID のタスクの説明文を編集 |
| `list --done` | 完了済みタスクのみ表示 |
| `list --pending` | 未完了タスクのみ表示 |

---

## エンハンス候補

### 機能追加

| 優先度 | 機能 | 概要 |
|--------|------|------|
| 中 | タスク優先度 | high/medium/low を付けてソート・フィルタ |
| 低 | カラー出力 | ターミナルカラーで状態を視覚的に区別 |
| 低 | macOS/Windows バイナリ | CI/CD でクロスプラットフォームビルドを追加 |
| 低 | 複数データファイル対応 | プロジェクト別にタスクを分けて管理 |
