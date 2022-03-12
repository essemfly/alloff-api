#!/bin/bash

ssh alloff-dev 'cd alloff-api && git pull && /usr/local/go/bin/go build -o apiServer ./cmd/apiServer/main.go && /usr/local/go/bin/go build -o grpcServer ./cmd/grpcServer/main.go && sudo systemctl restart alloff-dev.service && sudo systemctl restart alloff-grpc-dev.service'
