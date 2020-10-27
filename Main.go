package main

import "fmt"

func sum_array(x []int) int {
	summ := 0
	for i := 0; i < len(x); i++ {
		summ = summ + x[i]
	}
	return summ
}
func mult_array(x []int)int(
	mult := 1
for i := 0; i < len(x); i++ {
	mult *= x[i]
}
return mult
)
func main() {

	x := []int{-5, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	fmt.Println("--------------------------------NEW--------------------------------")
	//summ := 0
	//mult := 1
	middle := len(x) / 2
	// for i := 0; i < len(x); i++ {
	// 	if i < middle {
	// 		summ = summ + x[i]

	// 	} else {
	// 		mult = mult * x[i]
	// 	}

	// }
	// fmt.Println(summ, mult)//
	/*
			for i := 0; i < 9; i++ { // вот так
				summ = summ + x[i]
			}
			for i := middle; i < len(x); i++ {
				mult = mult * x[i]

			}
			fmt.Println(summ, mult)
		}
	*/
	summ := sum_array(x[0:middle])
	fmt.Println(x[middle:len(x)])
	fmt.Println(summ)

}
