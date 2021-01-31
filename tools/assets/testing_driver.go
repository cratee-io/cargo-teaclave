package assets

import (
	"github.com/cratee-io/cargo-teaclave/pkg/xpackr"
)

const testingDriverPath = "testing-driver"

func CopyTestingDriverTo(dst string) error {
	return xpackr.CopyDirFromBox(rootDir, testingDriverPath, dst)
}
