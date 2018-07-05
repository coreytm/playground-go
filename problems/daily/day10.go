package daily

import (
	"fmt"
	"time"
)

// This problem was asked by Apple.
// Implement a job scheduler which takes in a function f and an integer n, and calls f after n milliseconds.

var foo = func() {
	fmt.Printf("This is func foo!")
}

// schedule accept a function and a delay in ms as arguments and runs the function after the delay.
func scheduler(f func(), ms int64) {
	t := time.NewTimer(time.Millisecond * time.Duration(ms))
	<-t.C
	f()
}
