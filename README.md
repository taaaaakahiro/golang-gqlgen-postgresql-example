# golang-gqlgen-postgresql-example

## Run
```
$ docker-compose up -d #run db
$ make run
```
## Test
```
$ make test
```

## Tree
    - domain/entity: 構造体定義
    - config: 設定
    - graph: gqlgen自動生成(フロントへの返却)
    - io: DBクライアント生成
    - persistence: sqlクエリ(未実装)

## Procedure(gqlgen)
    1. スキーマ定義 → ./graph/schema.graphqlsの修正
    2. gqlコマンド実行 → go gqlgen ./...
    3. ロジック実装 → ./graph/schema.resolvers.goの修正

## wiki
    https://github.com/taaaaakahiro/golang-gqlgen-postgresql-example/wiki

## Docs
- gqlgen
    https://gqlgen.com/getting-started/
- validator
    https://zenn.dev/sntree/articles/b8585ee1ce219f
    https://github.com/Sntree2mi8/gqlgen-validator-sample
    https://pkg.go.dev/github.com/go-playground/validator/v10#hdr-Singleton