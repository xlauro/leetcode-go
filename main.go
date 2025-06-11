package main

import (
	"fmt"

	"github.com/xlauro/leetcode/solution"
)

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	target := 2
	result := solution.RemoveElement(nums, target)
	fmt.Println(result)
}
