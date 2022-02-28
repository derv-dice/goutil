package main

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type TestTask struct{}

func (t *TestTask) Do() {
	time.Sleep(time.Second * 1)
	successCounter++
}

var successCounter = 0

func TestNewWPool(t *testing.T) {
	wp := NewWPool(context.Background(), 10)
	wp.Start()

	clock := time.Now()
	wp.Put(new(TestTask))
	wp.Put(new(TestTask))
	wp.Put(new(TestTask))
	wp.Put(new(TestTask))
	wp.Put(new(TestTask))
	wp.Stop()

	check := time.Since(clock)

	assert.Equal(t, 5, successCounter)   // Все задачи выполнены
	assert.Less(t, check, time.Second*2) // Общее время выполнения должно быть меньше 2 секунд (будет 1с)
	// В 1 поток было бы минимум 5с.
}
