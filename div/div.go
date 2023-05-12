package main

import (
	"fmt"
	"log"
)

func main() {
	// fmt.Println(div(1, 0))
	fmt.Println(safeDiv(1, 0))
}

func div(a, b int) int {
	return a / b
}

func safeDiv(a, b int) (q int, err error) {
	// q && err are local variables in safeDiv
	// (just like a & b)
	defer func() {
		if e := recover(); e != nil {
			log.Println("ERROS:", e)
			err = fmt.Errorf("%v", e)
		}
	}()
	/**
	q = a / b
	return
	**/

	return a / q, nil
}
