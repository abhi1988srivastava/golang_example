/*
 *
 * user defined type and conversion
 *
 * type A struct {
    ax, ay int
   }
   type B struct {
    A
    bx, by float64
   }

   b := B{A{1, 2}, 3.0, 4.0}
   fmt.Println(b.ax, b.ay, b.bx, b.by)
   Prints 1 2 3 4


    type A struct { a int }
    type B struct { a, b int }
    type C struct { A; B }
    var c C
    Using c.a is an error: is it c.A.a or c.B.a?
    type D struct { B; b float64 }
    var d D
    Using d.b is OK: it's the float64, not d.B.b

 */

package main

import (
    "fmt"
)

 type foo struct { //anonymous struct field
    int
 }

 type bar foo //bar is an alias for foo

func main() {

 var b bar=bar{1}

 //var f foo=b //fails here : cannot use b (type bar) as type foo in assignment
 var f1 foo=foo(b) //correct way

 //fmt.Println("%v",f)
 fmt.Printf("%v",f1)
}
