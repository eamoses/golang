package main
import "fmt"

type myStruct struct{
  intField int
}

func (ms myStruct) myMethodName(param1 int, param2 int) int{
  return ms.intField + param1 + param2
}

func main(){
  sumert := myStruct{1}
  var retVal int
  retVal = sumert.myMethodName(3,4)
  fmt.Println(retVal)
}
