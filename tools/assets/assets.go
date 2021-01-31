package assets

import (
	"github.com/gobuffalo/packr/v2"
)

const (
	edlCratePrefix = "sgx_edl-"
	edlPathPrefix  = "rsgx-edls/" + edlCratePrefix
)

var rootDir = packr.New("assets", "../../assets")
