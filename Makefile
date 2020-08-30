.PHONY: tools generate

generate:
	PATH=$$GOPATH/bin:$$PATH; go generate ./... && go fmt ./...