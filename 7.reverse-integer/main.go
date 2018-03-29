package main

func reverse(x int) int {
	minus := 1
	if x < 0 {
		minus = -1
		x = -x
	}

	var b []int

	for x > 0 {
		b = append(b, x%10)
		x /= 10
	}

	result := 0
	for ; len(b) > 0; b = b[1:] {
		for i := 1; i < len(b); i++ {
			b[0] *= 10
		}
		result += b[0]
	}

	if result > 0x7fffffff {
		result = 0
	}

	return result * minus
}

func main() {}
