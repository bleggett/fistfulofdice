package main

import (
	"fmt"
	"github.com/bleggett/fistfulofdice/pkg/die"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) > 1 {
		count, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("Error parsing CLI: %s", err)
		}
		c := make(chan int, count)
		for i := 0; i < count; i++ {
			die.Roll(6, c)
			fmt.Printf("Result: %d\n", <- c)
		}
	} else {
		fmt.Print("You must specify a number of dice to roll...\n")
	}
}
