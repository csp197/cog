package cli

import (
	"os/exec"
)

func checkDockerInstallation() error {
	_, err := exec.LookPath("docker")
	if err != nil {
		return err
	}
	return nil
}
