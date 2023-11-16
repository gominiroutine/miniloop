package miniloop

import (
	"context"
	"time"
)

var mapRunningCtx = map[int64]context.Context{}

type queueCtx struct {
	ctx    context.Context
	cancel context.CancelFunc
	ticker *time.Ticker
}

func (e *queueCtx) run(execFunc func(*time.Ticker, int64), ctxKey string) {
	if taskId, ok := e.ctx.Value(ctxKey).(int64); ok {
		if mapRunningCtx[taskId] == nil {
			e.cancel()
		} else {
			execFunc(e.ticker, taskId)
		}
	}
}

func start(ctxParent context.Context, execFunc func(*time.Ticker, int64), ctxKey string) {
	var ctx, cancel = context.WithCancel(ctxParent)
	r := &queueCtx{
		ctx:    ctx,
		cancel: cancel,
		ticker: time.NewTicker(time.Second),
	}
loop:
	for range r.ticker.C {
		select {
		case <-ctx.Done():
			break loop
		default:
			r.run(execFunc, ctxKey)
		}
	}
}

func AddTask(ctx context.Context, taskId int64, execFunc func(*time.Ticker, int64), ctxKey string) {
	if ctx == nil || taskId == 0 || execFunc == nil {
		return
	}
	if mapRunningCtx[taskId] == nil {
		mapRunningCtx[taskId] = context.WithValue(ctx, ctxKey, taskId)
		go start(mapRunningCtx[taskId], execFunc, ctxKey)
	}
}

func DeleteTask(taskId int64) {
	delete(mapRunningCtx, taskId)
}

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
