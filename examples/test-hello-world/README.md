# test-hello-world

This example serves to demonstrate how to test a teaclave-sgx-sdk-ported with cargo-teaclave
plugin.

## Getting Started

### Run
```bash
docker run -it -v ${PWD}:/workspace -w /workspace \
  sammyne/rsgx-dcap:2.12.100.3-dcap1.9.100.3-rs20201025-go1.15.7-ubuntu18.04 bash

# from now on, we're within the container

cargo teaclave test
```

Once succeeded, we should see output like

```bash
running 1 tests
test test_hello_world::tests::it_works ... ok

test result ok. 1 passed; 0 failed; 0 ignored
```
