FROM golang:1.16.4-buster AS builder

ARG VERSION=dev
ARG GIT_INFO=

WORKDIR /go/src/app
COPY main.go .
RUN go build -ldflags="-X 'main.Env=${env}' -X 'main.GitInfo=$(GIT_INFO)' -X 'main.BuildTime=$(DATE)'" -o grpcServer2 ./cmd/grpcServer/main.go
RUN go build -ldflags="-X 'main.Env=${env}' -X 'main.GitInfo=$(GIT_INFO)' -X 'main.BuildTime=$(DATE)'" -o grpcServer2 ./cmd/grpcServer/main.go

FROM debian:buster-slim
COPY --from=builder /go/src/app/main /go/bin/main
ENV PATH="/go/bin:${PATH}"
CMD ["main"]