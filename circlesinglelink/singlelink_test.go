package main

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

//使用链表实现队列入列和出列
//需要改动的地方 ，就是出列是从第一个出列

type CircleLink struct {
	no   int
	name string
	next *CircleLink
}

func PushNode(head *CircleLink, node *CircleLink) {
	//先判断是不是第一只猫
	if head.next == nil {
		head.name = node.name
		head.no = node.no
		head.next = head
		return
	}
	tmp := head
	for {
		if tmp.next == head {
			break
		}
		tmp = tmp.next
	}

	tmp.next = node
	node.next = head
}

func PopNode(head *CircleLink) (*CircleLink, *CircleLink) {
	tmp := head
	helper := head
	//判断是否是空链表
	if head.next == nil {
		fmt.Println(`链表为空`)
		return tmp, head
	}

	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}
	helper.next = tmp.next
	head = tmp.next

	return tmp, head
}

var (
	lock  sync.Mutex
	count int
)

func TestSingleLink(t *testing.T) {
	head := &CircleLink{}
	getQueue := &CircleLink{}
	go func() {
		for {

			select {
			case <-time.After(time.Duration(rand.Intn(5)) * time.Second):
				lock.Lock()
				getQueue, head = PopNode(head)
				if getQueue.no != 0 {
					fmt.Println("1号协程服务-->", getQueue.no, "顾客", head)
				}

				getQueue.no = 0
				getQueue.next = nil
				lock.Unlock()
			}
		}
	}()

	go func() {
		for {
			select {
			case <-time.After(time.Duration(rand.Intn(5)) * time.Second):
				lock.Lock()
				getQueue, head = PopNode(head)
				if getQueue.no != 0 {
					fmt.Println("2号协程服务-->", getQueue.no, "顾客", head)
				}
				getQueue.no = 0
				getQueue.next = nil

				lock.Unlock()
			}
		}
	}()

	for {
		select {
		case <-time.After(time.Duration(rand.Intn(5)) * time.Second):
			lock.Lock()
			count++
			tmp := &CircleLink{
				no:   count,
				name: "",
				next: nil,
			}
			PushNode(head, tmp)
			lock.Unlock()
		}
	}
}
