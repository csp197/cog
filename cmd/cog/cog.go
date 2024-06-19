package main

import (
	"github.com/replicate/cog/pkg/cli"
	"github.com/replicate/cog/pkg/util/console"
	"github.com/replicate/cog/pkg/util/path"
)

func main() {
	if _, err := path.Exists("docker"); err != nil {
		console.Fatalf("%f", err)
	}
	cmd, err := cli.NewRootCommand()
	if err != nil {
		console.Fatalf("%f", err)
	}

	if err = cmd.Execute(); err != nil {
		console.Fatalf("%s", err)
	}
}
