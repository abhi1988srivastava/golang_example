/*
 *
 * methods with pointer receiver vs just named receiver
 *
 * use :
 * avoid copying value on each method call (efficient if the value is a large struct)
 * methods can modify the value that its receiver points to it
 *
 */

package main

import (
    "fmt"
    "math"
)
type vertex struct {
    x,y float64
}

func (v *vertex) Scale (f float64) {
    v.x=v.x*f
    v.y=v.y*f
}

func (v vertex) Scale1 (f float64) {
    v.x=v.x*f
    v.y=v.y*f
}

func (v *vertex) Abso () float64 {
    return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
    v1:=vertex{3,4}
    v2:=&vertex{5,6}
    v1.Scale1(6) //no change would  be there
    fmt.Println("v1 new scale after :",v1)
    fmt.Println("v1 before :",v1)
    v1.Scale(5)
    fmt.Println("v1 after :",v1)
    fmt.Println("v2 before :",v2)
    v2.Scale(5)
    fmt.Println("v2 after :",v2)
}

/*output
 *
 *
v1 new scale after : {3 4}
v1 before : {3 4}
v1 after : {15 20}
v2 before : &{5 6}
v2 after : &{25 30}*
*/
