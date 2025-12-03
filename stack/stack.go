package stack

import "fmt"

/*
вершина стека это последний добалвенный элемент
*/

const (
	CmdAdd   = iota // сложить два числа в стеке (top-1) = (top-1) + top
	CmdSub          // вычесть (top-1) = (top-1) - top
	CmdMul          // умножить (top-1) = (top-1) * top
	CmdDiv          // разделить (top-1) = (top-1) / top
	CmdPush         // поместить следующее число в стек
	CmdPop          // убрать число из стека
	CmdPrint        // печать верхнего элемента стека
	CmdSave         // сохранить число из стека в ячейку
	CmdLoad         // переместить число из ячейки в стек
)

func StackCalculator() {
	program := []int{CmdPush, 33, CmdPush, 44, CmdAdd, CmdPush, 567, CmdSub, CmdPush,
		-13, CmdMul, CmdPush, 5, CmdDiv, CmdPush, 45, CmdPush, 21, CmdAdd, CmdMul,
		CmdPrint, CmdSave, 'А', CmdPop, CmdPush, 3, CmdPush, 9, CmdPush, 7,
		CmdSub, CmdMul, CmdLoad, 'А', CmdMul, CmdPrint, CmdSave, 'Б',
		CmdLoad, 'А', CmdPush, 10230, CmdLoad, 'Б', CmdSub, CmdSub,
		CmdPush, 1000, CmdDiv, CmdPrint}

		for i:=0; i <= len(program); i++ {
			switch comand {
				case CmdPush:
					fmt.Println("fd")

				case CmdAdd:
					
				
				case CmdSub:
					
				
				case CmdMul:
					
				
				case CmdDiv:
					
				
				case CmdPop:
					
				
				case CmdPrint:
					
				
				case CmdSave:
					
				
				case CmdLoad:
				
			
		}

}











	// stack := make([]int, 0, 3)
	// stack = append(stack, 2, 3)

	// // fmt.Println(stack)
	// stack = append(stack, stack[0]+stack[1])
	// top := len(stack) - 1
	// valueTop := stack[top]
	// fmt.Println(stack)

	// stack = stack[:top]
	// fmt.Println("valueTop", valueTop)