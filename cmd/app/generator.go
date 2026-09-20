package main

import (
	"fmt"
	"os/exec"
)

func generate(inp input) error {
	// generate go.mod
	cmd := exec.Command("go", "mod", "init", inp.moduleName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to generate go.mod: %w, output: %s", err, output)
	}

	return nil
}
