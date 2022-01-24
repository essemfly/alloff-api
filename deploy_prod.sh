#!/bin/bash

ssh alloff 'cd alloff-api && git pull && /usr/local/go/bin/go build -o apiServer ./cmd/apiServer/main.go && /usr/local/go/bin/go build -o grpcServer ./cmd/grpcServer/main.go && sudo systemctl restart alloff.service && sudo systemctl restart alloff-grpc.service'
