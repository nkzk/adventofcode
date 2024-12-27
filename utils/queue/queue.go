package queue

import "fmt"

type FIFO[T any] struct {
	items []T
}

func (queue *FIFO[T]) Enqueue(item T) {
	queue.items = append(queue.items, item)
}

func (queue *FIFO[T]) Dequeue() (T, error) {
	var item T

	if len(queue.items) > 0 {
		item = queue.items[0]
		queue.items = queue.items[1:]
		return item, nil
	}

	return item, fmt.Errorf("queue is empty")
}

func (queue *FIFO[T]) Peek() (T, error) {
	var zero T
	if len(queue.items) > 0 {
		return queue.items[0], nil
	}

	return zero, fmt.Errorf("queue is empty")
}

func (queue *FIFO[T]) IsEmpty() bool {
	return len(queue.items) == 0
}

type LIFO[T any] struct {
	items []T
}

func (queue *LIFO[T]) Enqueue(item T) {
	queue.items = append(queue.items, item)
}

func (queue *LIFO[T]) Dequeue() error {
	if len(queue.items) > 0 {
		queue.items = queue.items[:len(queue.items)-1]
		return nil
	}

	return fmt.Errorf("queue is empty")
}

func (queue *LIFO[T]) Peek() (T, error) {
	var zero T
	if len(queue.items) > 0 {
		return queue.items[0], nil
	}

	return zero, fmt.Errorf("queue is empty")
}

func (queue *LIFO[T]) IsEmpty() bool {
	return len(queue.items) == 0
}
