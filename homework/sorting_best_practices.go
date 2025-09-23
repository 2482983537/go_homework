// Go语言排序最佳实践和实际应用场景
package main

import (
	"fmt"
	"sort"
	"strconv"

	"time"
)

// 实际业务场景的数据结构
type BestPracticeEmployee struct {
	ID         int
	Name       string
	Department string
	Salary     float64
	HireDate   time.Time
	Skills     []string
}

type BestPracticeOrder struct {
	ID       string
	Amount   float64
	Date     time.Time
	Priority int
	Status   string
}

// 排序最佳实践演示
func DemonstrateSortingBestPractices() {
	fmt.Println("=== Go语言排序最佳实践 ===")
	fmt.Println()

	// 1. 预计算排序键优化
	demonstratePrecomputedKeys()

	// 2. 复杂业务逻辑排序
	demonstrateComplexBusinessSorting()

	// 3. 多级排序策略
	demonstrateMultiLevelSorting()

	// 4. 字符串自然排序
	demonstrateNaturalStringSorting()

	// 5. 排序稳定性的实际应用
	demonstrateSortingStabilityUsage()

	// 6. 大数据排序优化
	demonstrateLargeDataSorting()

	// 7. 排序错误处理和边界情况
	demonstrateSortingEdgeCases()

	fmt.Println("排序最佳实践演示完成！")
}

// 1. 预计算排序键优化
func demonstratePrecomputedKeys() {
	fmt.Println("1. 预计算排序键优化:")
	fmt.Println()

	employees := []BestPracticeEmployee{
		{1, "张三", "技术部", 8000, time.Date(2020, 1, 15, 0, 0, 0, 0, time.UTC), []string{"Go", "Python"}},
		{2, "李四", "销售部", 6000, time.Date(2019, 3, 20, 0, 0, 0, 0, time.UTC), []string{"销售", "沟通"}},
		{3, "王五", "技术部", 12000, time.Date(2018, 7, 10, 0, 0, 0, 0, time.UTC), []string{"Java", "架构"}},
		{4, "赵六", "人事部", 7000, time.Date(2021, 5, 8, 0, 0, 0, 0, time.UTC), []string{"HR", "招聘"}},
	}

	fmt.Println("原始数据:")
	printBestPracticeEmployees(employees)

	// 方法1: 直接排序（每次比较都计算）
	employees1 := make([]BestPracticeEmployee, len(employees))
	copy(employees1, employees)

	sort.Slice(employees1, func(i, j int) bool {
		// 每次比较都要计算工作年限
		years1 := time.Since(employees1[i].HireDate).Hours() / (24 * 365)
		years2 := time.Since(employees1[j].HireDate).Hours() / (24 * 365)
		return years1 > years2 // 按工作年限降序
	})

	fmt.Println("方法1 - 直接排序（按工作年限降序）:")
	printBestPracticeEmployees(employees1)

	// 方法2: 预计算排序键（推荐）
	type EmployeeWithKey struct {
		BestPracticeEmployee
		WorkYears float64
	}

	employeesWithKeys := make([]EmployeeWithKey, len(employees))
	for i, emp := range employees {
		employeesWithKeys[i] = EmployeeWithKey{
			BestPracticeEmployee: emp,
			WorkYears:            time.Since(emp.HireDate).Hours() / (24 * 365),
		}
	}

	sort.Slice(employeesWithKeys, func(i, j int) bool {
		return employeesWithKeys[i].WorkYears > employeesWithKeys[j].WorkYears
	})

	fmt.Println("方法2 - 预计算排序键（推荐）:")
	for _, empWithKey := range employeesWithKeys {
		fmt.Printf("  %s (%.1f年) - %s - ¥%.0f\n",
			empWithKey.Name, empWithKey.WorkYears, empWithKey.Department, empWithKey.Salary)
	}
	fmt.Println()
}

// 2. 复杂业务逻辑排序
func demonstrateComplexBusinessSorting() {
	fmt.Println("2. 复杂业务逻辑排序:")
	fmt.Println()

	orders := []BestPracticeOrder{
		{"ORD001", 1500.0, time.Date(2023, 12, 1, 10, 0, 0, 0, time.UTC), 1, "pending"},
		{"ORD002", 800.0, time.Date(2023, 12, 1, 14, 0, 0, 0, time.UTC), 3, "processing"},
		{"ORD003", 2000.0, time.Date(2023, 12, 2, 9, 0, 0, 0, time.UTC), 1, "urgent"},
		{"ORD004", 500.0, time.Date(2023, 12, 1, 16, 0, 0, 0, time.UTC), 2, "pending"},
		{"ORD005", 3000.0, time.Date(2023, 12, 3, 11, 0, 0, 0, time.UTC), 2, "completed"},
	}

	fmt.Println("原始订单:")
	printBestPracticeOrders(orders)

	// 复杂业务排序规则:
	// 1. urgent状态优先级最高
	// 2. 然后按优先级数字排序（1最高）
	// 3. 相同优先级按金额降序
	// 4. 相同金额按时间升序
	sort.Slice(orders, func(i, j int) bool {
		orderI, orderJ := orders[i], orders[j]

		// 1. urgent状态优先
		if orderI.Status == "urgent" && orderJ.Status != "urgent" {
			return true
		}
		if orderI.Status != "urgent" && orderJ.Status == "urgent" {
			return false
		}

		// 2. 按优先级排序
		if orderI.Priority != orderJ.Priority {
			return orderI.Priority < orderJ.Priority
		}

		// 3. 按金额降序
		if orderI.Amount != orderJ.Amount {
			return orderI.Amount > orderJ.Amount
		}

		// 4. 按时间升序
		return orderI.Date.Before(orderJ.Date)
	})

	fmt.Println("复杂业务逻辑排序后:")
	printBestPracticeOrders(orders)
}

// 3. 多级排序策略
func demonstrateMultiLevelSorting() {
	fmt.Println("3. 多级排序策略:")
	fmt.Println()

	employees := []BestPracticeEmployee{
		{1, "张三", "技术部", 8000, time.Date(2020, 1, 15, 0, 0, 0, 0, time.UTC), []string{"Go", "Python"}},
		{2, "李四", "技术部", 8000, time.Date(2019, 3, 20, 0, 0, 0, 0, time.UTC), []string{"Java", "Spring"}},
		{3, "王五", "销售部", 6000, time.Date(2021, 7, 10, 0, 0, 0, 0, time.UTC), []string{"销售", "客户"}},
		{4, "赵六", "销售部", 6000, time.Date(2020, 5, 8, 0, 0, 0, 0, time.UTC), []string{"市场", "推广"}},
		{5, "钱七", "技术部", 12000, time.Date(2018, 7, 10, 0, 0, 0, 0, time.UTC), []string{"架构", "管理"}},
	}

	fmt.Println("原始员工数据:")
	printBestPracticeEmployees(employees)

	// 多级排序: 部门 -> 薪资(降序) -> 入职时间(升序)
	sort.Slice(employees, func(i, j int) bool {
		empI, empJ := employees[i], employees[j]

		// 第一级: 按部门排序
		if empI.Department != empJ.Department {
			return empI.Department < empJ.Department
		}

		// 第二级: 按薪资降序
		if empI.Salary != empJ.Salary {
			return empI.Salary > empJ.Salary
		}

		// 第三级: 按入职时间升序
		return empI.HireDate.Before(empJ.HireDate)
	})

	fmt.Println("多级排序后 (部门 -> 薪资降序 -> 入职时间升序):")
	printBestPracticeEmployees(employees)
}

// 4. 字符串自然排序
func demonstrateNaturalStringSorting() {
	fmt.Println("4. 字符串自然排序:")
	fmt.Println()

	files := []string{
		"file1.txt", "file10.txt", "file2.txt", "file20.txt",
		"file3.txt", "file11.txt", "file100.txt", "file21.txt",
	}

	fmt.Println("原始文件列表:")
	for _, file := range files {
		fmt.Printf("  %s\n", file)
	}

	// 普通字符串排序
	files1 := make([]string, len(files))
	copy(files1, files)
	sort.Strings(files1)

	fmt.Println("\n普通字符串排序:")
	for _, file := range files1 {
		fmt.Printf("  %s\n", file)
	}

	// 自然排序（数字部分按数值排序）
	files2 := make([]string, len(files))
	copy(files2, files)
	sort.Slice(files2, func(i, j int) bool {
		return naturalLess(files2[i], files2[j])
	})

	fmt.Println("\n自然排序:")
	for _, file := range files2 {
		fmt.Printf("  %s\n", file)
	}
	fmt.Println()
}

// 5. 排序稳定性的实际应用
func demonstrateSortingStabilityUsage() {
	fmt.Println("5. 排序稳定性的实际应用:")
	fmt.Println()

	// 学生成绩数据，已按姓名排序
	type StudentGrade struct {
		Name  string
		Score int
		Class string
	}

	students := []StudentGrade{
		{"Alice", 85, "A班"},
		{"Bob", 90, "B班"},
		{"Charlie", 85, "A班"},
		{"David", 90, "A班"},
		{"Eve", 85, "B班"},
	}

	fmt.Println("原始数据（已按姓名排序）:")
	for _, student := range students {
		fmt.Printf("  %s: %d分 (%s)\n", student.Name, student.Score, student.Class)
	}

	// 非稳定排序按分数
	students1 := make([]StudentGrade, len(students))
	copy(students1, students)
	sort.Slice(students1, func(i, j int) bool {
		return students1[i].Score > students1[j].Score
	})

	fmt.Println("\n非稳定排序按分数（可能改变相同分数学生的相对顺序）:")
	for _, student := range students1 {
		fmt.Printf("  %s: %d分 (%s)\n", student.Name, student.Score, student.Class)
	}

	// 稳定排序按分数
	students2 := make([]StudentGrade, len(students))
	copy(students2, students)
	sort.SliceStable(students2, func(i, j int) bool {
		return students2[i].Score > students2[j].Score
	})

	fmt.Println("\n稳定排序按分数（保持相同分数学生的原始顺序）:")
	for _, student := range students2 {
		fmt.Printf("  %s: %d分 (%s)\n", student.Name, student.Score, student.Class)
	}
	fmt.Println()
}

// 6. 大数据排序优化
func demonstrateLargeDataSorting() {
	fmt.Println("6. 大数据排序优化技巧:")
	fmt.Println()

	fmt.Println("大数据排序优化策略:")
	fmt.Println("  1. 预分配内存:")
	fmt.Println("     - 避免排序过程中的内存重分配")
	fmt.Println("     - 使用 make([]Type, 0, capacity)")
	fmt.Println()

	fmt.Println("  2. 减少比较函数复杂度:")
	fmt.Println("     - 预计算排序键")
	fmt.Println("     - 避免字符串操作和复杂计算")
	fmt.Println()

	fmt.Println("  3. 选择合适的排序算法:")
	fmt.Println("     - 基础类型使用专用函数")
	fmt.Println("     - 不需要稳定性时避免使用 SliceStable")
	fmt.Println()

	fmt.Println("  4. 考虑并行处理:")
	fmt.Println("     - 分块排序后合并")
	fmt.Println("     - 使用 goroutine 并行处理")
	fmt.Println()

	fmt.Println("  5. 内存优化:")
	fmt.Println("     - 原地排序 vs 复制排序")
	fmt.Println("     - 及时释放不需要的内存")
	fmt.Println()
}

// 7. 排序错误处理和边界情况
func demonstrateSortingEdgeCases() {
	fmt.Println("7. 排序错误处理和边界情况:")
	fmt.Println()

	fmt.Println("常见边界情况处理:")

	// 空切片
	var emptySlice []int
	sort.Ints(emptySlice)
	fmt.Printf("  空切片排序: %v (长度: %d)\n", emptySlice, len(emptySlice))

	// 单元素切片
	singleElement := []int{42}
	sort.Ints(singleElement)
	fmt.Printf("  单元素切片: %v\n", singleElement)

	// 相同元素切片
	sameElements := []int{5, 5, 5, 5}
	sort.Ints(sameElements)
	fmt.Printf("  相同元素切片: %v\n", sameElements)

	// nil 切片处理
	var nilSlice []int
	sort.Ints(nilSlice) // 不会panic
	fmt.Printf("  nil切片排序: %v (长度: %d)\n", nilSlice, len(nilSlice))

	fmt.Println()
	fmt.Println("错误处理建议:")
	fmt.Println("  1. 排序前检查切片是否为nil或空")
	fmt.Println("  2. 比较函数要处理相等情况")
	fmt.Println("  3. 注意浮点数的NaN和Inf值")
	fmt.Println("  4. 字符串排序考虑Unicode和本地化")
	fmt.Println("  5. 时间排序注意时区问题")
	fmt.Println()
}

// 辅助函数
func printBestPracticeEmployees(employees []BestPracticeEmployee) {
	for _, emp := range employees {
		fmt.Printf("  %s - %s - ¥%.0f - %s\n",
			emp.Name, emp.Department, emp.Salary, emp.HireDate.Format("2006-01-02"))
	}
	fmt.Println()
}

func printBestPracticeOrders(orders []BestPracticeOrder) {
	for _, order := range orders {
		fmt.Printf("  %s: ¥%.0f (优先级:%d, %s) - %s\n",
			order.ID, order.Amount, order.Priority, order.Status,
			order.Date.Format("2006-01-02 15:04"))
	}
	fmt.Println()
}

// 自然排序比较函数
func naturalLess(a, b string) bool {
	// 简化的自然排序实现
	// 实际项目中可能需要更复杂的实现

	i, j := 0, 0
	for i < len(a) && j < len(b) {
		// 如果都是数字，按数值比较
		if isDigit(a[i]) && isDigit(b[j]) {
			numA, nextI := extractNumber(a, i)
			numB, nextJ := extractNumber(b, j)

			if numA != numB {
				return numA < numB
			}
			i, j = nextI, nextJ
		} else {
			// 按字符比较
			if a[i] != b[j] {
				return a[i] < b[j]
			}
			i++
			j++
		}
	}

	return len(a) < len(b)
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func extractNumber(s string, start int) (int, int) {
	end := start
	for end < len(s) && isDigit(s[end]) {
		end++
	}

	num, _ := strconv.Atoi(s[start:end])
	return num, end
}

/*
Go语言排序最佳实践总结:

1. 性能优化:
   - 预计算排序键，避免重复计算
   - 基础类型使用专用排序函数
   - 大数据考虑并行处理

2. 业务逻辑:
   - 复杂排序规则要清晰分层
   - 多级排序按重要性排序
   - 考虑业务场景的特殊需求

3. 稳定性:
   - 需要保持相对顺序时使用稳定排序
   - 多次排序时考虑稳定性影响
   - 性能敏感场景权衡稳定性需求

4. 错误处理:
   - 处理边界情况（空、nil、单元素）
   - 比较函数要健壮
   - 考虑特殊值（NaN、Inf等）

5. 代码质量:
   - 排序逻辑要清晰易懂
   - 复杂比较函数考虑提取为独立函数
   - 添加适当的注释说明排序规则

6. 实际应用:
   - 根据数据特征选择算法
   - 测量实际性能
   - 平衡性能和可维护性
*/
