package main

import (
	"fmt"
	"runtime"
	"sync"
)

var numCpu int

func init() {
	numCpu = runtime.NumCPU()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	if len(src) <= 1 {
		return
	}
	p, q := 0, len(src)-1
	dataPerCpu := len(src) / numCpu
	pairs := make([][2]int, 0, numCpu)
	if dataPerCpu < 1 {
		dataPerCpu = 1
	}

	for i := 0; i < numCpu; i++ {
		if p > q {
			break
		}
		end := min(p+dataPerCpu-1, q)
		pair := [2]int{p, end}
		pairs = append(pairs, pair)
		p = end + 1
	}
	if pairs[len(pairs)-1][1] < q {
		pairs[len(pairs)-1][1] = q
	}

	var wg sync.WaitGroup
	for _, pair := range pairs {
		wg.Add(1)
		go DivideMerge(src, pair[0], pair[1], &wg)
	}
	wg.Wait()

	for len(pairs) > 1 {
		x := len(pairs) - 1
		var i int
		for 2 *i + 1 <= x {
			wg.Add(1)
			go Merge(src, pairs[2*i][0], pairs[2*i][1], pairs[2*i+1][1], &wg)
			pairs[i] = [2]int{pairs[2*i][0], pairs[2*i+1][1]}
			i++
		}
		wg.Wait()
		if 2 * i == x {
			pairs[i] = pairs[x]
			pairs = pairs[:i+1]
		} else {
			pairs = pairs[:i]
		}
	}
}

func Merge(src []int64, p, mid, q int, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
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

func DivideMerge(src []int64, p, q int, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}

	if p >= q {
		return
	}

	mid := (p + q) / 2
	DivideMerge(src, p, mid, nil)
	DivideMerge(src, mid+1, q, nil)
	Merge(src, p, mid, q, nil)
}

func main() {
	src := []int64{5, 1, 4, 3, 2}
	MergeSort(src)
	fmt.Println(src)
}
