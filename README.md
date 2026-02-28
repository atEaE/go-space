# Go-Space — Vampire Survivors Mini

## 概要

Go製の2Dサバイバルアクションゲームです。Vampire Survivorsにインスパイアされたミニゲームで、押し寄せる敵を倒しながらレベルアップし、生き残ることを目指します。

## 💻技術スタック

### ゲームエンジン・フレームワーク

- **[Go](https://go.dev/)** 1.25.2: プログラミング言語
- **[Ebiten](https://ebitengine.org/)** v2.9.8: 2Dゲームエンジン
- **[Donburi](https://github.com/yohamta/donburi)** v1.15.7: ECS（Entity Component System）フレームワーク

### 開発ツール

- **[golangci-lint](https://golangci-lint.run/)**: コードのリンティング（`.golangci.yml`で設定済み）

## 📦セットアップ

### 必要な環境

- Go 1.25.2 以上

### インストール

```bash
go mod download
```

## 📁プロジェクト構造

```text
.
├── main.go                  # エントリーポイント（ウィンドウ 960x720）
└── internal/
    ├── game/                # Game構造体、シーン管理（Title / Playing）
    ├── component/           # ECSコンポーネント・タグ定義
    ├── archetype/           # エンティティファクトリ関数
    ├── system/              # ECSシステム（更新ロジック・描画）
    ├── config/              # 画面サイズ定数
    ├── event/               # イベント型定義
    └── layer/               # 描画レイヤー順序
```

## 🏗️アーキテクチャ

ECS（Entity Component System）アーキテクチャを採用しています。

## 🛠️開発コマンド

| コマンド | 説明 |
| -------- | ---- |
| `go mod download` | 依存パッケージのダウンロード |
| `go build -o ./out/go-space` | ゲームのビルド |
| `./out/go-space` | ゲームの実行 |
| `go test ./... -cover` | テストの実行（カバレッジ付き） |
| `go tool golangci-lint run ./...` | Lintチェックの実行 |

## 🔄開発フロー

1. コードを実装・テスト
2. `go tool golangci-lint run ./...` でLintチェックを通過させる
3. `go test ./... -cover` で全テストを通過させる
4. `go build -o ./out/go-space` でビルドの成功を確認する
