.PHONY: build api docs

UNIT_TEST_PACKAGES=$(shell go list ./... | grep -v "gen")
COVERAGE_MIN=18

setup:
	@go install github.com/swaggo/swag/cmd/swag@latest
	@go install github.com/golang/mock/mockgen@v1.6.0
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@pre-commit install
	@golangci-lint cache clean

build:
	@go mod tidy -compat=1.17

server: build
	@go run main.go api profiles/development.yml

docs:
	@swag init

tests:
	@go test ./... -config=profiles/test.yml -cover=true -convey-story=true -test.v

watch-tests:
	@$$GOPATH/bin/goconvey -excludedDirs="profiles,mocks"

tests-ci:
	ENVIRONMENT=test GOMAXPROCS=2 go test -vet=off -p 1 -cover -coverprofile=coverage-temp.out $(UNIT_TEST_PACKAGES)
	bash ./filter-coverage-report.sh
	@cat ./filtered-coverage-temp.out | grep -v "mock.go" > ./coverage.txt
	@go tool cover -func=coverage.txt | gawk '/total:.*statements/ {if (strtonum($$3) < $(COVERAGE_MIN)) {print "ERR: coverage is lower than $(COVERAGE_MIN)%", "Current Coverage: " $$3; exit 1} else {print "Current Coverage: " $$3}}'
	@go tool cover -func=coverage.txt | grep total > current-coverage.txt

tests-ci-main:
	git fetch origin main:main
	git checkout main
	ENVIRONMENT=test GOMAXPROCS=2 go test -p 1 -cover -coverprofile=coverage-temp-main.out $(UNIT_TEST_PACKAGES)
	@cat ./coverage-temp-main.out | grep -v "mock.go" > ./coverage-main.txt
	@go tool cover -func=coverage-main.txt | grep total > main-coverage.txt

generate-mocks:
	bash ./generate_mock.sh

qa env:
	SECRET_ENV=qa sh generateEnv