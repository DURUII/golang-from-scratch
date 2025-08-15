package main

import "slices"

var dp = make([]int, 1000)
var prev = make([]int, 1000)

func largestDivisibleSubset(nums []int) []int {
	// sort
	slices.Sort(nums)
	// boss result
	bossIdx := 0
	// find the largest num of the subset ending with nums[i]
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		prev[i] = -1
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				prev[i] = j
			}
		}
		if dp[i] > dp[bossIdx] {
			bossIdx = i
		}
	}

	// render the solution
	ret := make([]int, 0, dp[bossIdx])
	for i := bossIdx; i >= 0; i = prev[i] {
		ret = append(ret, nums[i])
		if prev[i] == -1 {
			break
		}
	}
	return ret
}
