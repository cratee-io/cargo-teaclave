package testings

import (
	"fmt"
	"io/ioutil"

	"github.com/cratee-io/cargo-teaclave/pkg/cargo"
)

const (
	sdkFromApache = "https://github.com/apache/teaclave-sgx-sdk"
	sdkFromCratee = "https://github.com/cratee-io/incubator-teaclave-sgx-sdk"
)

func newCargoWorkspaceManifest(outPath, driverTag string) error {
	members := []string{"app", "enclave"}

	sdkPatch := cargo.Patch{
		Old:    sdkFromApache,
		NewGit: sdkFromCratee,
		Crates: make(map[string]string),
	}

	// @TODO: make this configurable
	crates := []string{
		"sgx_alloc", "sgx_libc", "sgx_rand", "sgx_tcrypto", "sgx_tdh", "sgx_tcrypto_helper",
		"sgx_tkey_exchange", "sgx_tprotected_fs", "sgx_trts", "sgx_tse", "sgx_tseal", "sgx_tstd",
		"sgx_types", "sgx_ucrypto", "sgx_urts",
	}
	for _, v := range crates {
		sdkPatch.Crates[v] = "v" + driverTag
	}

	workspace, err := cargo.NewWorkspaceManifest(members, sdkPatch)
	if err != nil {
		return fmt.Errorf("new manifest: %w", err)
	}

	return ioutil.WriteFile(outPath, workspace, 0644)
}
