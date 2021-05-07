func singleNumber(nums []int) int {
	hashMap := make(map[int]int)
	for _, v := range nums {
		hashMap[v]++
	}

	for k, v := range hashMap {
		if v == 1 {
			return k
		}
	}
	return -1
}