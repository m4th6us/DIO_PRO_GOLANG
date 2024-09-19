package main

import (
	"fmt"
	"log"
)

func main() {

	var x, y float64;

	fmt.Println("Digite valores para x e y")
	fmt.Println("x: ")
	fmt.Scan(&x)

	fmt.Println("y: ")
	fmt.Scan(&y)

	log.Println("sum: ", sum(x, y))
	log.Println("subtract: ", subtract(x, y))
	log.Println("multiply: ", multiply(x, y))
	log.Println("division: ", division(x, y))
	
}

func sum(x, y float64) float64 {
	return x + y
}

func subtract(x, y float64) float64{
	return x - y
}

func multiply(x, y float64) float64{
	return x * y
}

func division(x, y float64) float64{
	return x / y
}

