# syntax = docker/dockerfile:experimental

FROM golang:1.13 as tester
COPY go.mod /test/go.mod
COPY go.sum /test/go.sum
copy vendor /test/vendor
COPY internal /test/internal
COPY proto /test/proto
COPY cmd /test/cmd/

WORKDIR /test

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
     GO111MODULE=on CGO_ENABLED=0 GOOS=linux \
     go test \
     -mod=vendor \
     -v \
     ./...

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
     GO111MODULE=on CGO_ENABLED=0 GOOS=linux \
     go vet \
     -mod=vendor \
     -v \
     ./...

FROM golang:1.13 as builder

ARG SERVICE_DIR

copy --from=tester /test/go.mod app/go.mod
copy --from=tester /test/go.sum app/go.sum
copy --from=tester /test/vendor app/vendor/
copy --from=tester /test/internal app/internal
copy --from=tester /test/proto app/proto
copy --from=tester /test/${SERVICE_DIR} app/${SERVICE_DIR}/

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
     -o /server \
     ./${SERVICE_DIR}

FROM gcr.io/distroless/base

COPY --from=builder /server /server

CMD ["/server"]
