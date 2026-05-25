package algorithms

import (
	"errors"
	"fmt"
	"lab5/structures"
	"strconv"
)

func BuildExpressionTree(tokens []string) (*structures.Node, error) {
	stack := structures.NewStack()

	for _, token := range tokens {

		if isOperator(token) {

			right, err := stack.Pop()
			if err != nil {
				return &structures.Node{}, err
			}

			left, err := stack.Pop()
			if err != nil {
				return &structures.Node{}, err
			}

			node := &structures.Node{
				Value: token,
				Left:  left,
				Right: right,
			}

			stack.Push(node)
		} else {

			node := &structures.Node{
				Value: token,
				Left:  nil,
				Right: nil,
			}
			stack.Push(node)
		}
	}

	if stack.Size() != 1 {
		return &structures.Node{}, errors.New("incorrect postfix expression")
	}

	res, err := stack.Peek()
	if err != nil {
		return &structures.Node{}, err
	}

	return res, nil
}

func Evaluate(root *structures.Node) (float64, error) {
	if root == nil {
		return 0, errors.New("tree is empty")
	}

	if root.Left == nil && root.Right == nil {
		return strconv.ParseFloat(root.Value, 64)
	}

	leftVal, err := Evaluate(root.Left)
	if err != nil {
		return 0, err
	}
	rightVal, err := Evaluate(root.Right)
	if err != nil {
		return 0, err
	}

	switch root.Value {
	case "+":
		return leftVal + rightVal, nil
	case "-":
		return leftVal - rightVal, nil
	case "*":
		return leftVal * rightVal, nil
	case "/":
		if rightVal == 0 {
			return 0, errors.New("dividing by zero")
		}
		return leftVal / rightVal, nil
	default:
		return 0, fmt.Errorf("unknown operator: %s", root.Value)
	}
}

func ToInfix(root *structures.Node) string {
	if root == nil {
		return ""
	}

	if !isOperator(root.Value) {
		return root.Value
	}

	leftExpr := ToInfix(root.Left)
	rightExpr := ToInfix(root.Right)

	leftNeedsParen := false
	if isOperator(root.Left.Value) && precedence(root.Left.Value) < precedence(root.Value) {
		leftNeedsParen = true
	}

	if isOperator(root.Left.Value) && precedence(root.Left.Value) == precedence(root.Value) && (root.Value == "-" || root.Value == "/") {
		leftNeedsParen = true
	}

	rightNeedsParen := false
	if isOperator(root.Right.Value) && precedence(root.Right.Value) < precedence(root.Value) {
		rightNeedsParen = true
	}
	if isOperator(root.Right.Value) && precedence(root.Right.Value) == precedence(root.Value) && (root.Value == "-" || root.Value == "/") {
		rightNeedsParen = true
	}

	if leftNeedsParen {
		leftExpr = "(" + leftExpr + ")"
	}
	if rightNeedsParen {
		rightExpr = "(" + rightExpr + ")"
	}

	return leftExpr + " " + root.Value + " " + rightExpr
}

func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/"
}

func precedence(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	}
	return 0
}
