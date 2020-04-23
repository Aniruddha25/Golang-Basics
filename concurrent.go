//Large programs are divided into smaller sub-programs.
// Programs which run their smaller components at the same time is known as concurrency.
//The parts of an application that run concurrently are called goroutines. 
//Goroutines and channels are used for structuring concurrent programs.
/*A race condition occurs in Go when two or more goroutines try to access the same resource. 
It may happen when a variable attempts to read and write the resource without any regard to other routines.*/
package main  
import (  
   "fmt"  
   "time"  
   "sync"  
)  
var wg = sync.WaitGroup{}  
  
func main() {  
   wg.Add(3)  
   go fun1()  
   go fun2() 
   go func3() 
   wg.Wait()  
}  
func fun1(){  
   for  i:=0;i<10;i++{  
      fmt.Println("fun1,  ->",i)  
      time.Sleep(time.Duration(5*time.Millisecond))  
   }  
   wg.Done()  
}  
func fun2(){  
   for i:=0;i<10;i++{  
      fmt.Println("fun2,  ->",i)  
      time.Sleep(time.Duration(10*time.Millisecond))  
   }  
   wg.Done()  
} 
func func3() {
	fmt.Println("Hello");
	time.Sleep(time.Duration(3*time.Millisecond))
	wg.Done();
}