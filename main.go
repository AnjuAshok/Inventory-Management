package main

import (
	"fmt"
)

type product struct {
	brand         string
	cost_price    float64
	selling_price float64
}

func main() {
	x := make(map[string]product)
	var v, a string
	var m, n float64
	var y int
	fmt.Println("Enter the number of products")
	fmt.Scanln(&y)
	for i := 0; i < y; i++ {
		fmt.Println("Enter the product")
		fmt.Scanln(&v)
		fmt.Println("Enter the brand name of Product")
		fmt.Scanln(&a)
		fmt.Println("Enter the cost Price")
		fmt.Scanln(&m)
		fmt.Println("enter the selling price")
		fmt.Scanln(&n)
		x[v] = product{a, m, n}

	}
	fmt.Println(x)
}
