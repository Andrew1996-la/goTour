package generic

import "fmt"

var mysteryTypeValue any = "i am string!"

func typeConversion(anyType any) {
	switch v := anyType.(type) {
	case string:
		fmt.Printf("%#v is a string\n", v)
	case int:
		fmt.Printf("%#v is a int\n", v)
	default:
		fmt.Printf("I don't know about type %T\n", v)
	}
}
