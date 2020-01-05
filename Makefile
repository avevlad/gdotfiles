GIT_COMMIT=$(shell git rev-parse --short HEAD)
VERSION=0.1.0

GO_FLAGS += -mod=vendor
GO_FLAGS += -ldflags="\
-X 'github.com/avevlad/gdotfiles/internal/build.Revision=$(GIT_COMMIT)'\
-X 'github.com/avevlad/gdotfiles/internal/build.Version=$(VERSION)'\
"\

build:
	go build $(GO_FLAGS) -o bin/gdotfiles ./cmd/gdotfiles

test-common: test vet

test:
	go test ./internal/...

run:
	./bin/gdotfiles --name=Node
	./bin/gdotfiles -v

vet:
	go vet ./cmd/... ./internal/...
