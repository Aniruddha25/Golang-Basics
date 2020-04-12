package main  
import(  
   "fmt"  
   "reflect"  
)  
func main()  {  
   age := 27  
   rune:='B'//Rune in Golang
   fmt.Println(reflect.TypeOf(age)) //to find datatype of variable at runtime
   fmt.Println(reflect.TypeOf(rune)) 
} 