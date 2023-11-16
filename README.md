![gomini.png](gomini.png)

```go
// USAGE DEMO
/***********
	var ctx = context.Background()
	var count = 0
	var ctxKey = "ContextTaskId"
	var taskIdParam = int64(123456789)
	miniloop.AddTask(
		ctx,
		taskIdParam,
		func(tk *time.Ticker, taskId int64) {
			fmt.Println("Working...", count, taskId)
			count++
			if count > 9 {
				tk.Reset(time.Second * 3)
			}
			if count > 12 {
				miniloop.DeleteTask(taskId)
				fmt.Println("Stopped", count, taskId)
			}
		},
		ctxKey,
	)
***********/
// USAGE
```
