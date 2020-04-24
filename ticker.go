package main  
import "time"  
import "fmt"  
import "log"
func haveFun(s string) {  
    log.Printf("\tA: Let's have fun: %v", s)  
}  
func doPolling() {  
    for _ = range time.Tick(10 * time.Second) {  
        haveFun("\t B : Okay!")  
    }  
} 
func main() {  
    tickerValue := time.NewTicker(time.Millisecond * 100)  
    go func() {  
        for t := range tickerValue.C {  
            fmt.Println("Tick at", t)  
        }  
    }()  
    time.Sleep(time.Millisecond * 500)  
    tickerValue.Stop()  
	fmt.Println("Ticker stopped")  
	go doPolling()  
    select {} //The select statement lets a goroutine wait on multiple communication operations.
} 