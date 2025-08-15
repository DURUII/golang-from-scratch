package main

func firstUniqChar(s string) int {
	cnt := make(map[rune]int)
	for _, ch := range s {
		cnt[ch]++
	}
	for i, ch := range s {
		if cnt[ch] == 1 {
			return i
		}
	}
	return -1
}
