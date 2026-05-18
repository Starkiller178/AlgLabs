package structures

import (
	"errors"
)

type Queue struct {
	input   Stack
	output  Stack
	maxSize int
}

func NewQueue(maxSize int) *Queue {
	return &Queue{
		maxSize: maxSize,
	}
}

func (q *Queue) GetInput() *Stack {
	return &q.input
}

func (q *Queue) GetOutput() *Stack {
	return &q.output
}

func (q *Queue) Enqueue(x int) error {

	if q.Size() >= q.maxSize {
		return errors.New("Err queue overflow")
	}
	q.input.Push(x)
	return nil
}

func (q *Queue) Dequeue() (int, error) {

	if q.output.IsEmpty() {
		for !q.input.IsEmpty() {
			val, _ := q.input.Pop()
			q.output.Push(val)
		}
	}

	if q.output.IsEmpty() {
		return 0, errors.New("Err queue is empty")
	}

	return q.output.Pop()
}

func (q *Queue) Front() (int, error) {

	if q.output.IsEmpty() {
		for !q.input.IsEmpty() {
			val, _ := q.input.Pop()
			q.output.Push(val)
		}
	}

	if q.output.IsEmpty() {
		return 0, errors.New("Err queue is empty")
	}

	return q.output.Peek()
}

func (q *Queue) Size() int {
	return q.input.Size() + q.output.Size()
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}
