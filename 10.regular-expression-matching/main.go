package main

func isMatch(s string, p string) bool {
	if len(p) > 1 && p[1] == '*' {
		for ; len(s) >= 0; s = s[1:] {
			if isMatch(s, p[2:]) {
				return true
			}
			if len(s) > 0 && (p[0] == '.' || p[0] == s[0]) {
				continue
			}
			break
		}
		return len(s) == 0 && len(p[2:]) == 0
	} else if len(p) > 0 {
		if len(s) == 0 {
			return false
		}
		if p[0] != '.' && p[0] != s[0] {
			return false
		}
		return isMatch(s[1:], p[1:])
	}

	return len(s) == 0
}

func main() {}
