package main

func myAtoi(str string) int {
	minus := 1
	b := []byte(str)

	for len(b) > 0 {
		if b[0] == '-' || b[0] == '+' {
			if b[0] == '-' {
				minus = -1
			}
			b = b[1:]
			break
		} else if b[0] >= '0' && b[0] <= '9' {
			break
		} else if b[0] == ' ' {
			b = b[1:]
			continue
		}
		return 0
	}

	p := 0
	for ; p < len(b); p++ {
		if b[p] >= '0' && b[p] <= '9' {
			continue
		}
		break
	}
	b = b[:p]

	if len(b) == 0 {
		return 0
	}

	result := float64(0)
	tmp := float64(0)
	for ; len(b) > 0; b = b[1:] {
		tmp = float64(b[0] - '0')
		for i := 1; i < len(b); i++ {
			tmp *= 10
		}
		result += tmp
	}

	if minus > 0 && result > 0x7fffffff {
		result = 0x7fffffff
	} else if minus < 0 && result > 0x80000000 {
		result = 0x80000000
	}

	return int(result) * minus
}

func main() {}
