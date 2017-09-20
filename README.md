
golang并发数控制工具

示例：

```go
// 初始化，最大并发数64
connPool := pool.NewPool(64)
for _, task := range tasks {
	connPool.Get()
	go func(task string) {
		fmt.Println(task)
		connPool.Put()
	}(task)
}
// 等待所有任务执行结束
connPool.Wait()
```
