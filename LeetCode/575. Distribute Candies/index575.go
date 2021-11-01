func distributeCandies(candyType []int) int {
	mp := map[int]int{}
	for _, v := range candyType {
		mp[v]++
	}
	lc := len(candyType)
	if len(mp) > lc/2 {
		return lc / 2
	} else {
		return len(mp)
	}
}