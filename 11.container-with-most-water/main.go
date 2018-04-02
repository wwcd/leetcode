package main

func maxArea(height []int) int {
	area := 0

	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			h := height[i]
			if height[i] > height[j] {
				h = height[j]
			}
			w := j - i

			if h*w > area {
				area = h * w
			}
		}
	}

	return area
}

func main() {}
