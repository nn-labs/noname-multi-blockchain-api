.SILENT: deps lint clean gen-mock test build

CYAN=\033[0;36m
RESET=\033[0m

pprint = echo -e "${CYAN}::>${RESET} ${1}"
completed = $(call pprint,Completed!)

deps:
	$(call pprint, Downloading go libraries...)
	go mod download
	$(call completed)

lint:
	$(call pprint, Runnning linter...)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.42.1
	./bin/golangci-lint --version
	./bin/golangci-lint run ./...
	$(call completed)

clean:
	$(call pprint,Cleaning up...)
	rm -rf ./bin
	find . -name "mocks" | xargs rm -rf {}
	find . -name ".cover.out" | xargs rm -rf {}
	$(call completed)

gen-mock: clean deps
	$(call pprint, Generating mocks for tests...)
	go get github.com/golang/mock/mockgen
	export PATH=$PATH:$(go env GOPATH)/bin
	go generate ./...
	$(call completed)

test: gen-mock
	$(call pprint, Runnning tests...)
	cat > app.env << 'EOF'
	PORT=:5000
	APP_ENV=development
	GRPC_HOST=localhost
	BTC_RPC_ENDPOINT_TEST=localhost
	BTC_RPC_ENDPOINT_MAIN=localhost
	BTC_RPC_USER=user
	BTC_RPC_PASSWORD=password
	EOF
	go test ./... -coverprofile .cover.out
	$(call completed)

build: clean deps
	$(call pprint, Building app...)
	go build -o ./bin/server ./cmd
	$(call completed)