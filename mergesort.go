package main

import (
	"fmt"
	"runtime"
	"sync"
)

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	var wg sync.WaitGroup
	wg.Add(1)
	go DivideAndMerge(src, 0, len(src)-1, &wg)
	wg.Wait()
}

func DivideAndMerge(src []int64, p, q int, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	if p >= q {
		return
	}
	mid := (p + q) / 2

	if q - q + 1 >= 2000 {
		var waitGroup sync.WaitGroup
		waitGroup.Add(1)
		go DivideAndMerge(src, p, mid, &waitGroup)
		waitGroup.Add(1)
		go DivideAndMerge(src, mid+1, q, &waitGroup)
		waitGroup.Wait()
	} else {
		DivideAndMerge(src, p, mid, nil)
		DivideAndMerge(src, mid+1, q, nil)
	}
	temp := make([]int64, 0, q-p+1)
	i, j := p, mid+1
	for i <= mid && j <= q {
		if src[i] <= src[j] {
			temp = append(temp, src[i])
			i++
		} else {
			temp = append(temp, src[j])
			j++
		}
	}
	for i <= mid {
		temp = append(temp, src[i])
		i++
	}
	for j <= q {
		temp = append(temp, src[j])
		j++
	}
	copy(src[p:q+1], temp)
}

func main() {
	fmt.Println(runtime.NumCPU())
	src := []int64{5, 1, 4, 3, 2}
	MergeSort(src)
	fmt.Println(src)
}
