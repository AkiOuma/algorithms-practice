package stack

import (
	"errors"
	"strings"
)

type RPNStackItem struct {
	Key   string
	Value float64
}

// 运算符优先级计算
func SymbolRank(symbol string) (rank int, err error) {
	switch symbol {
	case "+":
		rank = 1
	case "-":
		rank = 1
	case "*":
		rank = 2
	case "/":
		rank = 2
	case "(":
		rank = -1
	default:
		err = errors.New("error: symbol can not be recognize")
	}
	return
}

// 中缀表达式转后缀表达式，并以切片的形式记录后缀表达式
func TranslateToRPN(expression string) (rpn []string) {
	items := strings.Split(expression, "")
	var symbols ArrStack
	for _, v := range items {
		// 若是运算符则进入判断逻辑
		if IsSymbol(v) {
			// 遇到运算符"("直接入栈
			if v == "(" {
				symbols.Push(v)
				continue
			}
			// 当遇到运算符")"时，则从符号栈中一直出栈，直到遇到"("为止，然后获取下一个符号
			if v == ")" {
				for symbol, _ := symbols.Pop(); symbol != "("; symbol, _ = symbols.Pop() {
					rpn = append(rpn, symbol.(string))
				}
				continue
			}
			// 获取栈顶，进行符号优先级比较，若空栈则直接推出循环，执行下一步的压栈操作
			for !symbols.IsEmpty() {
				stackTop, _ := symbols.Top()
				rankStackTop, _ := SymbolRank(stackTop.(string))
				rank, _ := SymbolRank(v)
				// 符号优先级比栈顶优先级要高，终止循环，进入下一步的压栈操作
				if rankStackTop < rank {
					break
				} else {
					// 栈顶优先级较高，执行一次出栈，并将结果放入后缀表达式切片，然后获取下一个栈顶与当前符号进行比较
					stackTopSymbol, _ := symbols.Pop()
					rpn = append(rpn, stackTopSymbol.(string))
				}
			}
			// 压栈
			symbols.Push(v)
		} else {
			// 若不是运算符则直接放入后缀表达式切片
			rpn = append(rpn, v)
		}
	}
	if !symbols.IsEmpty() {
		for !symbols.IsEmpty() {
			top, _ := symbols.Pop()
			rpn = append(rpn, top.(string))
		}
	}
	return
}

// 后缀表达式求值
func RPN(expression string, valueMap map[string]float64) float64 {
	var itemStack LinkStack
	RPNSlice := TranslateToRPN(expression)
	for _, v := range RPNSlice {
		// 不是运算符直接入栈
		if !IsSymbol(v) {
			itemStack.Push(RPNStackItem{Key: v, Value: valueMap[v]})
			continue
		}
		// 运算符，从栈中取值进行运算
		item1, _ := itemStack.Pop()
		item2, _ := itemStack.Pop()
		value1 := item1.(RPNStackItem).Value
		value2 := item2.(RPNStackItem).Value
		result := Caculate(value1, value2, v)
		itemStack.Push(RPNStackItem{Value: result})
	}
	result, _ := itemStack.Pop()
	return result.(RPNStackItem).Value
}

// 查找是否为运算符
func IsSymbol(item string) (result bool) {
	var Symbol []string = []string{"+", "-", "*", "/", "(", ")"}
	result = false
	for _, v := range Symbol {
		if v == item {
			result = true
			break
		}
	}
	return
}

// 依据运算符与值进行运算
func Caculate(value1 float64, value2 float64, symbol string) (result float64) {
	switch symbol {
	case "+":
		result = value2 + value1
	case "-":
		result = value2 - value1
	case "*":
		result = value2 * value1
	case "/":
		result = value2 / value1
	}
	return
}
