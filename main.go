package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/huh"
)

func main() {
	var inp input

	if err := getInput(&inp).Run(); err != nil {
		log.Fatal(err)
	}

	for {
		err := generate(inp)
		if err == nil {
			fmt.Println("Successfully generated! Check it out!")
			return
		}

		fmt.Fprintln(os.Stderr, err)

		retry, err := confirmRetry()
		if err != nil {
			log.Fatal(err)
		}

		if !retry {
			return
		}
	}
}

func confirmRetry() (bool, error) {
	var retry bool

	err := huh.NewConfirm().
		Title("Try generating again?").
		Value(&retry).
		Run()

	return retry, err
}
