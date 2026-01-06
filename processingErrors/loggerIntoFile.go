package processingErrors

import (
	"log"
	"os"
)

func loggerIntoFile() {
	file, err := os.OpenFile(
		"info.log",                          // наименование файла
		os.O_CREATE|os.O_APPEND|os.O_WRONLY, // если файл есть, то писать в него, если нет, то создать|дописывать дальше в файл|только для записи
		0644,                                // досупы
	)
	if err != nil {
		log.Fatal(err)
	}

	// откладываем закрытие файла
	defer file.Close()

	// устанавливаем назначение вывода в файл info.log
	log.SetOutput(file)
	log.Print("Logging to a file in Go!")
}
