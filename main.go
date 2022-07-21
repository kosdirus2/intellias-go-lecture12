package main

import (
	"fmt"
	"sync"
)

// Конкурентно порахувати суму кожного слайсу int, та роздрукувати результат.
// Потрібно використовувати WaitGroup.
// Приклад:
// [ [ 4, 6 ], [ 7, 9 ] ]
// Результат друку:
// Порядок друку не важливий.
// “slice 1: 10”
// “slice 2: 16”
func main() {
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}

	var wg sync.WaitGroup
	wg.Add(len(n))

	for i, v := range n {
		go func(i int, sliceInt []int) {
			defer wg.Done()
			sum(i, sliceInt)
		}(i+1, v)
	}

	wg.Wait()
}

func sum(i int, sliceInt []int) {
	var sumOfElems int
	for _, v := range sliceInt {
		sumOfElems += v
	}
	fmt.Printf("slice %d: %d\n", i, sumOfElems)
}
