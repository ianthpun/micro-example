# syntax = docker/dockerfile:experimental

FROM golang:1.13 as builder

ARG SERVICE_DIR

copy proto app/proto
copy internal app/internal
copy go.mod app/go.mod
copy go.sum app/go.sum
copy vendor app/vendor/
copy ${SERVICE_DIR} app/${SERVICE_DIR}/

WORKDIR app

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
     GO111MODULE=on CGO_ENABLED=0 GOOS=linux \
     go build \
     -mod=vendor \
     -v \
     -ldflags '-w -s' \
     -o /client \
     ./${SERVICE_DIR}

FROM gcr.io/distroless/base

COPY --from=builder /client /client

CMD ["/client"]
