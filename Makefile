bundle:
	go get github.com/rakyll/statik
	statik -src=./server/public
lint:
	go get github.com/securego/gosec/cmd/gosec
	gosec -quiet ./...
	golint ./...
test:
	go test ./...
dev: bundle
	go build -o ./bin/coveshare
	./bin/coveshare serve --base_url "http://localhost:8080" --config ./bin/config.yml 
publish: bundle lint test
	go build -ldflags="-s -w" -o "./bin/coveshare_$(shell date +'%y%m%d%H%M%S')"
