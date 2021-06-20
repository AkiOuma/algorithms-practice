package linear

import "fmt"

type PolyArrNode struct {
	Coef int
	Exp  int
}

// 公共使用数据
func ReadData() (P1 []PolyArrNode, P2 []PolyArrNode) {
	P1 = []PolyArrNode{{9, 12}, {15, 8}, {3, 2}}
	P2 = []PolyArrNode{{26, 19}, {-4, 8}, {-13, 6}, {82, 0}}
	return
}

// 使用顺序存储结构实现多项式相加
func AddPolyArr(P1 []PolyArrNode, P2 []PolyArrNode) {
	length := len(P1) + len(P2)
	result := make([]PolyArrNode, 0, length)
	ptr1 := 0
	ptr2 := 0
	p1_length := len(P1)
	p2_length := len(P2)
	for ptr1 < p1_length && ptr2 < p2_length {
		diff := P1[ptr1].Exp - P2[ptr2].Exp
		switch {
		case diff > 0:
			result = append(result, P1[ptr1])
			ptr1++
		case diff < 0:
			result = append(result, P2[ptr2])
			ptr2++
		default:
			result = append(result, PolyArrNode{P1[ptr1].Coef + P2[ptr2].Coef, P1[ptr1].Exp})
			ptr1++
			ptr2++
		}
	}
	if ptr1 < p1_length {
		for i := ptr1; i < p1_length; i++ {
			result = append(result, P1[i])
		}
	} else if ptr2 < p2_length {
		for i := ptr2; i < p2_length; i++ {
			result = append(result, P2[i])
		}
	}
	fmt.Printf("%v\n", result)
}

// 定义链表节点
type PolyLinkNode struct {
	Coef int
	Exp  int
	Next *PolyLinkNode
}

// 实现构建链表的函数
func BuildList(p []PolyArrNode) PolyLinkNode {
	var head PolyLinkNode
	current := &head
	for _, v := range p {
		node := PolyLinkNode{
			Coef: v.Coef,
			Exp:  v.Exp,
		}
		current.Next = &node
		current = &node
	}
	return head
}

// 使用链表了结构实现多项式相加
func AddPolyLink(P1 []PolyArrNode, P2 []PolyArrNode) {
	var head PolyLinkNode
	current := &head
	P1_link := BuildList(P1)
	P2_link := BuildList(P2)
	ptr1 := P1_link.Next
	ptr2 := P2_link.Next
	for ptr1 != nil && ptr2 != nil {
		diff := ptr1.Exp - ptr2.Exp
		switch {
		case diff > 0:
			current.Next = ptr1
			current = ptr1
			ptr1 = ptr1.Next
		case diff < 0:
			current.Next = ptr2
			current = ptr2
			ptr2 = ptr2.Next
		default:
			ptr1.Coef = ptr1.Coef + ptr2.Coef
			if ptr1.Coef != 0 {
				current.Next = ptr1
				current = ptr1
			}
			ptr1 = ptr1.Next
			ptr2 = ptr2.Next
		}
	}
	if ptr1 != nil {
		current.Next = ptr1
	} else if ptr2 != nil {
		current.Next = ptr2
	}
	PrintList(head)
}

// 实现打印链表的方法
func PrintList(p PolyLinkNode) {
	ptr := p.Next
	fmt.Print("[")
	for ptr != nil {
		fmt.Printf("{%d %d}", ptr.Coef, ptr.Exp)
		ptr = ptr.Next
		if ptr != nil {
			fmt.Print(" ")
		}
	}
	fmt.Println("]")
}
