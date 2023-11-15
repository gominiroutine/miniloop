# miniloop

```go
// USAGE DEMO
/***********
	var ctx = context.Background()
	var count = 0
	var ctxKey = "ContextTaskId"
	var shopIdParam = time.Now().Unix()
	miniloop.AddTask(
		ctx,
		shopIdParam,
		func(tk *time.Ticker, taskId int64) {
			fmt.Println("Working...", count, taskId)
			count++
			if count > 9 {
				tk.Reset(time.Second * 3)
			}
			if count > 12 {
				delete(miniloop.MapRunningCtx, taskId)
				fmt.Println("Stopped", count, taskId)
			}
		},
		ctxKey,
	)
***********/
// USAGE
```
