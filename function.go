package main
import "fmt"

func main ()  {
	var i int;
	i = num();
	fmt.Println("i :",i);
	var j int;
	j=10;
	var value int;
	value=3;
	switch value {
	case 1:
		square(j);
	case 2:
		add(i,j);
	case 3:
		fmt.Println("Sub Product Div");
		fmt.Println(smd(i,j));

		
	}

	i=modifynum(i);
	fmt.Println("i : ",i);

	
	


	
	
}
func num() int/*return type*/ {
	var i int;
	i=10;
	return i;
	
}
func square( i int) {
	fmt.Println(i*i);

}
func add(i int ,j int) {
	fmt.Println(i+j);
}
//Returning Multiple values
func smd(i int , j int) (int,int,int){
	return i-j,i*j,i/j;
}
func modifynum(j int) int {
	j+=5;
	return j;

}