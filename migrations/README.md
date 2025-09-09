# Migration
## Documentation
https://atlasgo.io/guides/orms/gorm/getting-started

## Installation
Install Atlas from macOS or Linux by running:
```shell
curl -sSf https://atlasgo.sh | sh
```

## Execute Migration Command
### Postgres
```shell
cd migrations
atlas migrate -c file://postgres_atlas.hcl diff --env gorm --var "postgres_url=$DEBANK_POSTGRES_MIGRATION"
```