package closure

import "fmt"

func once(fn func()) func() {
	isPrint := false

	return func() {
		if !isPrint {
			fn()
			isPrint = true
		}
	}
}

func print() {
	hello := once(func() {
		fmt.Println("hello")
	})

	hello()
	hello()

}
