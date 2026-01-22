package list

func Min(list []int) int {
	if len(list) == 0 {
		return 0
	}
	
	min := list[0]

	for _, v := range list {
		if v < min {
			min = v
		}
	}
	return min
}
