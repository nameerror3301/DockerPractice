package queue

import "sync"

type Queue struct {
	mux     sync.Mutex
	first   Node
	current Node
	size    int
}

type Node interface {
	SetNext(n Node)
	GetNext() Node
}

func NewQueue() *Queue {
	return &Queue{
		size:    0,
		first:   nil,
		current: nil,
	}
}

// Add data in queue
func (qu *Queue) Push(data Node) {
	qu.mux.Lock()
	defer qu.mux.Unlock()

	d := data //.(*Node)
	//First node
	if qu.first == nil {
		qu.first = d
		qu.current = d
		qu.size++
		return
	}

	qu.current.SetNext(d)
	qu.current = d
	qu.size++
}

// Give data from queue
func (qu *Queue) Pop() (int, Node) {
	// var ok bool

	if qu.size == 0 {
		return 0, nil
	}

	qu.mux.Lock()
	defer qu.mux.Unlock()

	data := qu.first
	sequence := qu.size

	// if qu.first, ok = qu.current.GetNext(); !ok {
	// 	return 0, nil, nil
	// }
	qu.first = qu.current.GetNext()
	qu.size--

	return sequence, data
}
