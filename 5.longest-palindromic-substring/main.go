package main

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}

	p := s[:1]

	for cusor := range s[:] {
		slice := s[cusor:]
		i := 0
		end := len(slice) - 1
		j := end
		for i < j {
			if slice[i] != slice[j] {
				i = 0
				end--
				j = end
				continue
			}
			i++
			j--
		}

		if end != j {
			if len(slice[:end+1]) > len(p) {
				p = slice[:end+1]
			}
		}
	}

	return p
}

func main() {}
