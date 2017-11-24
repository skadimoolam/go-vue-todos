# Simple Todos using GoLang and Vue
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fskadimoolam%2Fgo-vue-todos.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fskadimoolam%2Fgo-vue-todos?ref=badge_shield)


## Prerequisites

- Install [Go](https://golang.org).
- Install [Glide](https://glide.sh/). (If you do not use Glide check the Manual setup at the end)

## Go get project

```
$ go get -u github.com/skadimoolam/go-vue-todos
$ cd $GOPATH/src/github.com/skadimoolam/go-vue-todos
```

## Install dependencies

```
$ glide install
```

## Start app
```
$ go run todo.go
```

## Open app

Open `localhost:8080` in your browser.

## Manual setup without using Glide
 - `go get github.com/skadimoolam/go-vue-todos`
 - `go get -u github.com/labstack/echo` -- This installs ECHO
 - `go get -u github.com/mattn/go-sqlite3` -- This installs the sqlite package
 - `cd $GOPATH/src/github.com/skadimoolam/go-vue-todos`
 - `go run todo.go`  -- assuming you have Go already installed and configured
 - Open `localhost:8080` in your browser


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fskadimoolam%2Fgo-vue-todos.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fskadimoolam%2Fgo-vue-todos?ref=badge_large)