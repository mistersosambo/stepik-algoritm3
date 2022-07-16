package main

import "fmt"

func main() {
	array := [5]int{}
	var a int
	var max int = -100
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	for _, z := range array {
		if z > max {
			max = z
		} else {
			continue
		}
	}
	fmt.Println(max)
}
