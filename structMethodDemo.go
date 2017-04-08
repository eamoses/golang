package main

import "fmt"

type vertex struct{
 x, y float64
}

func (v *vertex) addVertex(v2 vertex){
  v.y += v2.y
  v.x += v2.x //methods are called on a struct using dot notation
}

func main(){
  vert := vertex{4.0, 2.0}
  vert2 := vertex{3.0, 3.0}
  vert.addVertex(vert2)
  fmt.Println(vert)
}

//struct example commented out below
// type gameCharacter struct{
//  vertex
//  life int
// }
//
// func main() {
//  mainChar := gameCharacter{
//    vertex:vertex{x:1.0},
//    life:10,
//  }
//
//  fmt.Println(mainChar)
// }
