build:
	go build -o bin/antilope main.go

run: build
	./bin/antilope

test:
	go test -v ./...