# 🦞 docketpunch 🥊

[![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Build Status](https://github.com/Sottiki/docketpunch/workflows/CI/badge.svg)](https://github.com/Sottiki/docketpunch/actions)
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

```bash
go install github.com/Sottiki/docketpunch@latest
```

インストール後、`docketpunch` コマンドが使えるようになります。

## ⚡ エイリアス設定（推奨）

インストール直後は `docketpunch` コマンドで使用しますが、より快適に使うためエイリアスの設定をおすすめします。

**設定前:**
```bash
docketpunch list  # 長い...
```

**設定後:**
```bash
docket list  # 短くて便利！
dkt list     # さらに短く！
```

### zsh の場合

`~/.zshrc` に以下を追加：

```bash
alias docket='docketpunch'
alias dkt='docketpunch'
```

### bash の場合

`~/.bashrc` または `~/.bash_profile` に以下を追加：

```bash
alias docket='docketpunch'
alias dkt='docketpunch'
```

### 設定を反映

```bash
# zsh の場合
source ~/.zshrc

# bash の場合
source ~/.bashrc
```

以降、このREADMEでは短縮形の `docket` を使用して説明します。

## 🚀 使い方

### タスクを追加

```bash
docket add "Go言語の勉強をする"
```

```
✓ Added task #1: Go言語の勉強をする
```

### タスク一覧を表示

```bash
docket list
```

```
[ |#1|Go言語の勉強をする (10/06)]
[ |#2|ドキュメントを書く (10/07)]
[◯|#3|コードレビュー (10/05→10/08)]
```

- `[ ]` - 未完了タスク
- `[◯]` - 完了済みタスク（パンチされた！）
- 日付は `月/日` 形式で表示
- 完了済みタスクは `作成日→完了日` で表示

### タスクを完了（パンチ！）

#### 最新のタスクを完了

```bash
docket punch
```

```
✓ Punched latest task #2: ドキュメントを書く
```

引数なしで実行すると、最新の未完了タスクを完了します。

#### 指定したタスクを完了

```bash
docket punch 1
```

```
✓ Punched task #1: Go言語の勉強をする
```

### タスクを削除

```bash
docket delete 2
```

```
✓ Deleted task #2: ドキュメントを書く
```

### 完了済みタスクを一括削除

```bash
docket clear
```

```
✓ Cleared 2 completed task(s)
```

完了済み（`[◯]`）のタスクのみを削除し、未完了タスクはそのまま残ります。

### エイリアスを使った例

エイリアス設定後は、より短いコマンドで操作できます：

```bash
dkt add "Go言語の勉強をする"
dkt list
dkt punch          # 最新タスクを完了
dkt punch 1        # 指定IDを完了
dkt delete 2
dkt clear
```

## 📚 コマンド一覧

| コマンド | 説明 | 使用例 |
|---------|------|--------|
| `add <description>` | 新しいタスクを追加 | `docket add "タスク名"` |
| `list` | タスク一覧をチケット形式で表示 | `docket list` |
| `punch [task-id]` | タスクを完了にする<br>引数なし: 最新タスクを完了<br>引数あり: 指定IDを完了 | `docket punch`<br>`docket punch 1` |
| `delete <task-id>` | 指定したタスクを削除 | `docket delete 2` |
| `clear` | 完了済みタスクを一括削除 | `docket clear` |
| `help` | ヘルプを表示 | `docket help` |

## 🛠️ 開発状況

- [x] プロジェクト初期化
- [x] 基本設計
- [x] `add` コマンド実装
- [x] `list` コマンド実装
- [x] `punch` コマンド実装
- [x] `delete` コマンド実装
- [x] `clear` コマンド実装

## 🎯 今後の予定

- [ ] `delete --all` オプション（全タスク削除）
- [ ] コードリファクタリング
- [ ] テストコードの追加
- [ ] CI/CD の設定

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
│   ├── punch.go
│   ├── delete.go
│   └── clear.go
├── internal/
│   ├── task/
│   │   └── task.go          # Task構造体
│   ├── docket/
│   │   └── docket.go        # Docketビジネスロジック
│   └── storage/
│       └── storage.go       # JSON読み書き処理
├── main.go
├── go.mod
├── go.sum
└── README.md
```

## 📝 ライセンス

[MIT License](LICENSE)

## 👤 作者

**Sottiki**

- GitHub: [@Sottiki](https://github.com/Sottiki)

## 💭 インスピレーション

このプロジェクトは、古き良きパンチカードシステムからインスピレーションを得ています。
タスクを完了させる爽快感を、レトロな「パンチ」という動作で表現しました。

---

⭐ このプロジェクトが役に立ったら、スターをつけていただけると嬉しいです！
