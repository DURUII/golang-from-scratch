package main

import "strings"

func numJewelsInStones(jewels string, stones string) int {
	ret := 0
	for _, j := range jewels {
		ret += strings.Count(stones, string(j))
	}
	return ret
}
