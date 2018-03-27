package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mergedNums := make([]int, 0, len(nums1)+len(nums2))

	for len(nums1) > 0 || len(nums2) > 0 {
		if len(nums1) == 0 {
			mergedNums = append(mergedNums, nums2[0])
			nums2 = nums2[1:]
			continue
		}

		if len(nums2) == 0 {
			mergedNums = append(mergedNums, nums1[0])
			nums1 = nums1[1:]
			continue
		}

		if nums1[0] > nums2[0] {
			mergedNums = append(mergedNums, nums2[0])
			nums2 = nums2[1:]
		} else {
			mergedNums = append(mergedNums, nums1[0])
			nums1 = nums1[1:]
		}
	}

	if len(mergedNums)%2 == 1 {
		return float64(mergedNums[len(mergedNums)/2])
	}

	return float64(mergedNums[len(mergedNums)/2-1]+mergedNums[len(mergedNums)/2]) / 2
}

func main() {}
