package task

//Реализовать коробку с методами Set и Get.

type Box[T any] struct {
	value T
}

func (b *Box[T]) Set(t T) {
	b.value = t
}

func (b Box[T]) Get() T {
	return b.value
}
