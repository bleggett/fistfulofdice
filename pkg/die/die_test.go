package die

import (
	"testing"
	"sync"
)

func TestReturnsInt(t *testing.T) {
	var testWg sync.WaitGroup
	var tests = []struct {
		faces int
		testChan chan int
		testWg sync.WaitGroup
		fakeRand int
	}{
		{6, make(chan int, 1), testWg, 4},
		{6, make(chan int, 1), testWg, 0},
		{6, make(chan int, 1), testWg, 5},
		{8, make(chan int, 1), testWg, 4},
		{8, make(chan int, 1), testWg, 7},
		{8, make(chan int, 1), testWg, 6},
	}
	for _, c := range tests {
		c.testWg.Add(1)
		var limitArg int

		getRandInt = func(limit int) int {
			limitArg = limit
			return c.fakeRand
		}

		go Roll(c.faces, c.testChan, &c.testWg)
		c.testWg.Wait()
		res := <- c.testChan

		if res != (c.fakeRand + 1) {
			t.Errorf("Roll() did not return an integer! %d", res)
		}
		if limitArg != c.faces {
			t.Errorf("Roll() did not pass face value %d as upper limit to rand func (%d)", c.faces, limitArg)
		}
	}
}
