package main

import "slices"

func topKFrequent(nums []int, k int) []int {
	frequency := make(map[int]uint)
	for _, num := range nums {
		if _, ok := frequency[num]; !ok {
			frequency[num] = 0
		}
		frequency[num]++
	}

	sortedFrequency := make([]int, 0, len(frequency))
	for num := range frequency {
		sortedFrequency = append(sortedFrequency, num)
	}
	slices.SortFunc(sortedFrequency, func(a, b int) int {
		return int(frequency[b] - frequency[a])
	})

	return sortedFrequency[:k]
}


func main() {

}
