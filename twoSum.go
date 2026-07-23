package main

import (
	"slices"
)

// https://neetcode.io/problems/two-integer-sum/question
// this is notated as the brute force solution, but is
// faster than 88% of the known solutions on neetcode
// execution time: 2m
func twoSum(nums []int, target int) []int {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// this is faster yet - 1ms
func twoSumWithMap(nums []int, target int) []int {
	intsSeen := map[int]int{}
	for i := range nums {
		complement := target - nums[i]
		if j, ok := intsSeen[complement]; ok && j != i {
			result := []int{i, j}
			slices.Sort(result)
			return result
		}
		intsSeen[nums[i]] = i
	}
	return []int{}
}