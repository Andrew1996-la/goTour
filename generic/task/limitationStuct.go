package task

type Number interface {
	~int | ~float64 | ~int64
}
type BoxNumber[T Number] struct {
	value T
}

func (b *BoxNumber[T]) Add(v T) {
	b.value += v
}
