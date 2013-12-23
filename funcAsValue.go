/*
 * function as values
 *
 */

 package main

 import (
    "fmt"
    "math"
 )

 func main() {
    hypo:=func(x,y float64) float64{
        return math.Sqrt(x*x + y*y)
    }
    fmt.Println(hypo(3,4)) //5
 }
