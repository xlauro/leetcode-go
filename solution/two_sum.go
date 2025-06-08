package solution

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, val := range nums {
		if j, ok := m[target-val]; ok {
			return []int{j, i}
		}
		m[val] = i
	}
	return nil
}
