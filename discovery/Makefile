.PHONY: all dep test lint build 

WORKSPACE ?= $$(pwd)

GO_PKG_LIST := $(shell go list ./... | grep -v /vendor/)

lint:
	@echo "Linting 👀"
	@golint -set_exit_status ${GO_PKG_LIST}
	@echo "Linting completed 💪"

dep:
	@echo "Resolving go package dependencies"
	@go mod tidy
	@go mod vendor
	@echo "Package dependencies completed"

update-sdk:
	@echo "Updating SDK dependencies"
	@export GOFLAGS="" && go get "github.com/Axway/agent-sdk@main"


${WORKSPACE}/apic_discovery_agent: dep
	@echo "Building discovery agent 🐿"
	@export time=`date +%Y%m%d%H%M%S` && \
	export version=`cat version` && \
	export commit_id=`git rev-parse --short HEAD` && \
	go build -tags static_all \
		-ldflags="-X 'github.com/Axway/agent-sdk/pkg/cmd.BuildTime=$${time}' \
				-X 'github.com/Axway/agent-sdk/pkg/cmd.BuildVersion=$${version}' \
				-X 'github.com/Axway/agent-sdk/pkg/cmd.BuildCommitSha=$${commit_id}' \
				-X 'github.com/Axway/agent-sdk/pkg/cmd.BuildAgentName=WSO2DiscoveryAgent'" \
		-a -o ${WORKSPACE}/bin/apic_discovery_agent ${WORKSPACE}/main.go

build:${WORKSPACE}/apic_discovery_agent
	@echo "Build complete 🏆"

wso2-gen-model: ## generate WSO2 resources
	@echo "Generating WSO2 resources"
	@openapi-generator generate -g go -i ./specs/publisher-v1.yaml --package-name models --output pkg/wso2/models/ --global-property modelDocs=false --global-property models --global-property apiDocs=false
	@echo "Done generating WSO2 resources"