package main

func removeDuplicates(nums []int) int {
	// 处理空数组的情况
	if len(nums) == 0 {
		return 0
	}

	// k用于记录不重复元素的位置
	k := 1

	// 遍历数组，从第二个元素开始
	for i := 1; i < len(nums); i++ {
		// 如果当前元素与前一个不重复元素不同
		if nums[i] != nums[k-1] {
			// 将当前元素放到k位置
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
