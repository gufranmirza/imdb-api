export BINARY_NAME = "imdb-api"
export CONFIG_FILE = "app-config.json"

# global command
.PHONY: all
all: dependencies build test

.PHONY: dependencies
dependencies:
	go mod download

.PHONY: build
build:
	env GOOS=linux GOARCH=amd64 go build -o ../bin/${BINARY_NAME} ./
	cp -r ${CONFIG_FILE} ../bin

.PHONY: run
run:
	go run ./main.go

.PHONY: start
start:
	../bin/${BINARY_NAME}

.PHONY: mock
mock:	
	mockgen -package servermock github.com/gufranmirza/imdb-api/web/server Server > web/server/servermock/mock.go
	mockgen -package routermock github.com/gufranmirza/imdb-api/web/router Router > web/router/routermock/mock.go
	mockgen -package healthmock github.com/gufranmirza/imdb-api/web/services/health Health > web/services/health/healthmock/mock.go

packages = \
	./config\
	./web/server\
	./web/router\
	./web/services/health\

.PHONY: test
test:
	@$(foreach package,$(packages), \
    		set -e; \
    		go test  $(package);)
