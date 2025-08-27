# go-server-scaffold

A golang server scaffold for developing.

## Prerequisites

- golang 1.25.0 or above

## How to use

```shell
git clone git@github.com:tsfans/go-server-scaffold.git
cd go-server-scaffold
go mod edit ${your-package}
DOCKER_BUILDKIT=1 docker compose -f .docker/docker-compose.yaml up -d
```

## Project Structure

- .docker // docker files
- .vscode // ide settings
- bin
  - conf // local config file
- framework
  - config // load yaml config
  - database // database operation
  - logger // log operation
  - server // http server
  - utils // utils
- server
  - controller // api entry
  - service // business logic
  - repository // database operation
  - model // business model
