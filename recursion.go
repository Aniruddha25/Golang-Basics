package main
import "fmt"
func main()  {
	var i int;
	i=5;
	var f int;
	f=fact(i);
	fmt.Println("factorial : ",f);
	var n int;
	n=3;
	fmt.Println("Power :",power(i,n));


	
}
func fact(i int) int  {
	if i<=1{
		return 1;
	}

	return i*fact(i-1);
	
}
func power(i int , n int) int {
	if(n==0){
	   return 1;
	}

	return i*power(i,n-1);
	
}