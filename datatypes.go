package main
import "fmt"

func main()  {
	
	
	var i int;//Hint : variable declaration is little bit similar to sql
	var j float64;
	var s string;
	var p bool;

	//Defining multiple variables at the same time
	var(
		e =20;
		f =30;
	
	
		
	)
	
	fmt.Println(e,f);
	i=10;
	j=20.0;
	s="Aniruddha";
	p=true;
	fmt.Printf("%v %v  \n" ,i,j)
	fmt.Printf("%q \n",s);//%q specifier used for strings
	fmt.Println("Length of string :" , len(s));//length of string
	fmt.Printf("%v \n",p);
	fmt.Println("Value of P :", p);
	//Complex Datatype
	var t complex64 =complex(10,5);
	fmt.Println(t);
	const pi float64 = 3.14;
	fmt.Println(pi);
	
	
}

