package gorutines

import (
	"context"
	"fmt"
	"time"
)

func foo(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("foo завершается")
			return
		default:
			fmt.Println("Foo")
			time.Sleep(1 * time.Second)
		}
	}
}

func boo(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("boo завершается")
			return
		default:
			fmt.Println("Boo")
			time.Sleep(1 * time.Second)
		}
	}
}

func contextGorutinex() {
	/*
		создание контекстов.
		parentContext - создается из главного Background контекста
		childContext - создается из parentContext

		когда мы вызываем childCancel то прекращают выполнение горутины куда мы прокинули childContext

		когда мы вызываем parentCancel то прекращают выполнение горутины куда мы прокинули parentContext
		а так же те отмена родительского контекста отменяет все дочерние
	*/
	parentContext, parentCancel := context.WithCancel(context.Background())
	childContext, childCancel := context.WithCancel(parentContext)

	go foo(parentContext)
	go boo(childContext)

	time.Sleep(5 * time.Second)
	childCancel()

	time.Sleep(5 * time.Second)
	parentCancel()

	time.Sleep(6 * time.Second)
	fmt.Println("main звершается")
}
