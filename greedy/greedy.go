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

	fmt.Println("Greedy")
	fmt.Printf("Maximum amount that can be robbed: %d\n", result)
	fmt.Printf("Runtime: %s\n", elapsed)
}

func rob(nums []int) int {
	inc := nums[0]
	exc := 0
	for i := 1; i < len(nums); i++ {
		ninc := exc + nums[i]
		nexc := max(inc, exc)
		inc = ninc
		exc = nexc
	}
	return max(inc, exc)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
