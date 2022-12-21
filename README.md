# golang-gqlgen-postgresql-example

## Offcial
    - https://gqlgen.com/getting-started/

## 構成
    - domain/entity: 構造体定義
    - config: 設定
    - graph: gqlgen自動生成(フロントへの返却)
    - io: DBクライアント生成
    - persistence: sqlクエリ(未実装)

## 開発手順(gqlgen)
    1. スキーマ定義 → ./graph/schema.graphqlsの修正
    2. gqlコマンド実行 → go gqlgen ./...
    3. ロジック実装 → ./graph/schema.resolvers.goの修正
   