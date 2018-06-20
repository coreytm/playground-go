package problems

import (
	"math/rand"
	"time"
)

// We want a random number from 1 to 7.
// How do we accomplish this with a 5 sided die?

func rand5() int {
	return rand.New(rand.NewSource(time.Now().Unix())).Intn(4) + 1
}

func rand7UsingRand5() int {
	// we can do this in two rolls
	// roll1 => 1..2
	// roll2 => 0..4

	// two rolls have 5^2 outcomes
	// but 25 ! divide 7, so retry if 22..25

	lookup := 25
	for lookup > 21 {
		rand1 := rand5()
		rand2 := rand5()

		lookup = (rand1-1)*5 + (rand2 - 1) + 1
	}

	return lookup%7 + 1
}
