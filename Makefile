.PHONY: setup expenses-go-api test

all: expenses-go-api
# all: expenses-go-api go-swagger

# setup:
# 	yarn
# 	make expenses-go-api

devinit: expenses-go-api.go
	GO111MODULE=on go mod init github.com/marian0/expenses-go-api
	go mod tidy

expenses-go-api: expenses-go-api.go
	GO111MODULE=on go build expenses-go-api.go

# swaggerlocal: go-swagger redoc

# go-swagger:
# 	go build -o ./swagger -i github.com/go-swagger/go-swagger/cmd/swagger
# 	./swagger generate spec -o ./docs/swagger.json


test:
	go test -cover -count=1