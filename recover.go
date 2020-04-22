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


}
