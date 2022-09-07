// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
)

var i int = 0

func incrementing(ch chan int, finished chan int) {
	for j := 0; j < 1000000; j++ {
		i := <-ch
		i++
		ch <- i
	}
	finished <- 1
}

func decrementing(ch chan int, finished chan int) {
	for j := 0; j < 1000000; j++ {
		i := <-ch
		i--
		ch <- i
	}
	finished <- 1
}

func main() {
	ch := make(chan int, 1) //a buffered channel
	finished := make(chan int, 1)
	ch <- i

	runtime.GOMAXPROCS(2)

	go incrementing(ch, finished)
	go decrementing(ch, finished)

	// We have no way to wait for the completion of a goroutine (without additional syncronization of some sort)
	// We'll come back to using channels in Exercise 2. For now: Sleep.
	<-finished
	<-finished
	i = <-ch
	Println("The magic number is:", i)
}
