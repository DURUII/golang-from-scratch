package main

func largestRectangleArea(heights []int) int {
	n := len(heights)
	// 单调栈：求每个数左边第一个比它小的数的下标
	stk := make([]int, 0, n)
	left, right := make([]int, n), make([]int, n)

	for i := 0; i < n; i++ {
		for len(stk) > 0 && heights[stk[len(stk)-1]] >= heights[i] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) == 0 {
			left[i] = -1
		} else {
			left[i] = stk[len(stk)-1]
		}
		stk = append(stk, i)
	}

	stk = make([]int, 0, n)
	for i := n - 1; i >= 0; i-- {
		for len(stk) > 0 && heights[stk[len(stk)-1]] >= heights[i] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) == 0 {
			right[i] = -1
		} else {
			right[i] = stk[len(stk)-1]
		}
		stk = append(stk, i)
	}

	ret := -1
	for i := 0; i < n; i++ {
		ret = max(ret, (right[i]-left[i]-1)*heights[i])
	}
	return ret
}
