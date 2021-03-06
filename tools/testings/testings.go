package testings

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/cratee-io/cargo-teaclave/pkg/cargo"
	assetsPkg "github.com/cratee-io/cargo-teaclave/tools/assets"
)

const testedCrateName = "wheel"

// NewWorkspace sets up a temporary workspace to test the crate at cratePath with
// the driver tagged by driverTag
func NewWorkspace(cratePath, driverTag string) (string, error) {
	workingDir, err := ioutil.TempDir("", "teaclave-testing-driver-")
	if err != nil {
		return "", fmt.Errorf("fail to make a temporary working directory: %w", err)
	}

	if err := assetsPkg.CopyTestingDriverTo(workingDir); err != nil {
		return "", fmt.Errorf("new driver project: %w", err)
	}

	edlsPath := filepath.Join(workingDir, "third_party", "rsgx-assets", "vendor", "sgx_edl")
	if err := assetsPkg.CopyEDLsTo(edlsPath, driverTag); err != nil {
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
	if _, err := assetsPkg.Tags4EDL(); err != nil {
		panic(fmt.Sprintf("no availabble EDLs: %v", err))
	}
}

func runCommand(workingDir, cmd string, args ...string) error {
	c := exec.Command(cmd, args...)
	c.Dir, c.Stderr, c.Stdout = workingDir, os.Stderr, os.Stdout
	return c.Run()
}
