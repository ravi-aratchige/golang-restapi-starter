# Golang RestAPI Starter

<img src="./docs/img/golang.jpg" alt="Golang Gopher">

This is my take on a boilerplate for a simple RestAPI built with Go (Golang).

This project uses the built-in <a href="https://pkg.go.dev/net/http">`net/http`</a> Golang package for the webserver and the <a href="https://gorm.io/">GORM</a> 🦦 library for database interaction.

This project also includes a demo SQLite database with mock data for you to try out yourself.

## Setup

Make sure you have <a href="https://go.dev/doc/install">Go</a> installed.

### Linux / MacOS Systems

Install all dependencies:

```shell
make install
```

Start the server:

```shell
make run
```

### Windows Systems

Install all dependencies:

```shell
go get .
```

Start the server:

```shell
go run .
```

---

Made with ❤️ by Ravindu Aratchige. Licensed under the <a href="https://github.com/ravi-aratchige/golang-restapi-starter/blob/main/LICENSE">Apache License<a>.
