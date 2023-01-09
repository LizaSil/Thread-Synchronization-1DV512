package main

import (
	"fmt"
	"sync"
)

func Task_3() {
	// Create a message queue
	mq := &MessageQueue{}
	go mq.Send("A")
	go mq.Send("B")
	go mq.Send("C")
	go mq.Send("D")
	go mq.Receive()
	// Infinite loop
	select {}
}

type IMessageQueue interface {
	/**
	 * Send a message.
	 * @return true if successful, otherwise false.
	 */
	Send(msg string) bool
	/**
	 * Receive a message.
	 * If queue has no message, Recv() will block until one arrives.
	 * @return A message.
	 */
	Receive() string
}

type MessageQueue struct {
	// Max buffer size
	queue [5]string
	// Semaphore to manage access to the queue
	mut sync.Mutex
}

func (q *MessageQueue) Send(msg string) {
	for {
		q.mut.Lock()
		for i := range q.queue {
			if q.queue[i] == "" {
				q.queue[i] = msg
				break
			}
		}
		q.mut.Unlock()
	}
}

func (q *MessageQueue) Receive() {
	var msg string
	for {
		q.mut.Lock()
		for i := range q.queue {
			if q.queue[i] != "" {
				msg = q.queue[i]
				q.queue[i] = ""
				fmt.Print(msg)
				break
			}
		}
		q.mut.Unlock()
	}
}
