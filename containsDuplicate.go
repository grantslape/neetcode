package main

// https://neetcode.io/problems/duplicate-integer/question?list=neetcode150
func hasDuplicate(nums []int) bool {
	intsSeen := map[int]struct{}{}
	for i := range nums {
		if _, ok := intsSeen[nums[i]]; ok {
			return true
		}
		intsSeen[nums[i]] = struct{}{}
	}
	return false
}
