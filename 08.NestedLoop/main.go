package main
import "fmt"

func main() {  
  var i, j int
  for i=1; i <= 10; i++   {
   for j=1; j <= 10; j++ {
      fmt.Println(i+j)
    }
  }
}
