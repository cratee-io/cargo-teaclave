# cargo-teaclave

A cargo plugin doing stuff targeting the [teaclave-sgx-sdk](https://github.com/apache/incubator-teaclave-sgx-sdk) project.

## Environment 
- cmake >= 3.10
- cargo = 1.44.0-nightly (6e07d2dfb 2020-03-31)

## Installation

```bash
go get -u -v github.com/sammyne/cargo-teaclave

# please update ~/.cargo/bin/ to your $CARGO_HOME if you have customized it
mv $(go env GOPATH)/bin/cargo-teaclave ~/.cargo/bin/
```
