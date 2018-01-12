package die

import (
	"fmt"
	"math/rand"
	"time"
)

// Roll a die
func Roll(faces int) int {
	rand.Seed(time.Now().UnixNano())
	result := rand.Intn(faces) + 1
	fmt.Printf("..a die! Result %d\n", result)
	return result
}
