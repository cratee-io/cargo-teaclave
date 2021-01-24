package testings

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/cratee-io/cargo-teaclave/pkg/assets"
	"github.com/cratee-io/cargo-teaclave/pkg/cargo"
	"github.com/cratee-io/cargo-teaclave/pkg/xpackr"
)

const (
	testedCrateName = "wheel"
	testingDriver   = "testing-driver"
)

// NewWorkspace sets up a temporary workspace to test the crate at cratePath with
// the driver tagged by testDriverTag
func NewWorkspace(cratePath, testingDriverTag string) (string, error) {
	workingDir, err := ioutil.TempDir("", "teaclave-testing-driver-")
	if err != nil {
		return "", fmt.Errorf("fail to make a temporary working directory: %w", err)
	}

	testingDriverPath := fmt.Sprintf("%s@%s", testingDriver, testingDriverTag)
	if !assets.Has(testingDriverPath) {
		tags, _ := assets.Tags(testingDriver)
		return "", fmt.Errorf("only drivers tagged by %v is supported, but got %s", tags,
			testingDriverTag)
	}

	if err := xpackr.CopyDirFromBox(assets.RootDir, testingDriverPath, workingDir); err != nil {
		return "", fmt.Errorf("fail to set up testing driver: %w", err)
	}

	driverEnclaveCratePath := filepath.Join(workingDir, "enclave")
	if err := cargo.UseLocalCrate4Dependency(driverEnclaveCratePath, testedCrateName,
		cratePath); err != nil {
		return "", fmt.Errorf("fail to set up workspace to test '%s': %w", cratePath, err)
	}

	return workingDir, nil
}

func NewWorkspaceV2(cratePath, driverTag string) (string, error) {
	workingDir, err := ioutil.TempDir("", "teaclave-testing-driver-")
	if err != nil {
		return "", fmt.Errorf("fail to make a temporary working directory: %w", err)
	}

	const driverPath = "testing-driver"
	if err := xpackr.CopyDirFromBox(assets.RootDir, driverPath, workingDir); err != nil {
		return "", fmt.Errorf("new driver project: %w", err)
	}

	if err := addEDL4Driver(workingDir, driverTag); err != nil {
		return "", fmt.Errorf("renew EDLs: %w", err)
	}

	driverEnclaveCratePath := filepath.Join(workingDir, "enclave")
	if err := cargo.UseLocalCrate4Dependency(driverEnclaveCratePath, testedCrateName,
		cratePath); err != nil {
		return "", fmt.Errorf("set up workspace to test '%s': %w", cratePath, err)
	}

	workspaceManifestPath := filepath.Join(workingDir, "Cargo.toml")
	if err := newCargoWorkspaceManifest(workspaceManifestPath, driverTag); err != nil {
		return "", fmt.Errorf("new cargo workspace manifest: %w", err)
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

func init() {
	if _, err := assets.Tags(testingDriver); err != nil {
		panic(fmt.Sprintf("%v", err))
	}
}

func runCommand(workingDir, cmd string, args ...string) error {
	c := exec.Command(cmd, args...)
	c.Dir, c.Stderr, c.Stdout = workingDir, os.Stderr, os.Stdout
	return c.Run()
}
