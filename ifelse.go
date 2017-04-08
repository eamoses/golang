package main
import (
  "fmt"
  "math/rand"
)

func main() {
  random := rand.Intn(15)
  fmt.Println(random)
  if random < 5 {
    fmt.Println("Our random number generated is less than five")
  } else if random < 3 {
    fmt.Println("Our random number generated is less than three")
  } else {
    fmt.Println("Our random number generated is greater than or equal to five.")
  }
  if randomNumber := rand.Intn(10); randomNumber < 3 {
    fmt.Println("our number is < 3")
  } else if randomNumber < 7 {
    fmt.Println("our number is <7")
  } else {
    fmt.Println("our number is >= 7")
  }
}
