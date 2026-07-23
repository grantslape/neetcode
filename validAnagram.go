package main

// https://neetcode.io/problems/is-anagram/question?list=neetcode150
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var (
		charactersInS = map[byte]uint{}
		charactersInT = map[byte]uint{}
	)

	for i := range len(s) {
		addRunetoSet(charactersInS, s[i])
		addRunetoSet(charactersInT, t[i])
	}

	for char, count := range charactersInS {
		countInT, ok := charactersInT[char]
		if !ok {
			return false
		}
		if countInT != count {
			return false
		}
		delete(charactersInS, char)
		delete(charactersInT, char)
	}

	if len(charactersInS) != 0 || len(charactersInT) != 0 {
		return false
	}

	return true
}

func addRunetoSet(set map[byte]uint, char byte) {
	if _, ok := set[char]; !ok {
		set[char] = 0
	}
	set[char] += 1
}
