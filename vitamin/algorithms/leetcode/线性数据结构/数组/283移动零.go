package main

func moveZeroes(nums []int) {
	writePos := 0
	for _, num := range nums {
		if num != 0 { // filter condition
			nums[writePos] = num
			writePos++
		}
	}
	for i := writePos; i < len(nums); i++ {
		nums[i] = 0
	}
}
