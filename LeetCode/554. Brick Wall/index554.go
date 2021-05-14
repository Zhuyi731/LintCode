package main

func main() {

}

func leastBricks(wall [][]int) int {
	hashMap := map[int]int{}
	curMax := 0
	for _, w := range wall {
		cur := 0
		for i, v := range w {
			if i != len(w)-1 {
				cur = cur + v
				hashMap[cur] = hashMap[cur] + 1
				if hashMap[cur] > curMax {
					curMax = hashMap[cur]
				}
			}
		}
	}
	return len(wall) - curMax
}
