package main

import (
	"fmt"
	"sort"
	"time"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}

	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= res[len(res)-1][1] {
			res[len(res)-1][1] = max(intervals[i][1], res[len(res)-1][1])
		} else {
			res = append(res, intervals[i])
		}
	}
	return res
}

// 优化版本的区间合并函数
func mergeOptimized(intervals [][]int) [][]int {
	// 优化1: 边界检查，避免空数组和单元素数组的额外处理
	if len(intervals) <= 1 {
		return intervals
	}

	// 优化2: 按起始时间排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 优化3: 预分配切片容量，减少内存重新分配
	result := make([][]int, 0, len(intervals))
	result = append(result, intervals[0])

	// 优化4: 减少重复的切片访问，使用局部变量缓存
	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		lastIdx := len(result) - 1
		last := result[lastIdx]

		// 优化5: 判断是否重叠
		if current[0] <= last[1] {
			// 重叠：更新结束时间为两者最大值
			if current[1] > last[1] {
				result[lastIdx][1] = current[1]
			}
			// 如果current[1] <= last[1]，说明current完全被包含，无需更新
		} else {
			// 不重叠：添加新区间
			result = append(result, current)
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// 测试用例
	testCases := [][][]int{
		{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		{{1, 4}, {4, 5}},
		{{4, 7}, {1, 4}},
		{{1, 1}, {2, 2}, {3, 3}},
		{{1, 10}, {2, 3}, {4, 5}, {6, 7}, {8, 9}},
	}

	fmt.Println("=== 功能测试 ===")
	for i, intervals := range testCases {
		// 需要复制数组，因为排序会修改原数组
		intervals1 := make([][]int, len(intervals))
		intervals2 := make([][]int, len(intervals))
		copy(intervals1, intervals)
		copy(intervals2, intervals)

		result1 := merge(intervals1)
		result2 := mergeOptimized(intervals2)

		fmt.Printf("测试%d: %v\n", i+1, intervals)
		fmt.Printf("原版本: %v\n", result1)
		fmt.Printf("优化版: %v\n", result2)
		fmt.Printf("结果一致: %v\n\n", equalSlices(result1, result2))
	}

	// 性能测试
	fmt.Println("=== 性能测试 ===")
	// 生成大量测试数据
	largeTest := generateLargeTestCase(10000)

	// 测试原版本
	start := time.Now()
	for i := 0; i < 100; i++ {
		testData := make([][]int, len(largeTest))
		copy(testData, largeTest)
		merge(testData)
	}
	duration1 := time.Since(start)

	// 测试优化版本
	start = time.Now()
	for i := 0; i < 100; i++ {
		testData := make([][]int, len(largeTest))
		copy(testData, largeTest)
		mergeOptimized(testData)
	}
	duration2 := time.Since(start)

	fmt.Printf("原版本耗时: %v\n", duration1)
	fmt.Printf("优化版耗时: %v\n", duration2)
	fmt.Printf("性能提升: %.2f%%\n", float64(duration1-duration2)/float64(duration1)*100)
}

// 辅助函数：比较两个二维切片是否相等
func equalSlices(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

// 生成大量测试数据
func generateLargeTestCase(n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		start := i * 2
		end := start + 1 + i%3 // 创建一些重叠的区间
		result[i] = []int{start, end}
	}
	return result
}
