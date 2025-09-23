package main

import (
	"fmt"
	"sync"
	"time"
)

/* 题目1：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
考察点：go 关键字的使用、协程的并发执行。
题目2：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
考察点：协程原理、并发任务调度。 */

// 题目1：协程打印奇偶数
func printOddEvenNumbers() {
	fmt.Println("=== 题目1：协程打印奇偶数 ===")
	var wg sync.WaitGroup
	wg.Add(2)

	// 协程1：打印奇数 (1, 3, 5, 7, 9)
	go func() {
		defer wg.Done()
		fmt.Println("奇数协程开始执行...")
		for i := 1; i <= 10; i++ {
			if i%2 != 0 {
				fmt.Printf("奇数: %d\n", i)
				time.Sleep(100 * time.Millisecond) // 模拟处理时间
			}
		}
		fmt.Println("奇数协程执行完成")
	}()

	// 协程2：打印偶数 (2, 4, 6, 8, 10)
	go func() {
		defer wg.Done()
		fmt.Println("偶数协程开始执行...")
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Printf("偶数: %d\n", i)
				time.Sleep(120 * time.Millisecond) // 模拟处理时间
			}
		}
		fmt.Println("偶数协程执行完成")
	}()

	// 等待所有协程完成
	wg.Wait()
	fmt.Println("所有协程执行完成\n")
}

// Task 表示一个任务
type Task struct {
	ID   int
	Name string
	Func func() interface{}
}

// TaskResult 表示任务执行结果
type TaskResult struct {
	TaskID      int
	TaskName    string
	Result      interface{}
	ExecuteTime time.Duration
	Error       error
}

// TaskScheduler 任务调度器
type TaskScheduler struct {
	tasks   []Task
	results []TaskResult
	mu      sync.Mutex
}

// NewTaskScheduler 创建新的任务调度器
func NewTaskScheduler() *TaskScheduler {
	return &TaskScheduler{
		tasks:   make([]Task, 0),
		results: make([]TaskResult, 0),
	}
}

// AddTask 添加任务
func (ts *TaskScheduler) AddTask(id int, name string, taskFunc func() interface{}) {
	task := Task{
		ID:   id,
		Name: name,
		Func: taskFunc,
	}
	ts.tasks = append(ts.tasks, task)
}

// ExecuteTasks 并发执行所有任务
func (ts *TaskScheduler) ExecuteTasks() {
	fmt.Println("=== 题目2：任务调度器并发执行 ===")
	var wg sync.WaitGroup
	resultChan := make(chan TaskResult, len(ts.tasks))

	// 启动协程执行每个任务
	for _, task := range ts.tasks {
		wg.Add(1)
		go func(t Task) {
			defer wg.Done()
			fmt.Printf("任务 [%s] 开始执行...\n", t.Name)
			
			startTime := time.Now()
			var result interface{}
			var err error
			
			// 执行任务并捕获可能的panic
			func() {
				defer func() {
					if r := recover(); r != nil {
						err = fmt.Errorf("任务执行panic: %v", r)
					}
				}()
				result = t.Func()
			}()
			
			executeTime := time.Since(startTime)
			
			taskResult := TaskResult{
				TaskID:      t.ID,
				TaskName:    t.Name,
				Result:      result,
				ExecuteTime: executeTime,
				Error:       err,
			}
			
			resultChan <- taskResult
			fmt.Printf("任务 [%s] 执行完成，耗时: %v\n", t.Name, executeTime)
		}(task)
	}

	// 等待所有任务完成
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// 收集结果
	for result := range resultChan {
		ts.mu.Lock()
		ts.results = append(ts.results, result)
		ts.mu.Unlock()
	}

	fmt.Println("\n所有任务执行完成")
}

// PrintResults 打印任务执行结果统计
func (ts *TaskScheduler) PrintResults() {
	fmt.Println("\n=== 任务执行结果统计 ===")
	totalTime := time.Duration(0)
	successCount := 0
	errorCount := 0

	for _, result := range ts.results {
		totalTime += result.ExecuteTime
		if result.Error != nil {
			errorCount++
			fmt.Printf("❌ 任务ID: %d, 名称: %s, 状态: 失败, 耗时: %v, 错误: %v\n",
				result.TaskID, result.TaskName, result.ExecuteTime, result.Error)
		} else {
			successCount++
			fmt.Printf("✅ 任务ID: %d, 名称: %s, 状态: 成功, 耗时: %v, 结果: %v\n",
				result.TaskID, result.TaskName, result.ExecuteTime, result.Result)
		}
	}

	fmt.Printf("\n📊 执行统计:\n")
	fmt.Printf("   总任务数: %d\n", len(ts.results))
	fmt.Printf("   成功任务: %d\n", successCount)
	fmt.Printf("   失败任务: %d\n", errorCount)
	fmt.Printf("   总耗时: %v\n", totalTime)
	if len(ts.results) > 0 {
		fmt.Printf("   平均耗时: %v\n", totalTime/time.Duration(len(ts.results)))
	}
}

// 示例任务函数
func calculateSum(n int) func() interface{} {
	return func() interface{} {
		sum := 0
		for i := 1; i <= n; i++ {
			sum += i
			time.Sleep(10 * time.Millisecond) // 模拟计算时间
		}
		return sum
	}
}

func calculateFactorial(n int) func() interface{} {
	return func() interface{} {
		if n < 0 {
			panic("负数无法计算阶乘")
		}
		result := 1
		for i := 1; i <= n; i++ {
			result *= i
			time.Sleep(15 * time.Millisecond) // 模拟计算时间
		}
		return result
	}
}

func simulateNetworkRequest(url string) func() interface{} {
	return func() interface{} {
		// 模拟网络请求
		time.Sleep(time.Duration(200+len(url)*10) * time.Millisecond)
		return fmt.Sprintf("响应来自: %s", url)
	}
}

func main() {
	// 执行题目1
	printOddEvenNumbers()

	// 执行题目2
	scheduler := NewTaskScheduler()

	// 添加各种类型的任务
	scheduler.AddTask(1, "计算1到100的和", calculateSum(100))
	scheduler.AddTask(2, "计算5的阶乘", calculateFactorial(5))
	scheduler.AddTask(3, "模拟网络请求1", simulateNetworkRequest("https://api.example1.com"))
	scheduler.AddTask(4, "计算1到50的和", calculateSum(50))
	scheduler.AddTask(5, "计算7的阶乘", calculateFactorial(7))
	scheduler.AddTask(6, "模拟网络请求2", simulateNetworkRequest("https://api.example2.com/data"))

	// 并发执行所有任务
	scheduler.ExecuteTasks()

	// 打印执行结果统计
	scheduler.PrintResults()
}
