
CYAN=\033[0;36m
RESET=\033[0m

pprint = echo -e "${CYAN}::>${RESET} ${1}"
completed = $(call pprint, Completed!)

.SILENT: deps
deps:
	$(call pprint, Downloading go libraries)
	go mod download
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
	go install github.com/golang/mock/mockgen@v1.6.0
	go install github.com/vektra/mockery/v2@latest
	go generate ./...
	$(call completed)

.SILENT: test
test: gen-mock
	$(call pprint, Running tests...)
	go test -v ./... -coverprofile .coverage.out
	$(call completed)