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
			fmt.Println("foo is done")
			return
		default:
			fmt.Println("foo")
			time.Sleep(time.Millisecond * 100)
		}

	}
}

func boo(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("boo is done")
			return
		default:
			fmt.Println("boo")
			time.Sleep(time.Millisecond * 100)
		}
	}
}

func gracefulShutdown() {
	parentContext, parentCloseContext := context.WithCancel(context.Background())
	childContext, childCloseContext := context.WithCancel(parentContext)

	go foo(parentContext)
	go boo(childContext)

	time.Sleep(time.Second * 1)
	childCloseContext()

	time.Sleep(time.Second * 3)
	parentCloseContext()

	time.Sleep(time.Second * 5)
	fmt.Println("program is done")
}
