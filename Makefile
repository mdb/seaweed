HOMEPAGE=https://github.com/mdb/seaweed
PREFIX=/usr/local

COVERAGE_FILE = coverage.out

VERSION=0.0.1
TAG=v$(VERSION)

ARCHIVE=seaweed-$(TAG).tar.gz
ARCHIVE_URL=$(HOMEPAGE)/archive/$(TAG).tar.gz

release: tag sha

tag:
	git tag --force latest
	git tag | grep $(TAG) || git tag --message "Release $(TAG)" --sign $(TAG)
	git push origin
	git push origin --force --tags

pkg/$(ARCHIVE): pkg/
	wget --output-document pkg/$(ARCHIVE) $(ARCHIVE_URL)

pkg/:
	mkdir pkg

sha: pkg/$(ARCHIVE)
	shasum pkg/$(ARCHIVE)

install: build
	go install

coverage: unit
	go tool cover -html=$(COVERAGE_FILE)

build: dependencies unit
	go build

test: dependencies
	go test -coverprofile=$(COVERAGE_FILE) -timeout 25ms

dependencies:
	go get -t
	go get golang.org/x/tools/cmd/cover

.PHONY: build coverage dependencies install release sha tag test
