# diatom-pub

POC srv

### Main functionality:

- webservice support with middleware architecture
- redis cache support
- golang implementation with modules

### More technical features:

- graceful shutdown
- semver versioning
- build information with git version tag, commit and date
- detailed logging
- deep healthcheck
- start time < 1s
- docker size ~13MB

Check Makefiles for all important stuff.

## Project structure

- [./cmd/srv](./cmd/srv) - server implementation
- [./internal](./internal) - internal packages that should not be shared with other projects
- [./template](./template/README.md) - html templates

## Requirements

- [golang](https://golang.org/doc/install) installation
- gui editor, eg [goland](https://www.jetbrains.com/go)

## API doc

Application serves
 
1. [api.yaml](./template/api.yaml) - oas3 definition of endpoints.

## DB

Application uses postgres database with [SQLC](https://docs.sqlc.dev/en/stable/index.html).

## Logging

Project is using standard logger from `log` library. It is configured in `main.go` and should be used in all logging
statements. Log is in format like: 

**NAME** : (**VERSION**, **SHA**) : **DATE-TIME** **FILE**: **MSG**

where:

- **NAME**: microservice name
- **VERSION**: [semver](https://semver.org/) version taken from annotated tag, `dev` otherwise
- **SHA**: git SHA in short version
- **DATE-TIME**: date time with microseconds
- **FILE**: source file name and line information
- **MSG**: log message

### Profiler

To use memory profiler use `curl -sK -v http://localhost:8080/debug/pprof/heap > heap.out` and `go tool pprof heap.out` or 
with interactive GUI `go tool pprof -http=:8080 heap.out`


### NOTES

```shell
# for testing full no downtime deployment
export ENDPOINT=https://1.2.3.4
while true; do echo "$(date) - $(curl -s -o /dev/null -w "%{http_code}\n" ${ENDPOINT}/api.yaml)"; sleep 0.1; done
```

### Links

- https://github.com/kyleconroy/sqlc
- https://www.youtube.com/c/TECHSCHOOLGURU/videos
- [JWT and PASETO](https://www.youtube.com/watch?v=Oi4FHDGILuY)
- [Go Time](https://podcasts.apple.com/no/podcast/go-time-golang-software-engineering/id1120964487)

License
-------
[![License: MIT](https://img.shields.io/badge/License-mit-brightgreen.svg)](https://opensource.org/licenses/MIT)
