# fog-wallet
OpenAPI based like mobilecoin full-service, but utilizing fog, but not really because I'm still learning how our clients work.

So really just an Go/HTTP/OpenAPI/ORM reference project.

What works:

- HTTP: using `echo`
- ORM: `gorm` with sqlite and auto-migrations, but everything _should_ work with postgres
- OpenAPI: Routes and models generation with `oapi-codegen`
- Dev experience: develop out of docker-compose with hot restarts and vscode remote debugging.

Still to do:

- Logging framework. Just using built in go `log`, could use full json logging. `echo` does json access/error logs by default.
- Might need some more robust database practices, limits/ranges on searches and stuff.
- Health check
- Metrics (Prometheus), Tracing (OpenTelemetry), Sentry 
- Figure out a whole lot about how to fog/mobilecoin. :)
- More documentation

## How to OpenAPI

### Resources

- [OpenAPI Specification](https://swagger.io/specification/)

### Adding/Modifying

1. Define your new API path and schema.
    One of the primary goals of OpenAPI is to encourage a "Design first" workflow.

    - Edit the `openapi/fog-wallet.yaml`.
    - Add new objects and save.
    - Generate new go boilerplate with `go generate -v api/api.go`
    - Rebuild `go build -v -o fog-wallet`

    If you're using the docker-compose environment, the dev container will do the generate and rebuild automatically on save.

### Swagger

Swagger is a UI for browsing the OpenAPI document. The included docker-compose starts up a swagger instance for browsing the API.

http://localhost:8081


## How to do Data Models

TODO

## Development

See [Setup](#go_setup) for requirements and creating a go development environment.

### Doing development

When you start work, source the go project `source_me.sh` file to setup your environment.

```
cd ~/gopath/things
source ./source_me.sh
```

Use `docker-compose` to build/run the app with sample DB and follow the logs:

```
cd ~/gopath/things/src/github.com/jgreat/things
docker-compose up --build
```

The app available on localhost.

* App endpoint: http://127.0.0.1:8080

When you make changes to the code the app container should automatically restart and rebuild the app binary.

### Debugging

When launched from `docker-compose` the app is run with headless `dlv` listening on `:2345`. 

There is a pre-configured `.vscode/launch.json` profile ready to attach to `dlv` to remote debug.

In `vscode` select the debug option and run `Attach Remote`, set your break points and have fun.

---

## Go Setup

### Prerequisites

This environment is pre-configured for running/compiling in docker with remote debugging and automatic rebuilds on code changes.

- `vscode` - https://code.visualstudio.com/download
- `go` 1.17 - https://golang.org/dl/
- `docker` - https://docs.docker.com/get-docker/
- `docker-compose` - https://docs.docker.com/compose/install/

### Install go

Download latest go 1.17 for your system: https://golang.org/dl/

Extract tar to `~/bin` (this will over wite the contents of the current `go`)

```
cd ~/bin
tar xvzf ~/Downloads/go1.17.3.linux-amd64.tar.gz
```

Move `go` to a versioned directory

```
mv go go-1.17.3
```

### Set up development environment

These instructions will help you create an isolated project path in you home directory.  More details on gopath setup can be found here: https://golang.org/doc/gopath_code.html

```
mkdir -p ~/gopath/things
cd ~/gopath/things
```

Add this script to your project base and point `GOROOT` at the version of go:

`source_me.sh`

```
#!/bin/bash

export GOROOT="${HOME}/bin/go-1.17.3"
export GOPATH="$(pwd)"
export PATH="${PATH}:${GOROOT}/bin:${GOPATH}/bin"
```

Create src path and clone the repo:

```
cd ~/gopath/fog-wallet
mkdir -p src/github.com/jgreat
cd src/github.com/jgreat
git clone git@github.com:jgreat/fog-wallet.git
```
