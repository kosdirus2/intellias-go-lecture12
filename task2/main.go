package main

import "fmt"

// Конкурентно порахувати суму усіх слайсів int, та роздрукувати результат.
// Приклад:
// [ [ 4, 6 ], [ 7, 9 ] ]
// Результат друку:
// “result: 26”
func main() {
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}
	c1 := make(chan int)
	var result int

	for _, v := range n {
		go calculate(v, c1)
	}

	for i := 0; i < len(n); i++ {
		result += <-c1
	}

	fmt.Println("result:", result)
}

func calculate(intSlice []int, c1 chan int) {
	var sum int
	for _, i := range intSlice {
		sum += i
	}

	c1 <- sum
}
