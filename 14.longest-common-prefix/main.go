package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 || len(strs[0]) == 0 {
		return ""
	}

	var i int

check:
	for i = range strs[0] {
		for _, v := range strs[1:] {
			if i+1 > len(v) || v[:i+1] != strs[0][:i+1] {
				i--
				break check
			}
		}
	}

	return strs[0][:i+1]
}

func main() {}
