package solution

import (
	"sort"
)

func Merge(nums1 []int, m int, nums2 []int, n int) {
	sort.Ints(nums1)
	nums1 = nums1[:m]
	sort.Ints(nums2)
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
}
