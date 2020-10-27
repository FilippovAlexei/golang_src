package main

import "fmt"

func main() {
	fmt.Println("------------------DZ3-------------")
	y := []int{-3, -2, -1, 0, 1, 2, 3}
	for i := 0; i < len(y); i++ {
		if y[i] <= 0 {
			fmt.Println(y[i])

		} else {
			fmt.Println("-----")
		}
	}

	fmt.Println("--------------------------------DZ4--------------------------------")
	x := []int{6, -5, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 34, -1}
	for i := 0; i < len(x); i++ {
		if x[i] <= 0 {
			fmt.Println("Отрицательные", x[i]) // не правильный перделать алгоритм

		} else {
			fmt.Println(" положительные ", x[i])
		}

	}
}
