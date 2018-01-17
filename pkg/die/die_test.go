package die

import (
	"sync"
	"testing"
)

func TestReturnsInt(t *testing.T) {
	var testWg sync.WaitGroup
	var testChan = make(chan int)
	var tests = []struct {
		faces    int
		fakeRand int
	}{
		{6, 4},
		{6, 0},
		{6, 5},
		{8, 4},
		{8, 7},
		{8, 6},
	}
	for _, c := range tests {
		testWg.Add(1)
		var limitArg int

		getRandInt = func(limit int) int {
			limitArg = limit
			return c.fakeRand
		}

		go Roll(c.faces, testChan, &testWg)
		res := <-testChan

		if res != (c.fakeRand + 1) {
			t.Errorf("Roll() did not return an integer! %d", res)
		}
		if limitArg != c.faces {
			t.Errorf("Roll() did not pass face value %d as upper limit to rand func (%d)", c.faces, limitArg)
		}
	}
}
