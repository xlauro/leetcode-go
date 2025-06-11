package solution

func mergeAlternately(word1 string, word2 string) string {
	result := ""

	if len(word1) >= len(word2) {
		for i := 0; i < len(word1); i++ {
			result += string(word1[i])
			if i < len(word2) {
				result += string(word2[i])
			}
		}
	} else {
		for i := 0; i < len(word2); i++ {
			if i < len(word1) {
				result += string(word1[i])
			}
			result += string(word2[i])
		}
	}
	return result
}
