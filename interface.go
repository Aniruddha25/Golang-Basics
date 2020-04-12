package main  
import (  
   "fmt"  
)

type product interface {  //similar to abstraction in c++
   purchase()  
}  
func details(p product)  {  
   fmt.Println(p)  
     
}  
type book struct {  
   name string  
   price int 
}  
func (b book) purchase()  {  
   fmt.Println("Purchase at ",b.price)  
     
}  
type tape struct {  
	name string  
	price int
}  
func (t tape) purchase()  {  
	fmt.Println("Purchase at ",t.price)  
	  
 }

func main() {  
 c1:=book{"HP1",100};
 t1:=tape{"t1",100};

   c1.purchase() 
   t1.purchase() 
     
   details(c1)
   details(t1)

}  

