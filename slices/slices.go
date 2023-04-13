package main

import (
	"fmt"
	"sort"
)

func main() {
	var s []int

	fmt.Println("len", len(s))

	if s == nil {
		fmt.Println("nil slice")
	}

	s2 := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Printf("s2 = %#v\n", s2)

	s3 := s2[1:4]
	fmt.Printf("s3 = %#v\n", s3)

	s3 = append(s3, 100)

	fmt.Printf("s3 (append) = %#v\n", s3)
	fmt.Printf("s2 (append) = %#v\n", s2)
	fmt.Printf("s2 len=%d cap=%d\n", len(s2), cap(s2))
	fmt.Printf("s3 len=%d cap=%d\n", len(s3), cap(s3))

	var s4 []int
	// s4 := make([]int, 0, 1_000 b nte)
	for i := 0; i < 1_000; i++ {
		s4 = appendInt(s4, i)
	}
	fmt.Printf("s4 len=%d cap=%d\n", len(s4), cap(s4))

	fmt.Println(concat([]string{"A", "B"}, []string{"C", "D", "E"})) // [A B C E D]

	fmt.Println(median([]float64{2, 3, 1}))    // [A B C E D]
	fmt.Println(median([]float64{2, 3, 1, 4})) // [A B C E D]

}

func median(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("median of empty slice")
	}

	// fix in order not to change values
	nums := make([]float64, len(values))

	copy(nums, values)
	sort.Float64s(nums)

	i := len(nums) / 2

	// len(nums)&1 == 1
	if len(nums)%2 == 1 {
		return nums[i], nil
	}
	return (nums[i-1] + nums[i]) / 2, nil
}

func concat(s1, s2 []string) []string {
	s3 := make([]string, len(s1)+len(s2))
	copy(s3, s1)
	copy(s3[len(s1):], s2)
	return s3
}
func appendInt(s []int, v int) []int {
	i := len(s)
	if len(s) < cap(s) {
		s = s[:len(s)+1]
	} else {
		fmt.Printf("reallocate: %d->%d\n", len(s), 2*len(s)+1)
		s2 := make([]int, 2*len(s)+1)
		copy(s2, s)
		s = s2[:len(s)+1]
	}

	s[i] = v
	return s

}
