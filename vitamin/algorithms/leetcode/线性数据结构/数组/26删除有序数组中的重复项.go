package main

// 过滤器模版题
func removeDuplicates(nums []int) int {
	w := 0                           // 代表长度，也代表写指针索引
	for r := 0; r < len(nums); r++ { // 读指针
		if r == 0 || nums[r] != nums[r-1] { // 注意越界问题
			nums[w] = nums[r]
			w++
		}
	}
	return w
}
