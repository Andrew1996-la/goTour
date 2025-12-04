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

	stack := make([]int, 0, 200)
	registers := make(map[rune]int)
	for i := 0; i < len(program); i++ {
		switch program[i] {
		case CmdPush:
			i++
			stack = append(stack, program[i])
		case CmdAdd:
			if len(stack) < 2 {
				fmt.Println("Не хватает элементов в стеке для операции сложения")
				return
			}

			stack[len(stack)-2] += stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case CmdSub:
			if len(stack) < 2 {
				fmt.Println("Не хватает элементов в стеке для операции вычитания")
				return
			}

			stack[len(stack)-2] -= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case CmdMul:
			if len(stack) < 2 {
				fmt.Println("Не хватает элементов в стеке для операции умножения")
				return
			}

			stack[len(stack)-2] *= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case CmdDiv:
			if len(stack) < 2 {
				fmt.Println("Не хватает элементов в стеке для операции деления")
				return
			}

			stack[len(stack)-2] /= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case CmdPop:

		case CmdPrint:
			fmt.Println(stack[len(stack)-1])
		case CmdSave:
			i++
			registers[rune(program[i])] = stack[len(stack)-1]
			fmt.Println("registers: ", registers)
		case CmdLoad:
			i++
			stack = append(stack, registers[rune(program[i])])
		}

		fmt.Println("stack: ", stack)
	}
}
