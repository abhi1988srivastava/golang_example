/*
 *
 * interface is defined by set of methods.
 * A value of interface type can hold any value that implements those methods
 *
 */

 package main

 import (
    "fmt"
    "math"
 )

 type Abser interface {
    Abs() float64
 }

 type MyFloat float64 //aliasing..since you cannot use normal types as receiver type

 func (f MyFloat) Abs() float64 {
    if f<0 {
        return float64(-f)
    }else{
        return float64(f)
    }

 }

 type vertex struct {
    x,y float64
 }

 func (v *vertex) Abs() float64{
    return math.Sqrt(v.x*v.x+v.y*v.y)
 }

 func main() {
    var a Abser
    f:=MyFloat(-math.Sqrt(2))
    v:=vertex{3,4}

    a=f
    a=&v //allowed since *vertex implements Abser
    //a=v //cannot use v (type vertex) as type Abser in assignment:
	    //vertex does not implement Abser (Abs method requires pointer receiver)


    fmt.Println(a.Abs())
 }
