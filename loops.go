package main
import "fmt"
func main(){
  mySlice := []string{"a","b","y","z"}
  for index, element := range mySlice { //_ if you dont care about printing index
    fmt.Println(index, element)  //will print "a", etc...
  }
  iq := map[string]int{"UT": 115, "OU": 90, "A&M": 75}
    for index := range iq {
      fmt.Println(index, iq)  //will print "UT", etc...
  }
  for _, element := range iq {
    fmt.Println(element)  //will print "115", etc...
  }
}
