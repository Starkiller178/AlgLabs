package main

import (
	"bufio"
	"fmt"
	"lab4/entity"
	"os"
)

func main() {
	var choice int
	dict := entity.NewMyMap()
	fmt.Println("Считать строку из консоли - 1")
	fmt.Println("Считать текст из файла MainTest - 2")
	fmt.Print("ВАШ ВЫБОР: ")
	fmt.Scan(&choice)
	bufio.NewReader(os.Stdin).ReadString('\n')

	switch choice {
	case 1:
		err := dict.CountWords()
		if err != nil {
			fmt.Println(err)
			return
		}
	case 2:
		err := dict.CountWordsF("MainTest")
		if err != nil {
			fmt.Println(err)
			return
		}
	default:
		fmt.Println("Непрвильный выбор!!!")
		return
	}

	fmt.Println("Найденные слова:")
	dict.PrintCountWords()

	word, err := dict.GetMostFrequentWord()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Самое частовстречающееся слово: %s\n", word)

	fmt.Println("Слова, встречающиеся более 1 раза:")
	dict.PrintNotUniqWords()

	count, err := dict.GetCountUniqWords()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Количество уникальных слов: %d", count)
}
