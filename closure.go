//A closure is a function which refers reference variable from 
//outside its body. The function may access and assign to the
// referenced variables.
package main
import 
( "fmt"
  
)

func main()  {
	var i float64;
	i=5;
	cubenum := func( ) float64 {
		i=i*i*i;
		
		return i;
		
	}
	fmt.Println("cube :",cubenum());


	
}
