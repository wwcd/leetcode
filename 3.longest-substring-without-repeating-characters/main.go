package main

func lengthOfLongestSubstring(s string) int {
	var len int

	for i, _ := range s {
		had := make(map[rune]bool)
		j := 0
		for _, c := range s[i:] {
			if had[c] {
				break
			}
			had[c] = true
			j++
		}
		if j > len {
			len = j
		}
	}

	return len
}

func main() {}
