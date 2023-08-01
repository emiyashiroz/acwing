package acwing

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	//a := []int{782, 23, 2, 4, 123, 5, 12, 12, 41, 4}
	//fmt.Println(a)
	//QuickSort(a, 0, 9)
	//fmt.Println(a)

	a := []int{782, 23, 2, 4, 123, 5, 12, 12, 41, 4}
	fmt.Println(a)
	MergeSort(a, 0, 9)
	fmt.Println(a)
}
