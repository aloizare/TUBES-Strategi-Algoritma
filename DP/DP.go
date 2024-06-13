package main

import (
	"fmt"
	"time"
)

func main() {
	nums := []int{2, 7, 9, 3, 1} //masukkan nilai-nilai rumah di sini {2, 7, 9, 3, 1} bisa diganti dengan nilai rumah yang diinginkan

	start := time.Now()
	result := rob(nums)
	elapsed := time.Since(start)

	fmt.Println("Dynamic Programming")
	fmt.Printf("Maximum amount that can be robbed: %d\n", result)
	fmt.Printf("Runtime: %s\n", elapsed)
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	prev2 := 0
	prev1 := 0
	for _, num := range nums {
		tmp := prev1
		prev1 = max(prev2+num, prev1)
		prev2 = tmp
	}
	return prev1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
