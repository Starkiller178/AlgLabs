package main

import (
	"fmt"
	"lab3/structures"
)

func main() {
	queue := structures.NewQueue(3)

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	fmt.Println(queue)

	if err := queue.Enqueue(40); err != nil {
		fmt.Printf("При переполнении возникает ошибка: %v\n", err)
	}

	fmt.Printf("Текущее количество элементов: %d\n", queue.Size())
	val, _ := queue.Front()
	fmt.Println(val)
	fmt.Printf("Количество элементов не изменилось: %d\n", queue.Size())

	val, _ = queue.Dequeue()
	fmt.Println(val)
	fmt.Printf("Количество элементов изменилось: %d\n", queue.Size())

	queue.Dequeue()
	queue.Dequeue()

	fmt.Println(queue)

	if _, err := queue.Dequeue(); err != nil {
		fmt.Printf("При попытке удатения из пустой очереди возникает ошибка: %v\n", err)
	}
}
