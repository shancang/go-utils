package go_utils

import "sync"

type WaitGroup struct {
	size      int
	pool      chan struct{}
	waitCount int64
	waitGroup sync.WaitGroup
}

func NewWaitGroup(size int) *WaitGroup {
	wg := &WaitGroup{
		size: size,
	}
	if size > 0 {
		wg.pool = make(chan struct{}, wg.size)
	}
	return wg
}

func (wg *WaitGroup) Add() {
	if wg.size > 0 {
		wg.pool <- struct{}{}
	}
	wg.waitGroup.Add(1)
}

// Done 代表一个并发结束
func (wg *WaitGroup) Done() {
	if wg.size > 0 {
		<-wg.pool
	}
	wg.waitGroup.Done()
}

// Wait 等待所有并发goroutine结束
func (wg *WaitGroup) Wait() {
	wg.waitGroup.Wait()
}

// PendingCount 返回所有pending状态的goroutine数量
func (wg *WaitGroup) PendingCount() int64 {
	return int64(len(wg.pool))
}
