package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	banner("Go", 6)
	banner("G😀", 6)

	s := "G☺"

	fmt.Println("len:", len(s))

	for i, r := range s {
		fmt.Println(i, r)
		if i == 0 {
			fmt.Printf("%c of type %T\n", r, r)
		}
	}

	b := s[0]

	fmt.Printf("%c of type %T\n", b, b)

	x, y := 1, "1"

	fmt.Printf("x=%v, y=%v\n", x, y)
	fmt.Printf("x=%#v, y=%#v\n", x, y)

	fmt.Printf("%20s\n", s)

	fmt.Println("g", isPalindrome("g"))
	fmt.Println("go", isPalindrome("go"))
	fmt.Println("gog", isPalindrome("gog"))
	fmt.Println("gogo", isPalindrome("gogo"))
	fmt.Println("g☺g", isPalindrome("g☺g"))
}

func isPalindrome(s string) bool {
	rs := []rune(s)
	for i := 0; i < len(rs)/2; i++ {
		if rs[i] != rs[len(rs)-1-i] {
			return false
		}
	}

	return true
}

func banner(text string, width int) {
	padding := (width - utf8.RuneCountInString(text)) / 2
	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}

	fmt.Println(text)

	for i := 0; i < width; i++ {
		fmt.Print("-")
	}

	fmt.Println()
}
