package main

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	result := make([]byte, 0, len(s))

	for i := 0; i < numRows; i++ {
		m := numRows + (numRows - 2) - 2*i
		n := numRows + (numRows - 2) - m

		for j := i; j < len(s); {
			if m > 0 && j < len(s) {
				result = append(result, s[j])
				j += m
			}

			if n > 0 && j < len(s) {
				result = append(result, s[j])
				j += n
			}
		}
	}

	return string(result)
}

func main() {}
