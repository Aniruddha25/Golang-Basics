package main  
import (  
   "fmt"  
)  
func main() {  
   x:=10;
   y:=20;
   modify(&x);
   fmt.Println(x);
   swap(&x,&y);
   fmt.Println(x);
   fmt.Println(y);
     
}  
func swap(x *int,y *int)  {
	var temp int;
	temp=*x;
	*x=*y;
	*y=temp;
	
}
func modify(x *int)  {
	*x+=5;
	
}
