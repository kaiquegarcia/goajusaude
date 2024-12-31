package main

import (
	"goajusaude/cmd"
	"goajusaude/pkg/shared"
	"goajusaude/pkg/shared/ajusaude"
	"os"
)

func main() {
	ajuSaudeClient := ajusaude.New(
		shared.NewHttpClient("https://aracaju-saude.voipy.com.br/portalcidadaoapi/portalcidadao"),
	)

	command := "cli"
	args := make([]string, 0)
	if len(os.Args) > 1 {
		if os.Args[1] == "bat" {
			command = "bat"
			if len(os.Args) > 2 {
				args = os.Args[2:]
			}
		}
	}

	switch command {
	case "cli":
		cmd.RunCLI(ajuSaudeClient)
	case "bat":
		cmd.RunBat(ajuSaudeClient, args)
	}
}
