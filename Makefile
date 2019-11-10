lint:
	go get github.com/securego/gosec/cmd/gosec
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b $(shell go env GOPATH)/bin v1.21.0
	gosec -quiet ./...
	golangci-lint run
dev:
	go test ./...
	go build -o ./bin/coveshare
	./bin/coveshare serve --config ./bin/config.yml
