package main

import (
	"github.com/replicate/cog/pkg/cli"
	"github.com/replicate/cog/pkg/util/console"
)

func main() {
	installDockerStr := "Cog uses Docker to create a container for your model. You'll need to install Docker before you can run Cog. If you install Docker Engine instead of Docker Desktop, you will need to install Buildx as well."
	err := cli.checkDockerInstallation()
	if err != nil {
		console.Errorf("%f", err)
		console.Fatalf(installDockerStr)
	}
	cmd, err := cli.NewRootCommand()
	if err != nil {
		console.Fatalf("%f", err)
	}

	if err = cmd.Execute(); err != nil {
		console.Fatalf("%s", err)
	}
}
