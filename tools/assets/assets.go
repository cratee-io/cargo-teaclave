package assets

import (
	"github.com/gobuffalo/packr/v2"
)

const (
	edlCratePrefix = "sgx_edl-"
	edlPathPrefix  = "rsgx-edls/" + edlCratePrefix
)

var RootDir = packr.New("assets", "../../assets")

// Has check if there is a asset within the RootDir
func Has(asset string) bool {
	return RootDir.HasDir(asset)
}
