package mergesort

import "math"

func MergeSort(list []int, p, r int) {
	if p < r {
		q := (p + r) / 2
		MergeSort(list, p, q)
		MergeSort(list, q+1, r)
		merge(list, p, q, r)
	}

}

func merge(list []int, p, q, r int) {
	n1 := q - p + 1
	n2 := r - q
	left := make([]int, n1+1)
	right := make([]int, n2+1)
	for i := 0; i < n1; i++ {
		left[i] = list[p+i]
	}
	for j := 0; j < n2; j++ {
		right[j] = list[q+1+j]
	}
	left[n1] = math.MaxInt32
	right[n2] = math.MaxInt32
	i := 0
	j := 0
	for k := p; k <= r; k++ {
		if left[i] <= right[j] {
			list[k] = left[i]

			i++
		} else {
			list[k] = right[j]
			j++
		}
	}
}
