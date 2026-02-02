package house

import "fmt"

type Thing struct {
	Name   string
	Weight int
}

type Room struct {
	Name   string
	Things []Thing
}

type House struct {
	Name  string
	Rooms [][]Room
}

func house() {
	house := House{
		Name: "Дом v1",
		// инициализируйте house нужными значениями
		Rooms: [][]Room{
			// Подвал
			{
				// 1 комната
				{Name: "Кладовка", Things: []Thing{
					{Name: "топор", Weight: 3000},
					{Name: "фонарик"},
					{Name: "брелок"},
				}},
				// 2 комната
				{Name: "Котельная", Things: []Thing{
					{Name: "верёвка", Weight: 200},
					{Name: "рюкзак", Weight: 500},
				}},
			},
			// 1-й этаж
			{
				// 1 комната
				{Name: "Столовая", Things: []Thing{
					{Name: "ручка"},
					{Name: "кольцо"},
					{Name: "карта"},
					{Name: "бинт"},
				}},
			},
		},
	}
	fmt.Println(house)
}
