SDK_ONLY_PKGS=$(shell go list ./... | grep -v "/vendor/")

all: build unit

help:
	@echo "Please use \`make <target>' where <target> is one of"
	@echo "  build                   to go build the SDK"
	@echo "  unit                    to run unit tests"
	@echo "  lint                    to lint the SDK"

build:
	@echo "go build SDK and vendor packages"
	@go build ${SDK_ONLY_PKGS}

unit:  build 
	@echo "go test SDK and vendor packages"
	@go test   $(SDK_ONLY_PKGS)

lint:  build 
	@echo "go lint netscaler package (ignoring generated packages)"
	@golint   netscaler | grep -v netscaler/resources.go || true

generate:
	@echo "Generate go schema from json schema"
	(cd tools; ./generate.sh)

