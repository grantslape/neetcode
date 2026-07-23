package main

import (
	"strconv"
)

const separator = "%"

type Solution struct{}

func (s *Solution) Encode(strs []string) (ret string) {
	for _, str := range strs {
		// %6%Hello%5%World%10%!@#$%^&*()
		ret += separator + strconv.Itoa(len(str)) + separator + str
	}
	return ret
}

func (s *Solution) Decode(encoded string) (ret []string) {
	runes := []rune(encoded)
	for i := 0; i < len(runes); i++ {
		if runes[i] == '%' {
			// read all the way to the next %, and then read the length
			end := i + 1
			for end < len(runes) && runes[end] != '%' {
				end++
			}
			length, err := strconv.Atoi(string(runes[i+1 : end]))
			if err != nil {
				panic(err)
			}
			i = end
			decoded := string(runes[i+1 : i+1+length])
			ret = append(ret, decoded)
			i += length
		}
	}
	return ret
}

func main() {

}
