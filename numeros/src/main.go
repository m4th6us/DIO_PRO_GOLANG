package main

import "fmt"

// create the function for the show the number with multiply three
func multiply_tree(){

	array_multiply_tree := make([]int, 0)

	for i:= 1; i <= 100; i ++ {

		if (i % 3 == 0) {
			array_multiply_tree = append(array_multiply_tree, i)
		}

	}

	fmt.Println("Os multiplos de 3 são:", array_multiply_tree)
}

// create the function to show numbers that are multiples of three and say pin or five and say pan
func pin_pan() {

	pin_pan_count := make(map[int]string)
	sum_pin := 0
	sum_pan := 0

	for i := 1; i <= 100; i ++ {
		
		if (i % 3 == 0) {

			pin_pan_count[i] = "Pin"
			sum_pin ++

		} else if (i % 5 == 0) {

			pin_pan_count[i] = "Pan"
			sum_pan ++

		} 

	}

	fmt.Println("Os números que receberam Pin ou Pan são esses:", pin_pan_count)
	fmt.Println("A quantidade de 'pins' é:", sum_pin)
	fmt.Println("A quantidade de 'pans' é:", sum_pan)

}

func main() {
	multiply_tree()
	pin_pan()
}