package main

import (
	"fmt"

	"github.com/xlauro/leetcode/solution"
)

func main() {
	nums := []int{3, 2, 4}
	target := 6
	result := solution.TwoSum(nums, target)
	fmt.Println(result)
}
