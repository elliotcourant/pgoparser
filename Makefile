.PHONY: tools generate

tools:
	go build -o $$GOPATH/bin/keywordgen github.com/elliotcourant/pgoparser/tools/keywordgen

generate: tools
	PATH=$$GOPATH/bin:$$PATH; go generate ./... && go fmt ./...