# For Developers

## Usage

### Migration

migration directory: [./migrations](./migrations)

See the [goose documentation](https://pressly.github.io/goose/) for details.

#### migrate sql files

```shell
docker compose exec rss-generator goose.sh up
```

#### rollback sql files

```shell
docker compose exec rss-generator goose.sh down
```

### Schema

#### Add schema

If you would like to create the "Hoge" schema, you can do so with the following command.

```shell
docker compose exec rss-generator go run -mod=mod entgo.io/ent/cmd/ent new Hoge
```

#### Describe schema

```shell
docker compose exec rss-generator go run -mod=mod entgo.io/ent/cmd/ent describe ./ent/schema
```

#### Modify Models

Only the directory under [./ent/schema](./ent/schema) needs to be changed.
After the change, you can update the model as follows.

```shell
docker compose exec rss-generator go generate .
```

### Package

#### Install packages and modify mod files

You can install the package with the following command.

```shell
docker compose exec rss-generator go get -u <package name>
docker compose build rss-generator
```
