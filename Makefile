CYAN=\033[0;36m
RESET=\033[0m

pprint = echo -e "${CYAN}::>${RESET} ${1}"
completed = $(call pprint, Completed!)

.SILENT: deps
deps:
	$(call pprint, Downloading go libraries)
	go mod download
	$(call completed)

.SILENT: lint
lint:
	$(call pprint, Runnning linter...)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.42.1
	./bin/golangci-lint --version
	./bin/golangci-lint run ./...
	$(call completed)

.SILENT: clean
clean:
	$(call pprint, Cleanning up...)
	rm -rf ./bin
	find . -name "mocks" | xargs rm -rf {}
	find . -name ".cover.out" | xargs rm -rf {}
	$(call completed)

.SILENT: gen-mock
gen-mock: clean deps
	$(call pprint, Generating mocks for tests...)
	go get github.com/golang/mock
	go generate ./...
	$(call completed)

.SILENT: test
test: gen-mock
	$(call pprint, Running tests...)
	go test -v ./... -coverprofile .coverage.out
	$(call completed)

.SILENT: build
build: clean deps
	$(call pprint, Building app...)
	go build -o ./bin/server ./cmd
	$(call completed)