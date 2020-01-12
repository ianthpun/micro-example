# syntax = docker/dockerfile:experimental

FROM golang:1.13 as builder

copy proto app/proto
copy internal app/internal
copy go.mod app/go.mod
copy go.sum app/go.sum
copy cmd/greeter/service app/cmd/greeter/service

WORKDIR app

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
     GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build \
     -v -ldflags '-w -s' \
     -o /server \
     ./cmd/greeter/service

FROM gcr.io/distroless/base

COPY --from=builder /server /server

CMD ["/server"]
