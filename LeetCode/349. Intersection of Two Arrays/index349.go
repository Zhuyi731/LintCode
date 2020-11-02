

func intersection(nums1 []int, nums2 []int) []int {
	mapper := make(map[int]bool)
	for _, v := range nums1 {
		mapper[v] = true
	}
	result := make([]int, 0)

	for _, v := range nums2 {
		if mapper[v] {
			result = append(result, v)
			mapper[v] = false
		}
	}
	return result
}