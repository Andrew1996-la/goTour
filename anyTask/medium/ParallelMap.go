package medium

import (
	"context"
	"sync"
)

func ParallelMap(ctx context.Context, items []int, limit int, fn func(context.Context, int) (int, error)) ([]int, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	res := make([]int, len(items))
	sem := make(chan struct{}, limit)

	wg := &sync.WaitGroup{}

	var firstError error
	var mu sync.Mutex

	for i, item := range items {
		sem <- struct{}{}

		wg.Add(1)
		go func() {
			defer wg.Done()
			defer func() {
				<-sem
			}()

			val, err := fn(ctx, item)
			if err != nil {
				mu.Lock()
				if firstError == nil {
					firstError = err
					cancel()
				}
				mu.Unlock()
				return 
			}

			res[i] = val
		}()
	}

	wg.Wait()

	if firstError != nil {
		return nil, firstError
	}

	return res, nil
}
