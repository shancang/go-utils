package goutils

import "sync"

type WaitGroup struct {
	size      int //并发大小
	pool      chan struct{}
	waitCount int64
	waitGroup sync.WaitGroup
}

// NewWaitGroup 创建一个新的 WaitGroup 对象
// 参数 size: WaitGroup 的大小
// 返回值: 创建的 WaitGroup 对象的指针
func NewWaitGroup(size int) *WaitGroup {
	wg := &WaitGroup{
		size: size,
	}
	if size > 0 {
		wg.pool = make(chan struct{}, wg.size)
	}
	return wg
}

// Add 向 WaitGroup 中添加一个任务
// 如果 WaitGroup 的大小大于 0，则向 WaitGroup 的池中添加一个 struct{}
// 然后调用 waitGroup 的 Add 方法
func (wg *WaitGroup) Add() {
	if wg.size > 0 {
		wg.pool <- struct{}{}
	}
	wg.waitGroup.Add(1)
}

// Done 表示一个任务已完成
// 如果 WaitGroup 的大小大于 0，则从 WaitGroup 的池中取出一个 struct{}
// 然后调用 waitGroup 的 Done 方法
func (wg *WaitGroup) Done() {
	if wg.size > 0 {
		<-wg.pool
	}
	wg.waitGroup.Done()
}

// Wait 等待所有的任务完成
// 调用 waitGroup 的 Wait 方法等待所有的任务完成
func (wg *WaitGroup) Wait() {
	wg.waitGroup.Wait()
}

// PendingCount 获取当前等待的任务数量
// 返回值: 当前等待的任务数量
func (wg *WaitGroup) PendingCount() int64 {
	return int64(len(wg.pool))
}
