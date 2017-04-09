package main
import "fmt"

func main(){
  // var aMap map[string]int //string is the type for the key an dint is the type for the value
  // aMap = map[string]int{} // set aMap equal to an empty map
  // bMap := map[string]int{} // shorthand
  // aMap["a"] = 1 //aMap now contains the value 1 for the key "a"
  // //delete remove key/val pair
  // delete(aMap, "a") // first parameter is map variable, second is the key

  myMap := map[string]int{}//declare a map and initialize to an empty map
  //check to see if there is a key value pair for key "a"
  val, ok := myMap["a"] //val is an int, ok is bool
  fmt.Println(val, ok)
  //add some key value pairs
  myMap["a"]=3
  myMap["b"]=1
  myMap["c"]=2
  //use if statement to check to see if there is a key value pair for key "a"
  if val, ok = myMap["a"]; ok{
  fmt.Println("there is a value for the a key")
  fmt.Println(val, ok)
  }
 }
