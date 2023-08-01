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

// MergeSort 归并排序
func MergeSort(q []int, l, r int) {
	if l >= r {
		return
	}
	j := (l + r) >> 1
	MergeSort(q, l, j)
	MergeSort(q, j+1, r)
	t := make([]int, r-l+1)
	m, n, k := l, j+1, 0
	for ; m <= j && n <= r; k++ {
		if q[m] <= q[n] {
			t[k] = q[m]
			m++
		} else {
			t[k] = q[n]
			n++
		}
	}
	for ; m <= j; k++ {
		t[k] = q[m]
		m++
	}
	for ; n <= r; k++ {
		t[k] = q[n]
		n++
	}
	for i := l; i <= r; i++ {
		q[i] = t[i-l]
	}
}
