package insertionsort

func Sort(list []int) []int {
	for j := 1; j < len(list); j++ {
		key := list[j]
		i := j - 1
		for ; i >= 0 && list[i] > key; i-- {
			list[i+1] = list[i]
		}
		list[i+1] = key
	}
	return list
}
