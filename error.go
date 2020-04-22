package main
import
(
	"fmt"
	"math"
	"errors"
)
func squareroot(b float64) (float64, error) {
	
	if b<0{
		return 0,errors.New("No negative value");
	}
	return math.Sqrt(b),nil;
	
}
func division(a int ,b int) (int,error) {
	if b==0 {
		return 0,errors.New("Division by zero error");
	}
	return (a/b),nil;
	
}
func checknumber(a int)(int,error)  {
	if a<0{
		return 0,errors.New("Negative number");
	}
	return a,nil;
	
}
func main()  {
	j,error:=squareroot(-25);
	if error!=nil{
		fmt.Println(error);
	} else {
       fmt.Println(j);
	}
	k,e:=division(5,0);
	if e!=nil{
		fmt.Println(e);

	}else{
		fmt.Println(k);
	}
	l,err1:=checknumber(10);
	if err1!=nil{
		fmt.Println(err1);

	}else{
		fmt.Println(l);
	}

	

	
}
