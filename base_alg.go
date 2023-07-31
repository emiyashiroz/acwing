package acwing

// QuickSort 快排
func QuickSort(q []int, l, r int) {
	if l >= r {
		return
	}
	i, j := l-1, r+1
	x := q[l]
	for i < j {
		for i++; q[i] < x; {
			i++
		}
		for j--; q[j] > x; {
			j--
		}
		if i < j {
			q[i], q[j] = q[j], q[i]
		}
	}
	QuickSort(q, l, j)
	QuickSort(q, j+1, r)
}
