#!/bin/bash

sudo cp ./productsCrawler ./productsCrawler-last
/usr/local/go/bin/go build -o productsCrawler ./cmd/productsCrawler/main.go
sudo cp ./productDiffNotifier ./productDiffNotifier-last
/usr/local/go/bin/go build -o productDiffNotifier ./cmd/productDiffNotifier/main.go
sudo cp ./paymentSweeper ./paymentSweeper-last
/usr/local/go/bin/go build -o paymentSweeper ./cmd/paymentSweeper/main.go
