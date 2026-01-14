package composition

import "fmt"

type Logger struct{}

func (l Logger) Log(v ...any) {
	for _, i := range v {
		fmt.Println(i)
	}
}

type Metric struct{}

func (m *Metric) Inc() {
	fmt.Println("Inc")
}

type Service struct {
	Logger
	Metric
}

func serviceStart() {
	service := Service{}

	service.Log("I am John")
	service.Inc()
}
