package main

import (
	"bufio"
	"fmt"
	"os"
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
