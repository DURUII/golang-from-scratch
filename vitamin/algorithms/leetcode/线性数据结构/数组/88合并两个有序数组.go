package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1 // dual read position
	for k := m + n - 1; k >= 0; k-- {
		readList, readPos := nums1, &i
		if j < 0 {
			break
		}
		if i < 0 || nums2[j] > nums1[i] {
			readList, readPos = nums2, &j
		}
		nums1[k] = readList[*readPos]
		*readPos--
	}
}
