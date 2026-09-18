package main

import (
	"log"

	"github.com/charmbracelet/huh"
)

func main() {
	form := huh.NewForm()

	if err := form.Run(); err != nil {
		log.Fatal(err)
	}
}
