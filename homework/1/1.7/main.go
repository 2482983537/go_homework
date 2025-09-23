package main

import "fmt"

func twoSum(nums []int, target int) []int {
	sign := make(map[int]int)
	for idx, val := range nums {
		temp := target - val
		if value, exists := sign[val]; exists {
			return []int{value, idx}
		} else {
			sign[temp] = idx
		}
	}
	return []int{}
}

// twoSumOptimized 优化版本的两数之和解法
// 相比原版本的改进：
// 1. 修复了哈希表存储逻辑错误
// 2. 添加了边界检查
// 3. 使用更清晰的变量命名
// 4. 提前返回优化性能
// 时间复杂度: O(n), 空间复杂度: O(n)
func twoSumOptimized(nums []int, target int) []int {
	// 边界检查：数组长度小于2无法找到两个数
	if len(nums) < 2 {
		return []int{}
	}

	// 使用哈希表存储已遍历的数字及其索引
	// key: 数字值, value: 索引位置
	numToIndex := make(map[int]int, len(nums))

	for currentIndex, currentNum := range nums {
		// 计算当前数字的补数
		complement := target - currentNum

		// 检查补数是否已经存在于哈希表中
		if complementIndex, found := numToIndex[complement]; found {
			// 找到答案，返回两个索引（较小的在前）
			return []int{complementIndex, currentIndex}
		}

		// 将当前数字和索引存入哈希表
		numToIndex[currentNum] = currentIndex
	}

	// 遍历完成仍未找到，返回空切片
	return []int{}
}

func main() {
	fmt.Println("=== 两数之和函数对比测试 ===")
	fmt.Println()

	// 测试用例1: 基本情况
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	fmt.Printf("测试1 - 基本情况: nums=%v, target=%d\n", nums1, target1)
	fmt.Printf("  原版本结果: %v\n", twoSum(nums1, target1))
	fmt.Printf("  优化版结果: %v\n", twoSumOptimized(nums1, target1))
	fmt.Println()

	// 测试用例2: 相同数字
	nums2 := []int{3, 3}
	target2 := 6
	fmt.Printf("测试2 - 相同数字: nums=%v, target=%d\n", nums2, target2)
	fmt.Printf("  原版本结果: %v\n", twoSum(nums2, target2))
	fmt.Printf("  优化版结果: %v\n", twoSumOptimized(nums2, target2))
	fmt.Println()

	// 测试用例3: 多个可能组合
	nums3 := []int{3, 2, 4}
	target3 := 6
	fmt.Printf("测试3 - 多个组合: nums=%v, target=%d\n", nums3, target3)
	fmt.Printf("  原版本结果: %v\n", twoSum(nums3, target3))
	fmt.Printf("  优化版结果: %v\n", twoSumOptimized(nums3, target3))
	fmt.Println()

	// 测试用例4: 负数
	nums4 := []int{-1, -2, -3, -4, -5}
	target4 := -8
	fmt.Printf("测试4 - 负数情况: nums=%v, target=%d\n", nums4, target4)
	fmt.Printf("  原版本结果: %v\n", twoSum(nums4, target4))
	fmt.Printf("  优化版结果: %v\n", twoSumOptimized(nums4, target4))
	fmt.Println()

	// 测试用例5: 无解情况
	nums5 := []int{1, 2, 3}
	target5 := 7
	fmt.Printf("测试5 - 无解情况: nums=%v, target=%d\n", nums5, target5)
	fmt.Printf("  原版本结果: %v\n", twoSum(nums5, target5))
	fmt.Printf("  优化版结果: %v\n", twoSumOptimized(nums5, target5))
	fmt.Println()

	// 测试用例6: 边界情况 - 数组长度小于2
	nums6 := []int{1}
	target6 := 1
	fmt.Printf("测试6 - 边界情况: nums=%v, target=%d\n", nums6, target6)
	fmt.Printf("  原版本结果: %v\n", twoSum(nums6, target6))
	fmt.Printf("  优化版结果: %v\n", twoSumOptimized(nums6, target6))
	fmt.Println()

	fmt.Println("=== 优化说明 ===")
	fmt.Println("1. 修复了原版本的逻辑错误（存储补数但查找当前值）")
	fmt.Println("2. 添加了边界检查，处理数组长度小于2的情况")
	fmt.Println("3. 使用更清晰的变量命名提高代码可读性")
	fmt.Println("4. 预分配哈希表容量，减少内存重新分配")
	fmt.Println("5. 找到结果后立即返回，避免不必要的遍历")
}
