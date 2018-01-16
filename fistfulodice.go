package main

import (
	"fmt"
	"github.com/bleggett/fistfulofdice/pkg/die"
	"os"
	"strconv"
	"sync"
	"time"
	"github.com/briandowns/spinner"
)

func main() {

	if len(os.Args) > 1 {
		count, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("Error parsing CLI: %s", err)
		}
		c := make(chan int, count)
		var wg sync.WaitGroup
		wg.Add(count)
		fmt.Printf("Chucking a fist of %d di(c)e!\n", count)
		spin := spinner.New(spinner.CharSets[4], 100* time.Millisecond)
		spin.Start()
		for i := 0; i < count; i++ {
			go die.Roll(6, c, &wg)
		}
		//If we used `range` here, we'd have to close the channel in Roll,
		//but Roll doesn't know how many times it is called, so can't close there.
		wg.Wait()
		spin.Stop()
		fmt.Print("\n")
		fmt.Print("Result: ")
		for j := 0; j < count; j++ {
			fmt.Printf("%d ", <- c)
		}
	} else {
		fmt.Print("You must specify a number of dice to roll...\n")
	}
}
