package main  
import (  
   "sync"  
   "time"  
   "math/rand"  
   "fmt"  
)  
var wait sync.WaitGroup  
var count int  
func  increment(s string)  {  
   for i :=0;i<10;i++ {  
      x := count  
      x++;  
      time.Sleep(time.Duration(rand.Intn(4))*time.Millisecond)  
      count = x;  
      fmt.Println(s, i,"Count: ",count)  
        
   }  
   wait.Done()  
     
}  
func main(){  
   wait.Add(2)  
   go increment("foo: ")  
   go increment("bar: ")  
   wait.Wait()  
   fmt.Println("last count value " ,count)  
}  
/*As you can see in the above example,
 the count resource is accessed by 2 go routines. 
 Each routine iterates to 10 times. In such case, 
 the count variable should be 20 at last.
  But it is not so because it is simulating race condition.*/


