package main

import (
	"fmt"
	"slices"
)

// https://neetcode.io/problems/anagram-groups/question?list=neetcode150
func sortString(s string) string {
	characters := []rune(s)
	slices.Sort(characters)
	return string(characters)
}

func groupAnagrams(strs []string) (ret [][]string) {
	res := make(map[string][]string)
	for _, str := range strs {
		sortedStr := sortString(str)
		res[sortedStr] = append(res[sortedStr], str)
	}
	for _, v := range res {
		ret = append(ret, v)
	}
	return ret
}

func main() {
	strs := []string{"act", "pots", "tops", "cat", "stop", "hat"}
	fmt.Println(groupAnagrams(strs))
}
