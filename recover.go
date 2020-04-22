package main

import (
  "fmt"
  "errors"
  "math"
)
func division(a,b int)int{
  defer func(){
    fmt.Println(recover());

  }();
  q:=(a/b);
  return q;
  

  
}
func div(a,c int)(int,error){
  if c==0{
    return 0,errors.New("DividebyZeroError");
  }
  return a/c,nil;

}
/*The defer keyword is generally used for cleaning purpose. The defer keyword postpones the execution of a function or statement until the end of the calling function.

It executes code (a function or expression) when the enclosing function returns before the closing curly brace }. It is also executed if an error occurs during the execution of the enclosing function.
*/
func squareroot(b float64) float64{
 defer func(){
  fmt.Println(recover());

  }()
  return math.Sqrt(b);

}

func main() {
  j:=division(6, 3);
  fmt.Println(j);
  k,e:=div(5, 0);
  if e!=nil{
    fmt.Println(e);
  }else{
    fmt.Println(k);
  }
  p:=squareroot(64);
  fmt.Println(p);
  defer print("Hello");


}
func print(s string)
{
	fmt.Println(s);
}
