.PHONY: build
build:
	set GOARCH=amd64
	go env -w GOARCH=amd64
	set GOOS=linux
	go env -w GOOS=linux


	go build -o deploy/app/rpc/workflow-rpc workflow/rpc/workflow.go
	go build -o deploy/app/api/workflow-api workflow/api/workflow.go


