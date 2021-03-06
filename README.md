# cargo-teaclave

![minimum go](https://img.shields.io/badge/go-1.15%2B-blue)
![rustc](https://img.shields.io/badge/rustc-ffa2e7ae8%202020--10--24-blue)
![cmake](https://img.shields.io/badge/cmake-3.10%2B-blue)
![test-hello-world](https://github.com/cratee-io/cargo-teaclave/workflows/test-hello-world/badge.svg)

A cargo plugin serves to ease developing apps with the [teaclave-sgx-sdk][2] project.

## Getting Started
These instructions will get you a copy of the project up and running on your local machine for
development and testing purposes.

### Prerequisites
- cmake 3.10+
- rustc 1.49.0-nightly (ffa2e7ae8 2020-10-24)
- go    1.14+

### Install

```bash
go get -u -v github.com/cratee-io/cargo-teaclave

# please update ~/.cargo/bin/ to your $CARGO_HOME if you have customized it
mv $(go env GOPATH)/bin/cargo-teaclave ~/.cargo/bin/
```

### Run
```bash
cargo teaclave
```

Just follow the hint shown as
```bash
teaclave helps to plays with apps written with teaclave-sgx-sdk

Usage:
  teaclave [command]

Available Commands:
  help        Help about any command
  test        test a given teaclave-sgx-sdk-ported crate

Flags:
  -h, --help   help for teaclave

Use "teaclave [command] --help" for more information about a command.
```

## Examples

example | description
-------:|:-------------
[test-hello-world][4] | test a teaclave-sgx-sdk-ported crate

## Recommend
During our daily development with [teaclave-sgx-sdk][2], we found it requires rigirous setting for
its SGX SDK. It takes non-trivial works to set up the SGX SDK and rust toolchain. Therefore, it's
recommended to just develop within the environment packaged by the official docker container (e.g.
sammyne/rsgx-dcap:2.12.100.3-dcap1.9.100.3-rs20201025-go1.15.7-ubuntu18.04 if you're playing with teaclave-sgx-sdk@v1.1.3).

## References 
- [README-Template][1]

[1]: https://gist.github.com/PurpleBooth/109311bb0361f32d87a2
[2]: https://github.com/apache/incubator-teaclave-sgx-sdk
[4]: ./examples/test-hello-world/README.md