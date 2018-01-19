package main

import (
	"fmt"
	"github.com/bleggett/fistfulofdice/internal/pkg/die"
	"github.com/briandowns/spinner"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	if len(os.Args) > 1 {
		diestrings := strings.Split(os.Args[1], "d")
		if len(diestrings) < 2 {
			fmt.Printf("Error parsing CLI. Use format XdX")
			return
		}
		count, err := strconv.Atoi(diestrings[0])
		faces, err := strconv.Atoi(diestrings[1])
		if err != nil {
			fmt.Printf("Error parsing CLI: %s\nUse format XdX", err)
			return
		}
		c := make(chan int, count)
		var wg sync.WaitGroup
		wg.Add(count)
		fmt.Printf("Chucking a fist of %d di(c)e!\n", count)
		spin := spinner.New(spinner.CharSets[4], 100*time.Millisecond)
		spin.Start()
		for i := 0; i < count; i++ {
			go die.Roll(faces, c, &wg)
		}
		//If we used `range` here, we'd have to close the channel in Roll,
		//but Roll doesn't know how many times it is called, so can't close there.
		wg.Wait()
		spin.Stop()
		fmt.Print("\nResult: ")
		for j := 0; j < count; j++ {
			fmt.Printf("%d ", <-c)
		}
	} else {
		fmt.Print("You must specify a number of dice to roll in the XdX format...\n")
	}
}
