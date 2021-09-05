package main

import "fmt"

func main() {

   var a int = 0
   Points: for a < 200 {
      if a == 5 {
         a = a + 10
         goto Points
      }
      fmt.Printf("A = %d ", a)
      a++     
   }  
}