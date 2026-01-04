package closure

func counter() func() int {
	var i int

	return func() int {
		i++
		return i
	}
}
