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

昔の車掌さんがチケットに穴を開けていたように、あなたのタスクリストに「完了の印」を刻みましょう。

## ✨ 特徴

- 🚀 **高速起動** - Goで書かれた単一バイナリ
- 📝 **シンプル** - 直感的なコマンド構造
- 💾 **軽量** - JSONファイルでローカル保存
- 🎨 **視覚的** - 美しいテーブル表示
- ⚡ **効率的** - キーボードだけで完結

## 📦 インストール

Coming soon ...

## ⚡ エイリアス設定（推奨）

より快適に使うため、エイリアスの設定をおすすめします。お好みのものを選んでください。

### zsh の場合

`~/.zshrc` に以下のいずれかを追加：

```bash
alias dkt='docket'   # 推奨：短くてタイプしやすい
alias dk='docket'    # より短く
alias dt='docket'    # さらに短く
```

### bash の場合

`~/.bashrc` または `~/.bash_profile` に以下のいずれかを追加：

```bash
alias dkt='docket'   # 推奨：短くてタイプしやすい
alias dk='docket'    # より短く
alias dt='docket'    # さらに短く
```

設定後、以下で反映：

```bash
# zsh の場合
source ~/.zshrc

# bash の場合
source ~/.bashrc
```

## 🚀 使い方

### タスクを追加

```bash
docket add "Go言語の勉強をする"
```

```
✓ タスクを追加しました (ID: 1)
```

### タスク一覧を表示

```bash
docket list
```

```
ID  状態    タスク                    作成日
─────────────────────────────────────────────
1   [ ]     Go言語の勉強をする        2025-10-06
2   [ ]     ドキュメントを書く        2025-10-06
3   [✓]     コードレビュー            2025-10-05
```

### タスクを完了

```bash
docket punch 1
```

```
🥊 パンチ！タスク #1 を完了しました
```

### タスクを削除

```bash
docket delete 2
```

```
✓ タスク #2 を削除しました
```

### エイリアスを使った例

エイリアス設定後は、より短いコマンドで操作できます：

```bash
dkt add "Go言語の勉強をする"
dkt list
dkt punch 1
dkt delete 2
```

## 📚 コマンド一覧

| コマンド | 説明 | 使用例 |
|---------|------|--------|
| `add` | 新しいタスクを追加 | `docket add "タスク名"` |
| `list` | タスク一覧を表示 | `docket list` |
| `punch` | タスクを完了にする | `docket punch <ID>` |
| `delete` | タスクを削除 | `docket delete <ID>` |
| `help` | ヘルプを表示 | `docket help` |

## 🛠️ 開発状況

- [x] プロジェクト初期化
- [x] 基本設計
- [ ] `add` コマンド実装
- [ ] `list` コマンド実装
- [ ] `punch` コマンド実装
- [ ] `delete` コマンド実装

## 🏗️ 技術スタック

- **言語**: Go
- **データ保存**: JSON

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
