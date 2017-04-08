package main
import "fmt"
func main() {
  os := "linux"
  switch os{
  case "OS X":
    fmt.Println("MAC")
  case "linux":
    fmt.Println("Bob's System")
  default:
    fmt.Printf("%s\n", os)
  }
}
