package testings

import (
	"fmt"
	"path/filepath"

	"github.com/cratee-io/cargo-teaclave/pkg/xpackr"
	"github.com/cratee-io/cargo-teaclave/tools/assets"
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
