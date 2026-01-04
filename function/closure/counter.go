package closure

func makeCounter(start int) func() int {
	s := start

	return func() int {
		s++
		return s
	}
}
