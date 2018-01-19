package die

import (
	"math/rand"
	"sync"
	"time"
)

// Roll a die
func Roll(faces int, c chan int, wg *sync.WaitGroup) {
	result := getRandInt(faces) + 1
	time.Sleep(time.Second)
	c <- result
	defer wg.Done()
}

//Make this a mockable private member.
var getRandInt = func(limit int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(limit)
}
