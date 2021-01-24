package testings

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/cratee-io/cargo-teaclave/pkg/cargo"
	"github.com/cratee-io/cargo-teaclave/pkg/xpackr"
	"github.com/cratee-io/cargo-teaclave/tools/assets"
)

const (
	sdkFromApache = "https://github.com/apache/teaclave-sgx-sdk"
	sdkFromCratee = "https://github.com/cratee-io/incubator-teaclave-sgx-sdk"
)

func addEDL4Driver(workdir, tag string) error {
	cratePath, err := assets.Path4EDL(tag)
	if err != nil {
		tags, _ := assets.Tags4EDL()
		return fmt.Errorf("tag(%s) is none of %v", tag, tags)
	}

	outPath := filepath.Join(workdir, "third_party", "rsgx-assets", "vendor", "sgx_edl")
	if err := xpackr.CopyDirFromBox(assets.RootDir, cratePath, outPath); err != nil {
		return fmt.Errorf("renew EDLs: %w", err)
	}

	return nil
}

func newCargoWorkspaceManifest(outPath, driverTag string) error {
	members := []string{"app", "enclave"}

	sdkPatch := cargo.Patch{
		Old:    sdkFromApache,
		NewGit: sdkFromCratee,
		Crates: make(map[string]string),
	}
	crates := []string{"sgx_tstd", "sgx_types"} // @TODO: make this configurable
	for _, v := range crates {
		sdkPatch.Crates[v] = "v" + driverTag
	}

	workspace, err := cargo.NewWorkspaceManifest(members, sdkPatch)
	if err != nil {
		return fmt.Errorf("new manifest: %w", err)
	}

	return ioutil.WriteFile(outPath, workspace, 0644)
}
