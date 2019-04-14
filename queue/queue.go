package queue

import (
    "errors"
)

var (
    ErrQueueEmpty = errors.New("queue empty")
    ErrQueueFull = errors.New("queue full")
)

type queueItem interface {}

type queue struct {
    list []queueItem
    size uint64
}

type Queuer interface {
    Dequeue() (queueItem, error)
    Enqueue(i queueItem) error
    Peek() (queueItem, error)
    //Count() uint64
}

func NewQueuer(size uint64) Queuer {
    return &queue{
        list: []queueItem{},
        size: size,
    }
}

func (q *queue) Dequeue() (queueItem, error) {
    if len(q.list) == 0 {
        return nil, ErrQueueEmpty
    }
    item := q.list[0]
    q.list = q.list[1:]
    return item, nil
}

func (q *queue) Enqueue(i queueItem) error {
    if uint64(len(q.list)) >= q.size {
        return ErrQueueFull
    }

    q.list = append(q.list, i)
    return nil
}

func (q *queue) Peek() (queueItem, error) {
    if len(q.list) == 0 {
        return nil, ErrQueueEmpty
    }
    return q.list[0], nil
}
