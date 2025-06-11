package solution

func RemoveElement(nums []int, val int) int {
	var aux []int

	for _, value := range nums {
		if value != val {
			aux = append(aux, value)
		}
	}
	nums = aux
	return len(nums)
}
