package goutils

import (
	"sync/atomic"
)

type AlgorithmInterface interface {
	Next() string
	AddItem(string)
	RemoveItemAt(int)
}

// RoundRobin 泛型结构体表示一个轮询状态
type RoundRobin[T any] struct {
	items []T          // 泛型候选项
	index atomic.Int64 // 原子操作的当前索引
}

// NewRoundRobin 创建一个新的 RoundRobin 实例
func NewRoundRobin[T any](items []T) *RoundRobin[T] {
	return &RoundRobin[T]{
		items: items,
	}
}

// Next 返回下一个候选项，并更新索引
func (r *RoundRobin[T]) Next() T {
	if len(r.items) == 0 {
		var zeroValue T
		return zeroValue // 如果没有候选项，返回类型的零值
	}
	// 原子增加索引，并确保索引循环回到切片范围内
	idx := r.index.Add(1) % int64(len(r.items))
	return r.items[idx]
}

// AddItem 添加一个新的候选项
func (r *RoundRobin[T]) AddItem(item T) {
	r.items = append(r.items, item)
}

// RemoveItemAt 移除指定索引的候选项
func (r *RoundRobin[T]) RemoveItemAt(index int) {
	if index >= 0 && index < len(r.items) {
		r.items = append(r.items[:index], r.items[index+1:]...)
	}
}
