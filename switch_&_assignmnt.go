/*
 * := vs =
 * := shorthand declaration which means that in the left hand we need to have atleast a new
 * variable declaration for it to be correct
 *
 * outside function every constraint begins with a keyword (var,func & so on) & := constraint
 * wont be available
 *
 * const can be declared with :=
 *
 */

package main

import (
    "fmt"
)

func main() {

    i:=0
    switch i {
    case 0: fmt.Println("Printme always")
            fallthrough
    case 1: fmt.Println("print me when fallthorugh in above")
    default: fmt.Println("I am default , print if no match only.no print if fallthrough")
    }
}
