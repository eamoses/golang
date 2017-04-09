package main
import "fmt"

func main(){
  arrA := [5]int {1,2,3,4,5}
  arrB := [5]int {3,4,5,6,7}
  arrC := [...]int {4,5,6,7,8}
  arrA=arrB
  arrA[2]=15
  fmt.Println(arrA,arrB)
  fmt.Printf("%p \n%p\n",&arrA,&arrB)
  fmt.Println(arrC)
}
