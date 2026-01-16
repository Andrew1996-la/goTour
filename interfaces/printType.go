package interfaces

import "fmt"

func PrintType(x any) {
	switch x := x.(type) {
	case int:
		fmt.Printf("type:%T value:%d\n", x, x)
	case string:
		fmt.Printf("type:%T value:%s\n", x, x)
	case float64:
		fmt.Printf("type:%T value:%f\n", x, x)
	case struct{}:
		fmt.Printf("type:%T value:%#v\n", x, x)
	default:
		fmt.Println("undefined type")
	}
}
