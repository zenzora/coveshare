dev:
	go test ./...
	go build -o ./bin/coveshare
	./bin/coveshare serve --config ./bin/config.yml
