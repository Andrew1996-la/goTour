package closure

import "fmt"

func counterList() {
	var counters []func() int

	for i := 0; i < 3; i++ {
		start := i * 10
		counters = append(counters, makeCounter(start))
	}

	for _, counter := range counters {
		fmt.Println(counter())
		fmt.Println(counter())
	}
}
