FROM golang:1.25.0 AS builder

WORKDIR /app
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -C server -o ../bin/server

FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/bin/server .
COPY --from=builder /app/bin/conf/server.yaml.template ./conf/server.yaml

ENV CONFIG_PATH=/app/conf/server.yaml
ENV CONFIG_AUTO_RELOAD=true
ENV CONFIG_RELOAD_INTERVAL=30

EXPOSE 8080
CMD ["./server"]
