package main

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

var (
	lock  sync.Mutex
	queue *CircleQueue
	count int
)

//
func TestQueue(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	lock = sync.Mutex{}
	queue := &CircleQueue{
		maxSize: 10,
		head:    0,
		tail:    0,
	}

	go func() {
		for {

			select {
			case <-time.After(time.Duration(rand.Intn(5)) * time.Second):
				lock.Lock()
				getQueue, err := queue.Pop()
				if err != nil {
					fmt.Println("从队列中取出数据失败", err)
				} else {
					fmt.Println("1号协程服务-->", getQueue, "顾客")
				}

				lock.Unlock()
			}
		}
	}()

	go func() {
		for {
			select {
			case <-time.After(time.Duration(rand.Intn(5)) * time.Second):
				lock.Lock()
				getQueue, err := queue.Pop()
				if err != nil {
					fmt.Println("从队列中取出数据失败", err)
				} else {
					fmt.Println("2号协程服务-->", getQueue, "顾客")
				}

				lock.Unlock()
			}
		}
	}()

	for {
		select {
		case <-time.After(time.Duration(rand.Intn(5)) * time.Second):
			lock.Lock()
			count++
			if err := queue.Push(count); err != nil {
				fmt.Println("添加队列失败", err)
			} else {
				fmt.Println(count, "排队成功")
			}

			lock.Unlock()
		}
	}

}
