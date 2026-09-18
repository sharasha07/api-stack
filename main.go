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
	)

	if err := form.Run(); err != nil {
		log.Fatal(err)
	}
}
