package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	inputer := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		inputer.Scan()
		text := inputer.Text()

		fmt.Println("Echoing: ", text)
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
