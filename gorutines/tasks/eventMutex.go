package tasks

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Event struct {
	ID    int
	Value int
}

type Stats struct {
	Processed int
	Sum       int
	mtx       sync.Mutex
}

func producerEvent(ctx context.Context, eventCh chan<- Event, wg *sync.WaitGroup) {
	defer wg.Done()

	i := 1

	for {
		event := Event{
			ID:    i,
			Value: 3 + rand.Intn(6),
		}

		select {
		case <-ctx.Done():
			fmt.Println("Producer завершает работу")
			return
		case eventCh <- event:
			i++
		}
	}
}

func workerEvent(ctx context.Context, eventCh <-chan Event, wg *sync.WaitGroup, stats *Stats) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker завершает работу")
			return
		case val, ok := <-eventCh:
			if !ok {
				return
			}
			stats.mtx.Lock()
			stats.Processed++
			stats.Sum += val.Value
			stats.mtx.Unlock()
		}
	}
}

func main() {
	context, cancelContext := context.WithTimeout(context.Background(), 3*time.Second)
	eventsCh := make(chan Event, 100)
	wgProducer := &sync.WaitGroup{}
	wgWorker := &sync.WaitGroup{}
	stats := Stats{}
	wgProducer.Add(1)
	go producerEvent(context, eventsCh, wgProducer)

	wgWorker.Add(1)
	go workerEvent(context, eventsCh, wgWorker, &stats)

	wgProducer.Wait()

	cancelContext()
	close(eventsCh)
	
	wgWorker.Wait()

	fmt.Println("Processed:", stats.Processed)
	fmt.Println("Sum:", stats.Sum)
}
