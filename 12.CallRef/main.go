
package main

import "fmt"


func swap(a, b *int)int{
	var o int
	o = *a
	*a = *b
	*b = o
	
return o
}

func main() {

var p int = 100
var q int = 200
fmt.Printf("p = %d and q = %d", p, q)


swap(&p, &q)
fmt.Printf("\np = %d and q = %d",p, q)
}
