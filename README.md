# codeid-boiler
This project is code.id boilerplate to create API in Golang Echo Framework
- PORT : 3030
- PATH : /

## Installation

``` bash
# clone the repo
$ git clone 

# go into app's directory
$ cd my-project
```

## Build & Run

Local environment
``` bash
# build 
$ go build

# run in development 
$ ENV=DEV go run main.go
$ ENV=DEV ./filego

# run in staging 
$ ENV=STAGING go run main.go
$ ENV=STAGING ./filego

# run in production 
$ ENV=PROD go run main.go
$ ENV=PROD ./filego
```

Docker environment
``` bash
# build 
$ docker build -t codeid-api:latest .

# config
sudo sysctl -w vm.max_map_count=262144 # it is required for elasticsearch config

# run
$ docker compose up
```

## Documentation

Install environment
``` bash
# get swagger package 
$ go get github.com/swaggo/swag

# move to swagger dir
$ cd $GOPATH/src/github.com/swaggo/swag

# install swagger cmd 
$ go install cmd/swag
```

Generate documentation
``` bash
# generate swagger doc
$ swag init --propertyStrategy snakecase
```
to see the results, run app and access {{url}}/swagger/index.html

## Description 
This project built in clean architecture that contains :
1. Http
2. Factory
3. Middleware 
4. Handler
5. Binder
6. Validation
7. Service
8. Repository
9. Model
10. Database

This project have some default function :
- Context
- Validator
- Transaction
- Pagination & Sort
- Filter
- Env
- Response
- Redis
- Elasticsearch
- Log

This project have some default endpoint :
- Auth 
  - Login
  - Register
- Sample
  - Get (with pagination, sort, & filter)
  - GetByID
  - Create (with transaction)
  - Update (with transaction)
  - Delete

# Author
CodeID Backend Team