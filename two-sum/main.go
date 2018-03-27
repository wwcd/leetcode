package main

func twoSum(nums []int, target int) []int {
	mapNums := make(map[int][]int)
	for i, v := range nums {
		mapNums[v] = append(mapNums[v], i)
	}

	for k, v := range mapNums {
		lk := target - k
		if lv, ok := mapNums[lk]; ok {
			if lv[0] == v[0] {
				if len(lv) > 1 {
					return lv[:2]
				}
				continue
			}
			return []int{v[0], lv[0]}
		}
	}

	return nil
}

func main() {
}
