package main
import "fmt"
func main()  {
	var arr [5]int;
	arr[0]=4;
	arr[1]=4;
	arr[2]=5;
	arr[3]=3;
	arr[4]=2;
	var a []int =arr[1:5];
	fmt.Println(a);
	var b []int =arr[0:1];
	fmt.Println(b);
	


	


	
}