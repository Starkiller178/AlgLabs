package tests

import (
	"lab3/structures"
	"testing"
)

func TestPush(t *testing.T) {
	stack := structures.NewStack()
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	if stack.Size() != 3 {
		t.Error("Stask Push method is incorrect. Size must be = 3")
	}

	if (*stack.GetItems())[2] != 30 {
		t.Errorf("Stask Push method is incorrect. Expected: 30, received: %d", (*stack.GetItems())[2])
	}
}

func TestPop(t *testing.T) {
	stack := structures.NewStack()
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	val, _ := stack.Pop()
	if val != 30 {
		t.Errorf("Stask Pop method is incorrect. Expected: 30, received: %d", val)
	}

	if stack.Size() != 2 {
		t.Error("Stask Pop method is incorrect. Size must be = 2")
	}

	stack.Pop()
	stack.Pop()

	if _, err := stack.Pop(); err == nil {
		t.Error("Stack Pop method is incorrect. err must be not nil")
	}
}

func TestPeek(t *testing.T) {
	stack := structures.NewStack()
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	val, _ := stack.Peek()
	if val != 30 {
		t.Errorf("Stask Peek method is incorrect. Expected: 30, received: %d", val)
	}

	if stack.Size() != 3 {
		t.Error("Stask Peek method is incorrect. Size must be = 3")
	}

	stack.Pop()
	stack.Pop()
	stack.Pop()

	if _, err := stack.Peek(); err == nil {
		t.Error("Stack Peek method is incorrect. err must be not nil")
	}

}

func TestEnqueue(t *testing.T) {
	queue := structures.NewQueue(3)

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	if queue.GetInput().Size() != 3 {
		t.Error("Queue Enqueue method is incorrect")
	}
}

func TestDequeue(t *testing.T) {
	queue := structures.NewQueue(3)

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	val, _ := queue.Dequeue()

	if val != 10 {
		t.Errorf("Queue Dequeue method is incorrect. Expected: 10, received: %d", val)
	}

	queue.Dequeue()
	queue.Dequeue()

	if _, err := queue.Dequeue(); err == nil {
		t.Error("Queue Dequeue method is incorrect. err must be not nil")
	}

	if !queue.IsEmpty() {
		t.Error("Queue Dequeue method is incorrect. Queue must be empty")
	}
}

func TestFront(t *testing.T) {
	queue := structures.NewQueue(3)

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	val, _ := queue.Front()
	if val != 10 {
		t.Errorf("Queue Front method is incorrect. Expected: 10, received: %d", val)
	}

	if queue.Size() != 3 {
		t.Error("Queue Front method is incorrect. Size must be = 3")
	}

	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()

	if _, err := queue.Front(); err == nil {
		t.Error("Queue Front method is incorrect. err must be not nil")
	}
}
