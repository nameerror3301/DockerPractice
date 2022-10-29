package queue

import (
	"fmt"
	"testing"
)

var queue = NewQueue()

type Unit struct {
	Data string
	next *Unit
}

func (ut *Unit) SetNext(n Node) {
	ut.next = n.(*Unit)
}

func (ut *Unit) GetNext() Node {
	return ut.next
}

func Typization(all int, val interface{}) (int, *Unit, error) {
	if value, ok := val.(Unit); ok {
		return all, &value, nil
	}

	return all, nil, fmt.Errorf("filed unit")
}

func TestQueue(t *testing.T) {
	val1 := &Unit{Data: "Test1"}
	val2 := &Unit{Data: "Test2"}
	val3 := &Unit{Data: "Test3"}
	val4 := &Unit{Data: "Test4"}

	queue.Push(val1)
	queue.Push(val2)
	queue.Push(val3)

	all, val, err := Typization(queue.Pop())
	if err == nil && all == 3 && val.Data == val1.Data {
		t.Fatal("Read 1 value")
	}

	queue.Push(val4)

	testCounter := 2
	for count := 3; count != 0; count-- {

		all, val, err := Typization(queue.Pop())
		if err == nil && all == count && val.Data != fmt.Sprintf("Test%d", testCounter) {
			if err != nil {
				t.Error(err)
			}

			t.Fatalf("read %d value", count)
		}
		testCounter++
	}

	if idx, _ := queue.Pop(); idx != 0 {
		t.Fatalf("last index is not null %d", idx)
	}

}
