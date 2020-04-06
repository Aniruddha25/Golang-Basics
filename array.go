package main
import "fmt"
func main()  {
	var arr [5]int;//1-d array declaration
	var i int;
	var j int;
	var matrix [2][2]int;
	var a =[3]int{1,4,9};//array declaration with defination
	for i=0;i<len(a);i++{
		
		fmt.Println(a[i]);

	}
	for i=0;i<len(arr);i++{
		arr[i]=i*i;
		fmt.Println(arr[i]);

	}
	for i=0;i<2;i++{
		for j=0;j<2;j++{
			matrix[i][j]=i+j;
			fmt.Print(matrix[i][j]);
		
		}
		fmt.Println("")
	}
	var mat = [6]int{9,3,10,5,4,7};
	//selection sort
	for i=0;i<5;i++{
		for j=i+1;j<6;j++{
			if(mat[i]>mat[j]){
				var temp int;
				temp=mat[i];
				mat[i]=mat[j];
				mat[j]=temp;
			}
			
		}
	}
	for i=0;i<6;i++{
		fmt.Print(mat[i]);
	}
	
	
	
}
