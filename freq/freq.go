package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

// Q: what is the mons common work in sherlock.txt

func main() {
	file, err := os.Open("sherlock.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer file.Close()

	words, err := wordFrequency(file)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(maxWord(words))
}

var wordReg = regexp.MustCompile(`[a-zA-Z]+`)

func MapDemo() {
	var stocks map[string]float64 //symbol -> price
	sym := "TTWO"
	price := stocks[sym]
	fmt.Printf("%s -> $%.2f\n", sym, price)

	if price, ok := stocks[sym]; ok {
		fmt.Printf("%s -> $%.2f\n", sym, price)
	} else {
		fmt.Printf("%s not found\n", sym)
	}

	// stocks = make(map[string]float64)
	// stocks[sym] = 136.73
	stocks = map[string]float64{
		sym:    136.76,
		"AAPL": 124.64,
	}

	if price, ok := stocks[sym]; ok {
		fmt.Printf("%s -> $%.2f\n", sym, price)
	} else {
		fmt.Printf("%s not found\n", sym)
	}

	for k, v := range stocks {
		fmt.Printf("%s -> $%.2f\n", k, v)
	}

	delete(stocks, "AAPL")
}

func maxWord(freqs map[string]int) (string, error) {
	if len(freqs) == 0 {
		return "", fmt.Errorf("empty map")
	}

	maxN, maxW := 0, ""
	for w, c := range freqs {
		if c > maxN {
			maxW, maxN = w, c
		}
	}

	return maxW, nil
}

func wordFrequency(r io.Reader) (map[string]int, error) {
	s := bufio.NewScanner(r)
	freqs := make(map[string]int)
	lnum := 0
	for s.Scan() {
		lnum++

		words := wordReg.FindAllString(s.Text(), -1)
		for _, w := range words {
			freqs[strings.ToLower(w)]++
		}
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	fmt.Println("Num of lines:", lnum)
	return freqs, nil
}
