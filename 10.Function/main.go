package main

import "fmt"

func main() {
   var a int = 10
   var b int = 20
   var c int
   c = adit(a,b)
   fmt.Printf( " Sum : %d\n", c )
}

func adit(num1, num2 int) int {
  return num1+num2
 
}