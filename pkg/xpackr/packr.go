package xpackr

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/gobuffalo/packr/v2"
)

func CopyDirFromBox(b *packr.Box, from, to string) error {
	return b.WalkPrefix(from, func(filename string, src packr.File) error {
		info, err := src.FileInfo()
		if err != nil {
			return fmt.Errorf("fail to read file info for '%s': %w", filename, err)
		}

		data, err := ioutil.ReadAll(src)
		if err != nil {
			return fmt.Errorf("fail to read '%s': %w", filename, err)
		}

		dstDir := strings.Replace(path.Dir(filename), from, to, 1)
		if err := os.MkdirAll(dstDir, 0644); err != nil {
			return fmt.Errorf("fail to mkdir destination '%s': %w", dstDir, err)
		}

		dst := filepath.Join(dstDir, path.Base(filename))
		return ioutil.WriteFile(dst, data, info.Mode())
	})
}
