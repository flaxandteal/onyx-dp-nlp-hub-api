BINPATH ?= build

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

BUILD_TIME=$(shell date +%s)
GIT_COMMIT=$(shell git rev-parse HEAD)
VERSION ?= $(shell git tag --points-at HEAD | grep ^v | head -n 1)

LDFLAGS = -ldflags "-X main.BuildTime=$(BUILD_TIME) -X main.GitCommit=$(GIT_COMMIT) -X main.Version=$(VERSION)"

.PHONY: all
all: delimiter-AUDIT audit delimiter-LINTERS lint delimiter-UNIT-TESTS test delimiter-COMPONENT_TESTS test-component delimiter-FINISH ## Runs multiple targets, audit, lint, test and test-component

.PHONY: audit
audit: ## Audits and finds vulnerable dependencies
	go list -json -m all | nancy sleuth

.PHONY: build 
build: Dockerfile ## Builds ./Dockerfile image name: scrubber
	docker build -t hub_container .

.PHONY: build-bin
build_bin: ## builds bin
	go build -tags 'production' $(LDFLAGS) -o $(BINPATH)/nlp_hub

.PHONY: clean
clean: ## Removes /bin folder
	rm -fr ./build
	rm -fr ./vendor

.PHONY: convey
convey: ## Runs Convey tests
	goconvey ./...
	
.PHONY: debug
debug: ## Runs application locally with debug mode on
	go build -tags 'debug' $(LDFLAGS) -o $(BINPATH)/dp-nlp-hub
	HUMAN_LOG=1 DEBUG=1 $(BINPATH)/dp-nlp-hub

.PHONY: delimiter-%
delimiter-%:
	@echo '===================${GREEN} $* ${RESET}==================='

.PHONY: fmt 
fmt: ## Formats the code using go fmt and go vet
	go fmt ./...
	go vet ./...

.PHONY: lint 
lint: ## Automated checking of your source code for programmatic and stylistic errors
	golangci-lint run ./...

.PHONY: lint-local 
lint-local: ## Automated checking of your source code for programmatic and stylistic errors
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@
	golangci-lint run ./...

.PHONY: run 
run: ## Runs container name: hub from image name: nlp_hub
	docker run -p 5000:5000 --name hub_container -ti --rm nlp_hub

.PHONY: run-locally 
run-locally: ## Run the app locally
	go run .
 
.PHONY: test
test: ## Runs standard unit test tests
	go test -race -cover ./... 

.PHONY: test-all
test: convey test-component	test ## Runs all tests with -race and -cover flags
	go test -race -cover ./...

.PHONY: test-component
test-component: ## Test components
	go test -cover -coverpkg=github.com/ONSdigital/dp-nlp-hub/... -component

.PHONY: update
update: ## Go gets all of the dependencies and downloads them
	go get .
	go mod download

help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)