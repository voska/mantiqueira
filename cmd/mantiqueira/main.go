package main

import (
	"github.com/voska/mantiqueira"
	"github.com/voska/vtexkit/cli"
)

// version is injected at build time via -ldflags.
var version = "dev"

func main() {
	cli.Main(cli.App{
		Store:       mantiqueira.Store,
		Version:     version,
		Description: "Mantiqueira em Casa egg CLI for humans and AI agents.",
	})
}
