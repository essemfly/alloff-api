#!/bin/bash

sudo cp ./apiServer ./apiServer-last
sudo cp ./grpcServer ./grpcServer-last
/usr/local/go/bin/go build -o apiServer ./cmd/apiServer/main.go && /usr/local/go/bin/go build -o grpcServer ./cmd/grpcServer/main.go && sudo systemctl restart alloff.service && sudo systemctl restart alloff-grpc.service