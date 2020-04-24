package main  
  
import "fmt"  
import (  
    "math/rand"  
    "time"  
)  
func main() {  
    
    fmt.Println(rand.Intn(50));
    fmt.Print(rand.Float64())   // will produce random number between 0 to 1  
    fmt.Println()  
      
    rand.Seed(time.Now().Unix())  // seeding do that random number will produced  
    num := random(1, 20)  
  
    fmt.Println(num)  
  
}  
  
func random(min, max int) int {  
    return rand.Intn(max - min) + min  
}  