BINARY_NAME=logparser
BUILD_DIR=build

# Misc
.DEFAULT_GOAL = help
.PHONY        = help test live clean

## —— Makefile ————————————————————————————————————————————————————————————————
help: ## Outputs this help screen
	@grep -E '(^[a-zA-Z0-9_-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}{printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'

## —— Tests ———————————————————————————————————————————————————————————————————
test: ## Run tests
	go clean -testcache
	go test ./...

coverage: ## Run tests with coverage
	mkdir -p ${BUILD_DIR}
	go test -coverprofile=${BUILD_DIR}/coverage.out ./...
	go tool cover -html=${BUILD_DIR}/coverage.out

build: ## Build the binary file
	mkdir -p ${BUILD_DIR}
	go build -o ${BUILD_DIR}/${BINARY_NAME} main.go

serve: build ## Run the binary file
	${BUILD_DIR}/${BINARY_NAME} serve

live: ## Run the binary file with live reload
	air

clean: ## Remove previous build
	go clean
	rm -f ${BINARY_NAME}
	rm -Rf ./${BUILD_DIR}
