# API-Server

>Simple api server for studying

## Feature

- [x] [dev env set](doc/dev_env_set.md)
- [x] [base dir structure](doc/base_dir_structure.md)
- [x] import cli & config pkg
- [x] Makefile with partial build arguments
    - cross compile
    - mod: `go build -mod=vendor`
    - tags
- [ ] config option pattern
- [ ] simple http server
    - [ ] data structure: byte, int, float, string, rune, slice, map, chan, etc.
    - [ ] field tag
    - [ ] byte alignment
    - [ ] struct & interface & reflect(option)
    - [ ] for & for range & goto label
    - [ ] goroutine
    - [ ] closure
    - [ ] internal dir
    - [ ] related specification
    - [ ] http server
    - [ ] test & benchmark
- [ ] go cmd: build & run & others
- [ ] vendor
- [ ] add version info
    - [ ] code version
    - [ ] API version: tag/branch/version file
- [ ] proto
    - [ ] idl define
    - [ ] proto generate
- [ ] simple rpc server & client
- [ ] Dockerfile
- [ ] multi env support

## TODO

- [ ] restart & hot reload
- [ ] swagger api
- [ ] chan pattern
- [ ] metrics
    - [ ] grafana & prometheus
    - [ ] metrics
    - [ ] metrics api: new port
- [ ] JWT(JSON Web Token)
- [ ] http ratelimit & interceptor
- [ ] Web ASM
- [ ] design pattern
    - [ ] sigleton
