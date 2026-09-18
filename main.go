package main

import (
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
	"golang.org/x/mod/module"
)

func main() {
	var inp input

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Enter a go module name:").
				Value(&inp.moduleName).
				Validate(func(str string) error {
					if err := module.CheckPath(str); err != nil {
						return fmt.Errorf("invalid go module path: %w", err)
					}

					return nil
				}),
		),

		huh.NewGroup(
			huh.NewSelect[envLoader]().
				Title("load environment variables:").
				Options(
					huh.NewOption("none", noLoader),
					huh.NewOption(".env + github.com/joho/godotenv/tree/main (load .env variables into process environment)",
						godotenv),
					huh.NewOption(".env + dotenvx CLI (load .env without Go code dependencies)", dotenvx),
				).
				Value(&inp.env.loader),

			huh.NewSelect[envParser]().
				Title("parse environment variables:").
				Options(
					huh.NewOption("none", noParser),
					huh.NewOption("github.com/caarlos0/env/v11 (map environment variables into typed Go struct)", caarlosEnv),
				).
				Value(&inp.env.parser),

			huh.NewSelect[envSecretManager]().
				Title("manage secrets:").
				Options(
					huh.NewOption("none", noSecretManager),
					huh.NewOption("dotenvx (encrypt .env secrets while keeping .env workflow)", dotenvxSecrets),
					huh.NewOption("AWS Secrets Manager", awsSecretsManager),
					huh.NewOption("Google Cloud Secret Manager", googleCloudSecretManager),
					huh.NewOption("Azure Key Vault", azureKeyVault),
				).
				Value(&inp.env.secretManager),
		),
	)

	if err := form.Run(); err != nil {
		log.Fatal(err)
	}
}
