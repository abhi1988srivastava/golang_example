/*
 * function demonstration
 *
 */

 package main

 import (
    "fmt"
 )

func g() {
    //fmt.Println(a) // undefined a
    //fmt.Println(a)
    fmt.Println("Inside g()")
}
func f() {
    a:=5
    fmt.Println(a)
    g()

}
 func main() {
    f()
 }

 /*output
  *
  * 5
Inside g()
*/
