package interfaces

import "fmt"

//Создай интерфейс Logger с методом Log(message string).
//Создай интерфейс AdvancedLogger, который встраивает Logger и добавляет метод Error(message string).
//Создай структуру ConsoleLogger, которая реализует AdvancedLogger.
//Протестируй вызовы методов через переменную типа AdvancedLogger.

type Logger interface {
	Log(message string)
}

type AdvancedLogger interface {
	Logger
	Error(message string)
}

type ConsoleLogger struct{}

func (c *ConsoleLogger) Log(message string) {
	fmt.Println(message)
}

func (c *ConsoleLogger) Error(message string) {
	fmt.Println(message)
}

func embedding() {
	var a AdvancedLogger = &ConsoleLogger{}
	a.Log("Hello World")
	a.Error("error test")
}
