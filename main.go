package main

import "fmt"

// statistics := [
//     [{userId:1, steps:1000}, {userId:2, steps: 1500}],
//     [{userId:2, steps: 1000}]
// ]

// champ := {userId: 2, steps: 2500}

type Stat struct {
	userId int
	steps  int
}

func main() {
	stats := [][]Stat{
		{{userId: 1, steps: 1000}, {userId: 2, steps: 1500}},
		{{userId: 2, steps: 1000}},
	}

	studendStat := make(map[int]int)

	for _, day := range stats {
		for _, studend := range day {
			studendStat[studend.userId] += studend.steps
		}
	}


    var champ Stat

    for id, steps := range studendStat {
        if steps > champ.steps {
            champ.userId = id
            champ.steps = steps
        }
    }

    fmt.Println(champ)
}
