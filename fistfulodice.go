package main

import (
	"fmt"
	"github.com/bleggett/fistfulofdice/pkg/die"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) > 1 {
		fmt.Printf("Rolling...\n")
		count, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("Error parsing CLI: %s", err)
		}
		for i := 0; i < count; i++ {
			die.Roll(6)
		}
	} else {
		fmt.Print("You must specify a number of dice to roll...\n")
	}
}
