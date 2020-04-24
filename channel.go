
 /*Channels are the pipes that connect concurrent goroutines.
 You can send values into channels from one goroutine and receive those values into another goroutine.*/
package main
import "fmt"
func main()  {
	
	m:=make(chan int)
	p:=make(chan string)
	go func ()  {
		m <-30;
		p <-"Hello";
		
	}()
	var j int;
	j= <- m;
	fmt.Println(j);
	k:= <-p
	fmt.Println(k);
}

