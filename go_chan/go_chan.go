package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("goroutine")
	fmt.Println("main")

	for i := 0; i < 3; i++ {
		// FIX 2: use a loop body variable
		i := i // "i" shadows "i" from the for loop
		go func() {
			fmt.Println(i)
		}()

		time.Sleep(10 * time.Millisecond)

		// FIX 1: Use a parameter
		// go func(n int) {
		// 	fmt.Println(n)
		// }(i)

		// BUG: All goroutines use the "i" for the for loop
		// go func() {
		// 	fmt.Println(i)
		// }()
	}

	ch := make(chan string)
	go func() {
		ch <- "hi" //send
	}()
	msg := <-ch //receive
	fmt.Println(msg)

	go func() {
		for i := 0; i < 3; i++ {
			msg := fmt.Sprintf("message #%v", i+1)
			ch <- msg
		}
		close(ch)
	}()

	for msg := range ch {
		fmt.Println("got", msg)
	}

	// for/range does this (syntax sugar)
	/* for {
		msg, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("got", msg)
	} */

	msg, ok := <-ch
	fmt.Printf("closed: %#v (ok=%v)\n", msg, ok)

	// ch <- "hi" // ch is closed

	values := []int{10, 2, 13, 7, 1, 24, 9, 5}

	fmt.Println(sleepSort(values))

}

/*
	For every value "n" in values, spin a goroutine that will
	- spleep "n" milliseconds
	- Send "n" over a channel

IN the funciton body, collect values from the channel to a slice and return it
*/

func sleepSort(values []int) []int {
	ch := make(chan int)
	for _, v := range values {
		v := v
		go func() {
			time.Sleep(time.Duration(v) * time.Millisecond)
			ch <- v
		}()
	}

	var out []int
	// for i := 0; i < len(values); i++
	for range values {
		n := <-ch
		out = append(out, n)
	}
	return out
}

/* Channel Semantics
- send & receive wil block until opposite operation(*)
	- buffered channels has cap(ch) non-blocking send operations
- receive from a closed channel will return the zerro value without blocking
- send to a closed channel will panic
- closing a closed channel will panic
- send/receive to a nil channel will block forever
*/
