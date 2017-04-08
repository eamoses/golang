package main
 import "fmt"
 type vertex struct{
   x, y float64
 }

 type gameCharacter struct{
   vertex
   life int
 }

 func main() {
   mainChar := gameCharacter{
     vertex:vertex{x:1.0},
     life:10,
   }

   fmt.Println(mainChar)
 }
