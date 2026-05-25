package main

import (
	"fmt"
	"lab5/algorithms"
	"strings"
)

func main() {
	token := strings.Fields("2 3 + 4 5 * +")
	tree, err := algorithms.BuildExpressionTree(token)
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := algorithms.Evaluate(tree)
	if err != nil {
		fmt.Println(err)
	}
	InfixForm := algorithms.ToInfix(tree)
	fmt.Printf("Постфиксная форма дерева: %s\n", strings.Join(token, " "))
	fmt.Printf("Постфиксная форма дерева: %s\n", InfixForm)
	fmt.Printf("Результат выражения: %.2f\n", res)
}
