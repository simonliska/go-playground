func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	if len(strs) > 0 {
		for _, str := range strs {
			for !strings.HasPrefix(str, prefix) {
				prefix = prefix[:len(prefix)-1]
			}
		}
	}
	return prefix
}