package assets

import (
	"errors"
	"strings"

	"github.com/gobuffalo/packr/v2"
)

var RootDir = packr.New("assets", "../../assets")

// Has check if there is a asset within the RootDir
func Has(asset string) bool {
	return RootDir.HasDir(asset)
}

func Tags(asset string) ([]string, error) {
	all := RootDir.List()

	var tags []string
	tagged := make(map[string]struct{})
	for _, v := range all {
		if !strings.HasPrefix(v, asset) {
			continue
		}
		if _, ok := tagged[v]; ok {
			continue
		}
		tagged[v] = struct{}{}

		assetAndTag := strings.SplitN(v, "@", 2)
		if len(assetAndTag) == 1 {
			tags = append(tags, "unknown")
		} else {
			tags = append(tags, assetAndTag[1])
		}
	}

	if len(tags) == 0 {
		return nil, errors.New("no such asset")
	}

	return tags, nil
}
