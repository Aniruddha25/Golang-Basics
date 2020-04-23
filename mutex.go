/*opposite process of race condition
Mutual Exclusion locks, or mutexes can be used to synchronize access to state and safely access data across many goroutines.
 It acts as a guard to the entrance of the critical section of code so that only one thread can enter the critical section at a time.
*/
package main  
import (  
   "sync"  
   "time"  
   "math/rand"  
   "fmt"  
)  
var wait sync.WaitGroup  
var count int  
var mutex sync.Mutex  
func  increment(s string)  {  
   for i :=0;i<10;i++ {  
      mutex.Lock()  
      x := count  
      x++;  
      time.Sleep(time.Duration(rand.Intn(10))*time.Millisecond)  
      count = x;  
      fmt.Println(s, i,"Count: ",count)  
      mutex.Unlock()  
        
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