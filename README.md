# GOLANG FIBER API (CLEAN ARCHITECTURE)
**_Best simple, lightweight, powerful and really fast Api with Golang (Fiber, Sqlc, Dbmate) PostgreSqL Database using Clean Architecture_**

status: **WIP**

LAYERS

- **App:**
  this directory has main setup files
- **Interface**
    - route handler here
    - midlewares
    - presenters
    - response configuration
    - validators
- **use case**
  this directory has application use cases (controllers)
- **infraestructure**
  this directory has database configuration, schemas, and repositories
- **Domain folder**
  domain folder contain entities and rules

## about api
API of an order flow inventory system built for the "**Versatil** Print shop".

## about frameworks
It is created with golang (fiber,Sqlc)

## how to run

**golderserver** requires [go 1.17+](https://golang.org/dl/) 
***postgres and dbmate*** installed to run.

**Step 1** configure env file

- create .env file
- copy .envexample content
- set your postgresql variables

```sh
touch .env
```

**Step 2** create database

to create database in postgres _check "infraestructure/database/base" folder_

```sh
dbmate up
```

**Step 3** Install sqlc cli


```sh
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
```

**Step 4** Sqlc generate model
_check sqlc.yaml file_

```sh
sqlc generate
```

**Step 4** Install the dependencies and start the server.

```sh
make download
make run
```

then, check 

```sh
127.0.0.1:3000
```

## License
APACHE LICENSE 2.0

