package main
// import "fmt"

struct automobile interface{
  canStart()
}

struct toyota int{
  reliabilityFactor
}

func (t toyota) canStart() bool{
  return false
}

func main(){
  var myCar toyota
  myCar.cost = 10000.0
  bool starts = myCar.canStart()
  var genericCar automobile
  genericCar = myCar
}

// func (a automobile) canStart() bool{
//   return true
// }
// struct toyota type{
//   automobile
//   reliabilityFactor int
// }
// func main(){
//   var myCar toyota
//   myCar.cost = 10000.0
//   bool starts = myCar.canStart()
// }
