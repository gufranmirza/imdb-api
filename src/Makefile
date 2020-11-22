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
	mockgen -package jwtmock github.com/gufranmirza/imdb-api/auth/jwt TokenAuth > mocks/jwtmock/mock.go
	mockgen -package servermock github.com/gufranmirza/imdb-api/web/server Server > mocks/servermock/mock.go
	mockgen -package routermock github.com/gufranmirza/imdb-api/web/router Router > mocks/routermock/mock.go
	mockgen -package healthmock github.com/gufranmirza/imdb-api/web/services/health Health > mocks/healthmock/mock.go
	mockgen -package movieservicemock github.com/gufranmirza/imdb-api/web/services/v1/movieservice MovieService > mocks/movieservicemock/mock.go
	mockgen -package authservicemock github.com/gufranmirza/imdb-api/web/services/v1/authservice AuthService > mocks/authservicemock/mock.go
	mockgen -package moviedalmock github.com/gufranmirza/imdb-api/dal/moviedal MovieDal > mocks/moviedalmock/mock.go
	mockgen -package userdalmock github.com/gufranmirza/imdb-api/dal/userdal UserDal > mocks/userdalmock/mock.go
	mockgen -package tokendalmock github.com/gufranmirza/imdb-api/dal/tokendal TokenDal > mocks/tokendalmock/mock.go

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
api-docs:
	export PATH=$(go env GOPATH)/bin:$PATH
	swag init --output=./web/docs

migrate-up:
	../bin/${BINARY_NAME} migrate up 

migrate-down:
	../bin/${BINARY_NAME} migrate down

migrate-new:
	../bin/${BINARY_NAME} migrate new "new migration script"

.PHONY: dataloader
dataloader:
	../bin/${BINARY_NAME} dataloader ./data/imdb.json