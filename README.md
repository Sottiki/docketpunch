# 🦞 docketpunch 🥊

[![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![CI](https://github.com/Sottiki/docketpunch/actions/workflows/ci.yml/badge.svg)](https://github.com/Sottiki/docketpunch/actions/workflows/ci.yml)
[![Release](https://github.com/Sottiki/docketpunch/actions/workflows/release.yml/badge.svg)](https://github.com/Sottiki/docketpunch/actions/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/Sottiki/docketpunch)](https://goreportcard.com/report/github.com/Sottiki/docketpunch)

**A CLI task management tool - Punch your tasks to completion!**  
CLIタスク管理ツール - パンチカードのように、タスクに穴を開けて完了させる

---

## 🎫 コンセプト

レトロなパンチカードをイメージしたタスク管理CLIツール。  
コマンド+エンターで、タスクを完了するたびに、まるでチケットにパンチするように。

## ✨ 特徴

- 🚀 **高速起動** - Goで書かれた単一バイナリ
- 📝 **シンプル** - 直感的なコマンド構造
- 💾 **軽量** - JSONファイルでローカル保存（`~/.docket/tasks.json`）
- 🎨 **視覚的** - チケット風の美しい表示
- ⚡ **効率的** - キーボードだけで完結

## 📦 インストール

### 方法1: Go install（推奨）

```bash
go install github.com/Sottiki/docketpunch@latest
```

インストール先: `~/go/bin/docketpunch`

### 方法2: ビルド済みバイナリ（Linux）

#### Linux (amd64 - Intel/AMD)

```bash
curl -L https://github.com/Sottiki/docketpunch/releases/latest/download/docketpunch-linux-amd64 -o docketpunch
chmod +x docketpunch
sudo mv docketpunch /usr/local/bin/
```

#### Linux (arm64 - ARM)

```bash
curl -L https://github.com/Sottiki/docketpunch/releases/latest/download/docketpunch-linux-arm64 -o docketpunch
chmod +x docketpunch
sudo mv docketpunch /usr/local/bin/
```

### アンインストール

```bash
rm $(which docketpunch)
```

## ⚡ エイリアス設定（推奨）

`~/.zshrc` または `~/.bashrc` に以下を追加：

```bash
alias docket='docketpunch'
alias dkt='docketpunch'
```

以降、このREADMEでは短縮形の `docket` を使用して説明します。

## 🚀 使い方

### タスクを追加

```bash
docket add "Go言語の勉強をする"
docket add --priority high "重要なバグ修正"   # 優先度付きで追加
docket add -p medium "ドキュメント更新"       # 短縮フラグも使用可
```

```
Added task: #1 : Go言語の勉強をする
[ |#1|Go言語の勉強をする (10/06)]
```

### タスク一覧を表示

```bash
docket list
```

```
[ |#1|Go言語の勉強をする (10/06)]
[ |#2|ドキュメントを書く [high] (10/07)]
[o|#3|コードレビュー (10/05→10/08)]
```

- `[ ]` - 未完了タスク（作成日を表示）
- `[o]` - 完了済みタスク（作成日→完了日を表示）
- `[high]` / `[medium]` / `[low]` - 優先度タグ（設定時のみ表示）

#### フィルター表示

```bash
docket list --pending          # 未完了のみ
docket list --done             # 完了済みのみ
docket list --priority high    # 優先度でフィルタ
docket list --priority high --pending  # 組み合わせも可
```

### タスクを完了（パンチ！）

#### 最新のタスクを完了

```bash
docket punch
```

```
✓ Punched latest task #2: ドキュメントを書く
[ |#1|Go言語の勉強をする (10/06)]
[◯|#2|ドキュメントを書く (10/07→10/08)]
[◯|#3|コードレビュー (10/05→10/08)]
```

#### 指定したタスクを完了

```bash
docket punch 1
```

### タスクの説明を編集

```bash
docket edit 1 "新しいタスク名"
```

```
Edited task #1: 新しいタスク名
[ |#1|新しいタスク名 (10/06)]
[ |#2|ドキュメントを書く [high] (10/07)]
```

### タスクを削除

```bash
docket delete 2
```

```
✓ Deleted task #2: ドキュメントを書く
```

#### 全タスクを削除（連番は継続）

```bash
docket delete --all
```

```
✓ Deleted all 3 tasks.
```

### リセット

```bash
docket clear
```

```
✓ Cleared all tasks and reset ID counter.
```

全タスクを削除し、IDの連番を **1からリセット** します。

## 📚 コマンド一覧

| コマンド | 説明 | 使用例 |
|---------|------|--------|
| `add <description>` | 新しいタスクを追加 | `docket add "タスク名"` |
| `add --priority <level>` | 優先度付きでタスクを追加 | `docket add -p high "重要タスク"` |
| `list` | タスク一覧をチケット形式で表示 | `docket list` |
| `list --pending` | 未完了タスクのみ表示 | `docket list --pending` |
| `list --done` | 完了済みタスクのみ表示 | `docket list --done` |
| `list --priority <level>` | 優先度でフィルタ表示 | `docket list --priority high` |
| `edit <task-id> <description>` | タスクの説明文を編集 | `docket edit 1 "新しい名前"` |
| `punch [task-id]` | タスクを完了にする<br>引数なし: 最新タスクを完了<br>引数あり: 指定IDを完了 | `docket punch`<br>`docket punch 1` |
| `delete <task-id>` | 指定したタスクを削除 | `docket delete 2` |
| `delete --all` | 全タスクを削除（連番継続） | `docket delete --all` |
| `clear` | 全タスク削除 + IDを1にリセット | `docket clear` |
| `help` | ヘルプを表示 | `docket help` |

## 🛠️ 開発状況

- [x] プロジェクト初期化・基本設計
- [x] `add` / `list` / `punch` / `delete` / `clear` コマンド実装
- [x] `delete --all` オプション（全タスク削除）
- [x] `clear` コマンドの完全リセット機能（連番リセット）
- [x] テストコードの追加（unit / integration）
- [x] CI/CD の設定
- [x] `list --done` / `list --pending` フィルター表示
- [x] `edit` コマンド（タスクの説明文を編集）
- [x] タスク優先度 high/medium/low（`--priority` フラグ）

## 🏗️ 技術スタック

- **言語**: Go 1.23+
- **CLI フレームワーク**: [Cobra](https://github.com/spf13/cobra)
- **データ保存**: JSON (`~/.docket/tasks.json`)

## 📂 プロジェクト構造

```
docketpunch/
├── cmd/
│   ├── root.go
│   ├── add.go
│   ├── list.go
│   ├── edit.go
│   ├── punch.go
│   ├── delete.go
│   ├── clear.go
│   ├── formatter.go     # チケット表示ヘルパー
│   └── cmd_test.go      # コマンド統合テスト
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
│       ├── ci.yml        # CI: main PR・マージ時にテスト実行
│       └── release.yml   # リリース: タグプッシュ時にバイナリ配布
├── main.go
├── go.mod
└── go.sum
```

## 📝 ライセンス

[MIT License](LICENSE)

## 👤 作者

**Sottiki**

- GitHub: [@Sottiki](https://github.com/Sottiki)

## 💭 インスピレーション

このプロジェクトは、古き良きパンチカードシステムからインスピレーションを得ています。
タスクを完了させる爽快感を、レトロな「パンチ」という動作で表現しました。

