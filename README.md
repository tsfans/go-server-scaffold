# go-server-scaffold

A golang server scaffold for developing.

## Prerequisites

- golang 1.25.0 or above

## How to use

```shell
git clone git@github.com:tsfans/go-server-scaffold.git
cd go-server-scaffold
git submodule update --init --recursive
DOCKER_BUILDKIT=1 docker compose -f .docker/docker-compose.yaml up -d --build
```

## Project Structure

- .docker // docker files
- .vscode // ide settings
- bin
  - conf // yaml config template, rename to yaml to local debug
- server
  - controller // api entry
  - service // business logic
  - repository // database operation
  - model // business model
- [framework](https://github.com/tsfans/go-framework) // git submodule: a golang framework for developing server.
