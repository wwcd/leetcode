package main

func isPalindrome(x int) bool {
	var b []int

	if x < 0 {
		return false
	}

	for x > 0 {
		b = append(b, x%10)
		x /= 10
	}

	i := 0
	j := len(b) - 1
	for i < j {
		if b[i] != b[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func main() {}
