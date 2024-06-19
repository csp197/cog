package path

import (
	"fmt"
	"os/exec"
)

func Exists(path string) (bool, error) {
	_, err := exec.LookPath(path)
	if err == nil {
		return true, nil
	}
	return false, fmt.Errorf("Failed to determine if %s exists: %w\nCog uses Docker to create a container for your model.\nYou'll need to install Docker before you can run Cog.\nIf you install Docker Engine instead of Docker Desktop, you will need to install Buildx as well.", path, err)
}
