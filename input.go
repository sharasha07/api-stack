package main

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"golang.org/x/mod/module"
)

type input struct {
	moduleName   string
	wantMakefile bool

	env struct {
		loader envLoader
		parser envParser
	}
}

func getInput(inp *input) *huh.Form {
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
			huh.NewConfirm().
				Title("Add a Makefile for common commands?").
				Value(&inp.wantMakefile),
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
		),
	)

	return form
}
