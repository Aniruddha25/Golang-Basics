package main  
import "fmt"  
import "strings"
import "regexp"
func main() {  
   fmt.Println("Ascii value of B is ","B"[0])  
   var i string;
   i = "amk";
   //length of string
   fmt.Println(len(i) );
   //Converting string  to uppercase
   fmt.Println(strings.ToUpper(i));
   //Converting string  to lowercase
   fmt.Println(strings.ToLower(i));
   //Repeating the same string again
   fmt.Println(strings.Repeat(i,4));
   //strings contains substring / not
   fmt.Println(strings.Contains(i,"mk"));
   //for printing index of substring
   fmt.Println(strings.Index(i,"m"));
   //for counting frequency of substring
   fmt.Println(strings.Count(i,"m"));
   //for replacing substring
   fmt.Println(strings.Replace(i,"a","A",1));
   //splitting the string
   var s string;
   s = i[0:2];
   fmt.Println(s);
   //for comparing two strings
   fmt.Println(strings.Compare(i,"mk"));//similar to strcmp() in c++
   //for removing blank spaces
   fmt.Println(strings.TrimSpace("    hello   "));
   //To find string using regular expressions
   re := regexp.MustCompile(".com")  
   fmt.Println(re.FindString("google.com"))
   fmt.Println(re.FindStringIndex("google.com"))
   //we can also use FindStringSubmatch() method which returns a slice of strings 
   //having the text of the leftmost match and the matches. 
   //If no match is found, the return value is an empty string. 
   rex := regexp.MustCompile("f([a-z]+)ing")  
   fmt.Println(rex.FindStringSubmatch("flying")) 














} 