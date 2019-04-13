all: updatedeps test install

updatedeps:
	@go tool cover 2>/dev/null; if [ $$? -eq 3 ]; then \
		go get -u golang.org/x/tools/cmd/cover; \
	fi
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

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
