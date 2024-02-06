GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell git show -s --format='%ct')

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

L1POOL_ABI_ARTIFACT := bridge-contracts/out/L1PoolManager.sol/L1PoolManager.json
L2POOL_ABI_ARTIFACT := bridge-contracts/out/L2PoolManager.sol/L2PoolManager.json
MSG_ABI_ARTIFACT := bridge-contracts/out/MessageManager.sol/MessageManager.json

acorus:
	env GO111MODULE=on go build -v $(LDFLAGS) ./cmd/acorus

clean:
	rm acorus

test:
	go test -v ./...

lint:
	golangci-lint run ./...

bindings: binding-l1p binding-l2p binding-msg

binding-l1p:
	$(eval temp := $(shell mktemp))

	cat $(L1POOL_ABI_ARTIFACT) \
	| jq -r .bytecode > $(temp)

	cat $(L1POOL_ABI_ARTIFACT) \
	| jq .abi \
	| abigen --pkg bindings \
	--abi - \
	--out bindings/l1_pool_manager.go \
	--type L1PoolManager \
	--bin $(temp)

	rm $(temp)

binding-l2p:
	$(eval temp := $(shell mktemp))

	cat $(L2POOL_ABI_ARTIFACT) \
	| jq -r .bytecode > $(temp)

	cat $(L2POOL_ABI_ARTIFACT) \
	| jq .abi \
	| abigen --pkg bindings \
	--abi - \
	--out bindings/l2_pool_manager.go \
	--type L2PoolManager \
	--bin $(temp)

	rm $(temp)

binding-msg:
	$(eval temp := $(shell mktemp))

	cat $(MSG_ABI_ARTIFACT) \
	| jq -r .bytecode > $(temp)

	cat $(MSG_ABI_ARTIFACT) \
	| jq .abi \
	| abigen --pkg bindings \
	--abi - \
	--out bindings/message_manager.go \
	--type MessageManager \
	--bin $(temp)

	rm $(temp)

.PHONY: \
	acorus \
	bindings \
	bindings-l1p \
	bindings-l2p \
	bindings-msg \
	clean \
	test \
	lint
