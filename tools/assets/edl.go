package assets

import (
	"fmt"
	"os"
	"strings"

	"github.com/cratee-io/cargo-teaclave/pkg/xpackr"
	"github.com/gobuffalo/packd"
)

func CopyEDLsTo(dst, rev string) error {
	cratePath, err := Path4EDL(rev)
	if err != nil {
		tags, _ := Tags4EDL()
		return fmt.Errorf("rev(%s) is none of %v", rev, tags)
	}

	return xpackr.CopyDirFromBox(RootDir, cratePath, dst)
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
