package main
import "fmt"

func main(){
  arrA := [5]int{1,2,3,4,5}
  slcA := arrA[:]
  fmt.Println(slcA, len(slcA), cap(slcA))
  fmt.Printf("\n%p\n%p\n", &arrA, &slcA[0])
  for i := 0; i < 10; i++ {
    slcA = append(slcA,99)
    fmt.Println(slcA, len(slcA), cap(slcA))
  }
  fmt.Println(slcA, len(slcA), cap(slcA))
  fmt.Printf("\n%p\n%p\n", &arrA, &slcA[0])

  arrA[2]=-100
  fmt.Println(arrA, "\n",slcA, )
}
