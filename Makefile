.PHONY: run
run: build
	./onchain-merklized-issuer-demo

.PHONY: build
build:
	go build -o onchain-merklized-issuer-demo .

.PHONY: test
test:
	go test -v ./...

.PHONY: lint
lint:
	golangci-lint run
