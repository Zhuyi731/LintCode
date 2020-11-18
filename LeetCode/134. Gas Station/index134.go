package main

func main() {
	canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2})
	canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3})
}

func canCompleteCircuit(gas []int, cost []int) int {
	sum := make([]int, len(gas))
	min := 0x7fffff
	minIndex := 0
	for i := range gas {
		sum[i] = gas[i] - cost[i]
		if i != 0 {
			sum[i] = sum[i] + sum[i-1]
		}

		if sum[i] < min {
			min = sum[i]
			minIndex = i
		}
	}

	nextIndex := (minIndex + 1) % len(gas)

	if sum[len(gas)-1] >= 0 && gas[nextIndex]-cost[nextIndex]-sum[nextIndex]+sum[minIndex] >= 0 {
		return nextIndex
	} else {
		return -1
	}
}
