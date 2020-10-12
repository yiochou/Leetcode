func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	index := 0

	for {
		for i := 0; i < len(strs); i++ {
			if index >= len(strs[i]) {
				return string(strs[0][:index])
			}
		}
		for i := 0; i < len(strs)-1; i++ {
			if strs[i][index] != strs[i+1][index] {
				return string(strs[0][:index])
			}
		}
		index++;
	}
}