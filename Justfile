_default:
    @just --list

alias t := test
alias b := build

# run all tests
test:
    go test ./...

# build all modules
build:
    go build -v ./...

# run the ci suite
ci: build test
