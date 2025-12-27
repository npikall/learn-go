# Learn go with tests

[![Go Test](https://github.com/npikall/learn-go/actions/workflows/ci.yml/badge.svg)](https://github.com/npikall/learn-go/actions/workflows/ci.yml)

This is the implementation of the guide [learn go with tests]

## Useful Commands

```bash
# to run the tests
go test (-v)

# run tests with code coverage
go test -cover

# to run benchmarks (with information about memory allocation [-benchmem])
go test -bench=. (--benchmem)

# running something like a linter
go vet
```

![Gopher](https://raw.githubusercontent.com/egonelbre/gophers/master/.thumb/vector/projects/network-side.png)

[learn go with tests]: https://quii.gitbook.io/learn-go-with-tests
