package main  
import (  
   "os"  
   "log"  
  "io/ioutil"  
   "fmt"  
)  
func main() {  
   file,err:=os.Create("file.csv");
   if err!=nil{
	   log.Fatal(err);
   }
   file.WriteString("1,2,3,4,5");
   
   r,error:=ioutil.ReadFile("file.csv");
   if error!=nil{
	   log.Fatal(error)
   }  
   rf:=string(r);
  
   fmt.Println(rf);
}  