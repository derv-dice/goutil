package main

import (
	"context"
	"sync"
)

type Task interface {
	Do()
}

// WPool - workers pool
type WPool struct {
	ctx context.Context

	count int
	in    chan Task
	ins   []chan Task
	out   chan Task
	stop  context.CancelFunc
	wg    sync.WaitGroup

	enabled bool
}

func NewWPool(ctx context.Context, count int) *WPool {
	if count <= 0 {
		count = 1
	}

	childCtx, cancel := context.WithCancel(ctx)

	wp := &WPool{
		ctx:   childCtx,
		stop:  cancel,
		count: count,
		in:    make(chan Task),
		ins:   []chan Task{},
		out:   make(chan Task),
	}

	for i := 0; i < count; i++ {
		wp.ins = append(wp.ins, make(chan Task))
	}

	return wp
}

// Put - Добавление задачи в пул
func (w *WPool) Put(task Task) {
	w.in <- task
	w.wg.Add(1)
}

// Start - Включение работы пула задач.
// После этого можно использовать WPool.Put(t Task) чтобы добавлять задачи на выполнение
func (w *WPool) Start() {
	if w.enabled {
		return
	}

	// Распределение задач между каналами
	go func() {
		for {
			for _, c := range w.ins {
				select {
				case <-w.ctx.Done():
					return
				case task := <-w.in:
					c <- task
				}
			}
		}
	}()

	// Обработка задач из каналов
	for i := 0; i < w.count; i++ {
		go func(ctx context.Context, i int) {
			for {
				select {
				case <-ctx.Done():
					return
				case task := <-w.ins[i]:
					task.Do()
					w.wg.Done()
				}
			}
		}(w.ctx, i)
	}

	w.enabled = true
	return
}

// Stop - Ждет, пока оставшиеся задачи закончат свое выполнение и выключает работу пула.
func (w *WPool) Stop() {
	if !w.enabled {
		return
	}
	w.wg.Wait()

	w.stop()
	w.enabled = false
}
