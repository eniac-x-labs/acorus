GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell git show -s --format='%ct')

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

acorus:
	env GO111MODULE=on go build -v $(LDFLAGS) ./cmd/acorus

clean:
	rm acorus

test:
	go test -v ./...

lint:
	golangci-lint run ./...

.PHONY: \
	acorus \
	bindings \
	bindings-scc \
	clean \
	test \
	lint
