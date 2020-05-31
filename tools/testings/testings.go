package testings

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/sammyne/cargo-teaclave/pkg/cargo"

	"github.com/sammyne/cargo-teaclave/pkg/config"
	"github.com/sammyne/cargo-teaclave/pkg/xpackr"
)

const testedCrateName = "wheel"

func NewWorkspace(cratePath string) (string, error) {
	workingDir, err := ioutil.TempDir("", "teaclave-testing-driver-")
	if err != nil {
		return "", fmt.Errorf("fail to make a temporary working directory: %w", err)
	}

	if err := xpackr.CopyDirFromBox(config.AssetsDir, "testing-driver@v1.1.2",
		workingDir); err != nil {
		return "", fmt.Errorf("fail to set up testing driver: %w", err)
	}

	driverEnclaveCratePath := filepath.Join(workingDir, "enclave")
	if err := cargo.UseLocalCrate4Dependency(driverEnclaveCratePath, testedCrateName,
		cratePath); err != nil {
		return "", fmt.Errorf("fail to set up workspace to test '%s': %w", cratePath, err)
	}

	return workingDir, nil
}

func Run(workspace string) error {
	buildDir := filepath.Join(workspace, "build")
	if err := os.RemoveAll(buildDir); err != nil {
		return fmt.Errorf("fail to remove build dir: %w", err)
	}
	if err := os.MkdirAll(buildDir, 0644); err != nil {
		return fmt.Errorf("fail to make build dir: %w", err)
	}
	defer os.RemoveAll(buildDir)

	if err := runCommand(buildDir, "cmake", ".."); err != nil {
		return fmt.Errorf("fail to run 'cmake ..'; %w", err)
	}

	if err := runCommand(buildDir, "make"); err != nil {
		return fmt.Errorf("fail to run 'make'; %w", err)
	}

	if err := runCommand(buildDir, "make", "test-sgx"); err != nil {
		return fmt.Errorf("fail to run 'make test-sgx'; %w", err)
	}

	return nil
}

func runCommand(workingDir, cmd string, args ...string) error {
	c := exec.Command(cmd, args...)
	c.Dir, c.Stderr, c.Stdout = workingDir, os.Stderr, os.Stdout
	return c.Run()
}
