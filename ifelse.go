package main

import "fmt"

func main(){
	var a int;
	fmt.Println("Enter the input :");
	fmt.Scanln(&a);//user-defined input
	if(a%2==0){
		fmt.Println(a," is divisible by 2");
		if(a%3==0){
			fmt.Println("It is also divisible by 6");//nesting of if
		}


	}else
	{
		fmt.Println(a," is not divisible by 2");

	}
	var marks int;
	marks=75;
	//ifelse statements
	if(marks<40){
		fmt.Println("Fail");
	} else if(marks>=40  && marks<=50){
		fmt.Print("D grade");
	} else if(marks>50  && marks<=60){
		fmt.Print("C grade");
	} else if(marks>60  && marks<=70){
		fmt.Print("B grade");
	} else {
		fmt.Print("A grade");
	}
	fmt.Println(" ");
	fmt.Println(" Switch case Result :");
	//Switch Case
	switch a {
	case 1:
		fmt.Println(a/2);
	case 2:
		fmt.Println(a/3);
	case 3:
		fmt.Println(a/4);
	default:
		fmt.Println(a);
		
	}


	
}