all: test install

fmt:
	gofmt -s -w .

test:
	go vet
	go test -cover

cover:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out
	rm coverage.out

install:
	go install
