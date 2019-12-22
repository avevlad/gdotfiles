GIT_COMMIT=$(shell git rev-parse --short HEAD)

GO_FLAGS += -ldflags="-X 'main.revision=$(GIT_COMMIT)'"

build:
	go build $(GO_FLAGS) -o bin/gignore ./cmd/gignore

test-common: test vet

test:
	go test ./internal/...

vet:
	go vet ./cmd/... ./internal/...
