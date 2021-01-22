package assets

import (
	"os"
	"strings"

	"github.com/gobuffalo/packd"
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

func Path4EDL(tag string) (string, error) {
	cratePath := edlPathPrefix + tag
	if !RootDir.HasDir(cratePath) {
		return "", os.ErrNotExist
	}

	return cratePath, nil
}

func Tags4EDL() ([]string, error) {

	var tags []string
	tagged := make(map[string]bool)
	_ = RootDir.WalkPrefix(edlPathPrefix, func(filename string, _ packd.File) error {
		breadcrumbs := strings.SplitN(filename, string(os.PathSeparator), 3)
		if len(breadcrumbs) < 2 || tagged[breadcrumbs[1]] {
			return nil
		}
		tagged[breadcrumbs[1]] = true

		tags = append(tags, strings.TrimPrefix(breadcrumbs[1], edlCratePrefix))
		return nil
	})

	return tags, nil
}
