# docketpunch プロジェクト概要

> 最終更新: 2026-04-20

## プロジェクト概要

レトロなパンチカードをイメージした Go 製 CLI タスク管理ツール。
タスクを完了するたびにパンチカードに穴を開けるような爽快感を提供する。

- **バージョン**: 0.1.0
- **言語**: Go 1.23+ (go.mod は 1.25.1)
- **CLI フレームワーク**: [Cobra](https://github.com/spf13/cobra) v1.10.1
- **データ保存先**: `~/.docket/tasks.json` (JSON)
- **リポジトリ**: https://github.com/Sottiki/docketpunch

---

## プロジェクト構造

```
docketpunch/
├── main.go
├── cmd/
│   ├── root.go          # ルートコマンド・ウェルカム画面・--version フラグ
│   ├── add.go           # add コマンド
│   ├── list.go          # list コマンド
│   ├── punch.go         # punch コマンド
│   ├── delete.go        # delete コマンド (--all フラグあり)
│   ├── clear.go         # clear コマンド (完全リセット)
│   ├── formatter.go     # formatTaskAsTicket() ヘルパー
│   └── cmd_test.go      # コマンド統合テスト (18ケース)
├── internal/
│   ├── task/
│   │   ├── task.go
│   │   └── task_test.go
│   ├── docket/
│   │   ├── docket.go
│   │   └── docket_test.go
│   └── storage/
│       ├── storage.go
│       └── storage_test.go
├── .github/
│   └── workflows/
│       ├── ci.yml        # CI: main への PR・マージ時にテスト & ビルド
│       └── release.yml   # リリース: タグプッシュ時にバイナリ配布
├── go.mod
├── go.sum
├── README.md
├── OVERVIEW.md
└── requirements.md
```

---

## データモデル

### Task 構造体 (`internal/task/task.go`)

| フィールド    | 型         | JSON キー       | 説明                          |
|-------------|------------|----------------|-------------------------------|
| ID          | int        | `id`           | 自動採番 (1始まり)             |
| Description | string     | `description`  | タスクの説明文                 |
| Done        | bool       | `done`         | 完了フラグ                    |
| CreatedAt   | time.Time  | `created_at`   | 作成日時                      |
| CompletedAt | *time.Time | `completed_at` | 完了日時 (未完了時は omitempty) |

### Docket 構造体 (`internal/docket/docket.go`)

| フィールド | 型       | JSON キー  | 説明                     |
|----------|---------|-----------|--------------------------|
| Tasks    | []*Task | `tasks`   | タスク一覧                |
| NextID   | int     | `next_id` | 次に割り当てる ID (自動採番) |

---

## 実装済みコマンド

### `docket add <description>`
- 新規タスクを追加し、追加後の全タスク一覧を表示
- **制約**: 引数は必ず1つ (`cobra.ExactArgs(1)`)

### `docket list`
- 全タスクをチケット形式で表示
- フォーマット: `[ |#ID|説明 (作成月/日)]` / `[◯|#ID|説明 (作成月/日→完了月/日)]`
- タスクがなければ `"No tasks found"` を表示

### `docket punch [task-id]`
- **引数なし**: 最新の未完了タスクを完了
- **引数あり**: 指定 ID のタスクを完了
- 完了後に全タスク一覧を表示
- **制約**: 引数は最大1つ (`cobra.MaximumNArgs(1)`)

### `docket delete <task-id> / --all`
- `delete <id>`: 指定 ID のタスクを削除
- `delete --all`: 全タスク削除（NextID は維持）

### `docket clear`
- 全タスク削除 + NextID を 1 にリセット
- **制約**: 引数なし (`cobra.NoArgs`)

### `docket` / `docket --version`
- ウェルカム画面を表示 (ASCII アート + バージョン番号)

---

## ビジネスロジック (`internal/docket/docket.go`)

| メソッド                      | 処理内容                                     |
|-----------------------------|----------------------------------------------|
| `NewDocket()`               | 空の Docket を初期化 (NextID=1)              |
| `AddTask(description)`      | 新タスク追加・NextID インクリメント          |
| `PunchTask(id)`             | 指定 ID の未完了タスクを完了                 |
| `GetLatestIncompleteTask()` | 逆順スキャンして最新未完了を返す             |
| `DeleteTask(id)`            | 指定 ID のタスクをスライスから削除           |
| `DeleteAllTasks()`          | 全タスク削除 (NextID 維持)                   |
| `ResetDocket()`             | 全タスク削除 + NextID=1 にリセット           |
| `ClearCompletedTasks()`     | 完了タスクを全削除し削除したタスクを返す     |

---

## ストレージ (`internal/storage/storage.go`)

- **ファイルパス**: `~/.docket/tasks.json`
- **`Init()`**: ディレクトリ作成 + ファイル未存在時は空の Docket を保存
- **`Save(d)`**: `json.MarshalIndent` でインデント付き JSON 書き込み
- **`Load()`**: `Init()` 呼び出し後に JSON 読み込み・パース

---

## テスト

| パッケージ           | テスト内容                                                                                                   |
|--------------------|--------------------------------------------------------------------------------------------------------------|
| `internal/task`    | `NewTask()`, 初期状態                                                                                        |
| `internal/docket`  | `NewDocket`, `AddTask`, `PunchTask`, `GetLatestIncompleteTask`, `DeleteTask`, `ClearCompletedTasks`, `ResetDocket`, `DeleteAllTasks` |
| `internal/storage` | `Save`/`Load` 往復, `Init` 冪等性, 不正 JSON エラー, `CompletedAt` 永続化, JSON フォーマット               |
| `cmd`              | 各コマンドの正常系・異常系 18ケース（os.Pipe で stdout キャプチャ、一時 HOME で分離）                       |

---

## CI/CD

### CI (`.github/workflows/ci.yml`)
- **トリガー**: main への PR 作成・更新時 / main へのマージ時
- **ジョブ**: テスト (`-race`) → ビルド（test 成功後）

### Release (`.github/workflows/release.yml`)
- **トリガー**: `v*.*.*` タグのプッシュ
- **成果物**: Linux amd64 / arm64 バイナリ + GitHub Release
- **制限**: macOS・Windows バイナリは未対応（`go install` で代替）
