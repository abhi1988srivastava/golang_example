/*
 *
 * user defined type and conversion
 *
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
