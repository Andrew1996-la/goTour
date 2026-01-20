package temporaryAccess


import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func Counter(count int, t time.Time) string {
	formatedTime := t.Format("2.01.2006")

	return strconv.Itoa(count) + " " + formatedTime
}

func ParseCounter(data string) (int, time.Time, error) {
	parts := strings.Split(data, " ")

	count, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, time.Time{}, err
	}

	t, err := time.Parse("2.01.2006", parts[1])
	if err != nil {
		return 0, time.Time{}, err
	}

	return count, t, nil
}

// Limits возвращает количество дней и запусков.
func Limits() (int, int, error) {
	app, err := os.Executable()
	if err != nil {
		return 0, 0, nil
	}

	// получаем путь и имя текстового файла
	name := filepath.Join(filepath.Dir(app), "data.txt")

	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			out := Counter(1, time.Now())                // записываем начальные значения
			err := os.WriteFile(name, []byte(out), 0644) // 0644 - устанавливаем разрешение на чтение и запись

			return 0, 1, err
		}
		return 0, 0, nil
	}

	var data []byte

	data, err = os.ReadFile(name)
	if err != nil {
		return 0, 0, nil
	}

	counter, t, err := ParseCounter(string(data))
	if err != nil {
		return 0, 0, nil
	}

	counter++

	if err = os.WriteFile(name, []byte(Counter(counter, t)), 0644); err != nil {
		return 0, 0, err
	}

	duration := time.Now().Sub(t)

	return int(duration.Hours()) / 24, counter, nil
}

func temporaryAccess() {
	days, counter, err := Limits()

	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Printf("Количество дней: %d\nКоличество запусков: %d\n", days, counter)

	if days > 14 || counter > 50 {
		fmt.Println("Запросите новую версию")
		return
	}

	fmt.Println("Программа готова к работе")
}
