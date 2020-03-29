package main
import 
("fmt"
  "math"
)
func  main()  {
  var a int;
  for a=0;a<5;a++{
	  fmt.Println(a);
  }
  a=0;
  //while loop fashion
  for a<5{
	  fmt.Println(a);
	  a++;
  }
  var sum int;
  sum=0;
  nums :=[]int{1,3,5,7,9};
  //Range construct for accesing array elements
  for _,i := range nums{//_, is to ignore the index
	  sum+=i;

  }
  var i int;
  for i=0;i<len(nums);i++{
	  fmt.Println(nums[i]);
  }

  fmt.Println("sum :",sum);
  arr := [5]float64{1,4,9,16,25};
  for i=0;i<5 ;i++{
	fmt.Println(math.Sqrt(arr[i]));
}


}




