package main
import "fmt"
type date struct{
	day int;
	month int;
	year int;
};
type student struct{
	name string;
	rn int;
	marks int;
	birthdate date;//nested structure
}
func studentdetails(s student)  {
	fmt.Println("Name        : ",s.name);
	fmt.Println("RN          :",s.rn);
	fmt.Println("Marks       :",s.rn);
	fmt.Println("Birthdate   :",s.birthdate);

	
}
	



func main()  {
	
	a := student{name:"Aniruddha",rn:1,marks:35,birthdate:date{22,5,1998}};
	fmt.Println(a);
	fmt.Println(a.birthdate);
	studentdetails(a);
	
	


	
}