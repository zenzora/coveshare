lint:
	go get github.com/securego/gosec/cmd/gosec
	gosec -quiet ./...
	golint ./...
dev:
	go test ./...
	go build -o ./bin/coveshare
	./bin/coveshare serve --config ./bin/config.yml
publish:
	go build -ldflags="-s -w" -o "./bin/coveshare_$(shell date +'%y%m%d%H%M%S')"
