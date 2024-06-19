package path

import (
	"os/exec"
	"testing"
)

func TestExists(t *testing.T) {
	t.Run("Docker found in path", func(t *testing.T) {
		// Mocking exec.LookPath to return expected path
		// TODO: fix library selector error
		exec.LookPath = func(path string) (string, error) {
			return "/usr/bin/docker", nil
		}

		if _, err := path.Exists("docker"); err != nil {
			t.Errorf("path.Exists() returned false, expected true")
		}
	})

	t.Run("Docker not found in path", func(t *testing.T) {
		// Mocking exec.LookPath to return an error
		exec.LookPath = func(path string) (string, error) {
			return "", exec.ErrNotFound
		}

		if _, err := path.Exists("docker"); err == nil {
			t.Errorf("path.Exists() return true, expected false")
		}
	})

}
