# test-hello-world

This example serves to demonstrate how to test a teaclave-sgx-sdk-ported with cargo-teaclave
plugin.

## Getting Started

### Run
```bash
docker run -it -v ${PWD}:/workspace -w /workspace baiduxlab/sgx-rust:1804-1.1.2 bash

# from now on, we're within the container
# install go
export GO_VERSION=1.14.3                                     
wget "https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz"
tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz         
ln -sf /usr/local/go/bin/* /usr/bin                               
go version
go env -w GO111MODULE=on
go get -u -v github.com/sammyne/cargo-teaclave@master
mv $(go env GOPATH)/bin/cargo-teaclave ~/.cargo/bin/

cargo teaclave test
```

Once succeeded, we should see output like

```bash
running 1 tests
test test_hello_world::tests::it_works ... ok

test result ok. 1 passed; 0 failed; 0 ignored
```
