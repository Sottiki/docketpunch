# docketpunch 開発ガイド

## 開発方針

- 小さなコミットを心がける
- 機能単位で完結した、小さな PR を作る
- テストを先に書いてから実装する（TDD）
- 仕様を変更したときはドキュメントを更新する
- 新しい機能を追加したときはドキュメントを更新する

## コマンド

```bash
go build ./...      # ビルド確認
go test ./...       # 全テスト実行
go run . <command>  # 直接実行
go install .        # ~/go/bin/docketpunch にインストール
```

## アーキテクチャ

```
cmd/          コマンド定義（cobra）
  add.go      タスク追加（--priority/-p フラグ）
  list.go     一覧表示（--done / --pending / --priority フィルタ）
  edit.go     説明文編集
  punch.go    タスク完了
  delete.go   削除
  clear.go    全リセット
  formatter.go  チケット表示ヘルパー
  cmd_test.go   統合テスト

internal/
  task/       Task 構造体（ID, Description, Priority, Done, CreatedAt, CompletedAt）
  docket/     Docket 構造体 + ビジネスロジック
  storage/    JSON 読み書き（~/.docket/tasks.json）
```

## テスト規約

- `cmd_test.go` の `run(args...)` と `withTempHome(t)` ヘルパーを再利用する（再定義しない）
- `withTempHome(t)` で HOME を一時ディレクトリに差し替えてテスト間の状態を分離する

## データ形式

`~/.docket/tasks.json` に保存。`priority` フィールドは省略可能（`omitempty`）。
既存データとの後方互換性あり（priority なしのデータはそのまま動作する）。
