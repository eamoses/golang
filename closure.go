package main
import "fmt"

// func MyCounter() func(){
//   theCount:=0
//   increment:=func(){
//     theCount++
//     fmt.Println("The count is ", theCount)
//   }
//   return increment
// }

type incFunc func()
type getFunc func() int

func MyCounter(initialCount int) (incFunc, getFunc){
  theCount:= initialCount
  increment := func(){
    theCount++
    fmt.Println("The count is", theCount)
  }
  get := func() int{
    fmt.Println("The count is", theCount)
    return theCount
  }
  return increment, get
}

func main(){
  incFunc, getFunc  := MyCounter(10)
  for i := 0; i < 5; i++ {
    incFunc()
  }
  fmt.Println("The final value is ", getFunc())
}
