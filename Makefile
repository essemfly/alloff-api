GO=go

build:
	$(eval GIT_INFO=$(shell git show --pretty=format:%cs --no-patch)-$(shell git show --pretty=format:%h --no-patch))
	$(eval DATE=$(shell echo `date` `time`))
	$(GO) build -ldflags="-X 'main.Env=${env}' -X 'main.GitInfo=$(GIT_INFO)' -X 'main.BuildTime=$(DATE)'" -o grpcServer2 ./cmd/grpcServer/main.go
	$(GO) build -ldflags="-X 'main.Env=${env}' -X 'main.GitInfo=$(GIT_INFO)' -X 'main.BuildTime=$(DATE)'" -o apiServer2 ./cmd/apiServer/main.go