// 解释器模式（Interpreter Pattern）提供了评估语言的语法或表达式的方式，它属于行为型模式。这种模式实现了一个表达式接口，该接口解释一个特定的上下文。这种模式被用在 SQL 解析、符号处理引擎等。
package main

import "fmt"

type iExpression interface {
	Interpret(string) bool
}

type ExpressionAnd struct {
	expr1 iExpression
	expr2 iExpression
}

func (e ExpressionAnd) Interpret(s string) bool {
	return e.expr1.Interpret(s) && e.expr2.Interpret(s)
}

type Expression []string

func (e Expression) Interpret(s string) bool {
	for _, ex := range e {
		if ex == s {
			return true
		}
	}
	return false
}

func main() {
	expr1 := Expression([]string{"A", "B"})
	expr2 := Expression([]string{"A", "C", "D"})

	expr := ExpressionAnd{expr1, expr2}
	fmt.Println(expr.Interpret("A"))
	fmt.Println(expr.Interpret("D"))
}
