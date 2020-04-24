package main

import (
	"fmt"
	"sort"
)
type sortbylength []string;
func (s sortbylength)Len()int{
return len(s);

}
func (s sortbylength)Swap(a,b int){
  s[a],s[b]=s[b],s[a];
}
func (s sortbylength)Less(a,b int)bool{
return len(s[a])<len(s[b]);
}


func main() {
	a:=[]string{"Aniruddha","Rajat","Nikita"};
	sort.Sort(sortbylength(a));
	fmt.Println(a)
}