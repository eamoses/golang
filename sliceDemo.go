package main
import "fmt"

func main(){
  // var mySlice []int //a slice of ints, length is variable. below: an array of ints that is 5 elements long
  // var myArray [5]int
  // var aSlice []int //nil slice, length and capacity are zero
  // var bSlice = [int]{1,2,3}
  // cSlice := [int]{3,4,5}
  // arrayName[startIndexInclusive:endIndexExclusive]
  // myArray := [6]int{1,2,3,4,5,6}
  // sliceA := myArray[2:5] // calue will be {3,4,5}
  // sliceB := myArray[:3]  // value will be {0,1,2}
  // silceC := myArray[3:]  // claue will be {2,3,4,5,6}
  // sliceD := myArray[:]   // value will be {1,2,3,4,5,6}
  //3 parts: A pointer to an element in the array, a length and a capacity
  arrA := [6]int{1,2,3,4,5,6}
  slcA := arrA[0:3]
  fmt.Println( "\n",arrA,"\n",slcA,"\n")  //print out the value of both our array and  slice
  for i := 0; i < 10; i++ {
    slcA = append (slcA,90+i) // append an into to the end of our []int slice
    //print out the address of the first element of the slice, length and capacity
    fmt.Printf("P:%p L:%2v C:%2v  S:%v\n",   &slcA[0],len(slcA), cap(slcA),  slcA,)
  }
  fmt.Println( "\n",arrA,"\n",slcA)  //print out the final value of our array and slice
}
