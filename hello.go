package main
import "fmt"
func main() {
  type MyBool bool
  const True = false
  const TypedTrue bool = true
  var mb MyBool
  mb = true // OK
  // mb = True // OK
  // mb = TypedTrue // BAD because I'm mixing const with var
  type MyFloat64 float64
  const Zero = 0.0
  const TypedZero float64 = 0.0
  var mf MyFloat64
  mf = 0.0
  mf = Zero
  type MyComplex128 complex128
  const I = (0.0 + 1.0i)
  const TypedI complex128 = (0.0 + 1.0i)
  var mc MyComplex128
  mc = (0.0 +1.0i)
  mc = I
  // mc = TypedI
  type MyInt int
  const Three = 3
  const TypedThree int = 3
  var mi MyInt
  mi = 3
  mi = Three
  // mi = TypedThree
  var e float32 = 1
  var i int = 1.000
  var u uint32 = 1e3 - 99.0*10.0 - 9
  var c float64 = '\x01'
  var p uintptr = '\u0001'
  var r complex64 = 'b' - 'a'
  var b byte = 1.0 + 3i - 3.0i
  var f = 'a' * 1.5
  fmt.Println(f)
  fmt.Println(e, i, u, c, p, r, b)
  fmt.Println(mi)
  fmt.Println(mc)
  fmt.Println(mf)
  fmt.Println(mb)
  fmt.Printf("%T: %s", "", "Hello\n")
  fmt.Printf("%T %v\n", 0, 0)
  fmt.Printf("%T %v\n", 0.0, 0.0)
  fmt.Printf("%T %v\n", '界', '界')
  fmt.Printf("%T %v\n", 0i, 0i)
}
