package entity

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type MyMap struct {
	dict map[string]int
}

func NewMyMap() *MyMap {
	return &MyMap{
		dict: make(map[string]int),
	}
}

func (d *MyMap) CountWords() error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите строку: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		return errors.New("reading err")
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	words := strings.Fields(text)

	for _, w := range words {
		cleaned := cleanWord(w)
		if cleaned == "" {
			continue
		}
		d.dict[cleaned]++
	}

	return nil
}

func (d *MyMap) CountWordsF(filename string) error {

	file, err := os.Open(filename)
	if err != nil {
		return errors.New("file opening err")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)

		for _, w := range words {
			cleaned := cleanWord(w)
			if cleaned == "" {
				continue
			}
			d.dict[cleaned]++
		}
	}

	if err := scanner.Err(); err != nil {
		return errors.New("file reading err")
	}

	return nil
}

func cleanWord(w string) string {
	var result strings.Builder

	for _, ch := range w {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			result.WriteRune(ch)
		}
	}

	return result.String()
}

func (d *MyMap) PrintCountWords() {
	if len(d.dict) == 0 {
		fmt.Println("Строка не содержит слов")
		return
	}

	for word, count := range d.dict {
		fmt.Printf("%s -> %d\n", word, count)
	}
}

func (d *MyMap) GetMostFrequentWord() (string, error) {
	if len(d.dict) == 0 {
		return "", errors.New("Err map is empty")
	}

	maxCnt := 0
	resWord := ""
	flag := true
	lastCount := 0
	for word, count := range d.dict {
		if resWord == "" {
			lastCount = count
		}

		if count != lastCount {
			flag = false
			lastCount = count
		}
		if count > maxCnt {
			resWord = word
			maxCnt = count
		}
	}

	if len(d.dict) > 1 && flag == true {
		return "Все слова встречаются одинаковое количество раз", nil
	}

	return resWord, nil
}

func (d *MyMap) PrintNotUniqWords() {
	if len(d.dict) == 0 {
		fmt.Println("Строка не содержит слов")
		return
	}

	cnt := 0
	for word, count := range d.dict {
		if count > 1 {
			fmt.Printf("%s -> %d\n", word, count)
			cnt++
		}
	}

	if cnt == 0 {
		fmt.Println("Все слова встречаются по одному разу")
	}
}

func (d *MyMap) GetCountUniqWords() (int, error) {
	if len(d.dict) == 0 {
		return 0, errors.New("Err map is empty")
	}

	cntUniqWords := 0
	for _, count := range d.dict {
		if count == 1 {
			cntUniqWords++
		}
	}

	return cntUniqWords, nil
}
