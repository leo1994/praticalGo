package main

import (
	"fmt"
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
}

/* Channel Semantics
- send & receive wil block until opposite operation(*)
- receive from a closed channel will return the zerro value without blocking
- send to a closed channel will panic
- closing a closed channel will panic
- send/receive to a nil channel will block forever
*/
