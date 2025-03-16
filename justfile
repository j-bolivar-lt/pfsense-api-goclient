build:
    docker build -t pfsense-api-test:latest .

run: build
    docker run --rm pfsense-api-test:latest

test:
	go test -v -race -coverprofile coverage.txt -covermode atomic ./pfsenseapi/...
