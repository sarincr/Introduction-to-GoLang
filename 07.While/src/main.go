package main

import "fmt"

func main() {
	sum := 0
	i := 0
	for i <= 100 {
		sum = sum+i
		i=i+1
	}
	fmt.Println(sum)
}
