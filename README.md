# GOLANG FIBER API (CLEAN ARCHITECTURE)
**_Best simple, lightweight, powerful and really fast Api with Golang (Fiber, REL, Dbmate) PostgreSqLDatabase using Clean Architecture_**

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
It is created with golang (fiber and REL)

## how to run

**golderserver** requires [go 1.17+](https://golang.org/dl/) 
***postgres and dbmate*** installed to run.

create .env file, copy .envexample content and set your postgres variables
```sh
touch .env
```

Creating database:

to create database in postgres _check "infraestructure/database/base" folder_

```sh
dbmate up
```

Install the dependencies and devDependencies and start the server.

```sh
make download
make run
```

then, check 

```sh
127.0.0.1:3000
```

## License
MIT

