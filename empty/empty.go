package main

import "fmt"

func main() {
	var i any

	i = 7
	fmt.Println(i)

	i = "hi"
	fmt.Println(i)

	v := i.(string)
	fmt.Println("v:", v)

	n, ok := i.(int)

	fmt.Println(n, ok)

	switch i.(type) {
	case int:
		fmt.Println("Is int")
	case string:
		fmt.Println("Is string")
	default:
		fmt.Printf("unknown type: %T\n", i)
	}
	fmt.Println(maxInts([]int{1, 2, 3}))
	fmt.Println(maxFloat64s([]float64{1, 2, 3}))
	fmt.Println(max([]float64{1, 2, 3}))
	fmt.Println(max([]int{1, 2, 3}))
}

func max[T int | float64](nums []T) T {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for _, n := range nums[1:] {
		if n > max {
			max = n
		}
	}

	return max
}

func maxInts(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for _, n := range nums[1:] {
		if n > max {
			max = n
		}
	}

	return max
}

func maxFloat64s(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for _, n := range nums[1:] {
		if n > max {
			max = n
		}
	}

	return max
}
