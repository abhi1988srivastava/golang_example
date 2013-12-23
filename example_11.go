/*
 *
 * The copy built-in function copies elements from a source slice into a destination slice.
 * (As a special case, it also will copy bytes from a string to a slice of bytes.)
 * The source and destination may overlap. Copy returns the number of elements copied,
 * which will be the minimum of len(src) and len(dst).
 */

 package main

 import (
    "fmt"

 )

 func main() {
    s:="abhinav"
    b:=make([]byte,len(s))
    x:=copy(b,s)
    fmt.Println(x)
    fmt.Println(string(b))
 }
