build:
	go build -o ./bin/mochain

run: build
	./bin/mochain

test: 
	go test -v ./...