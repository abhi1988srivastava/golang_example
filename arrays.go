/*
 *arrays cannot grow, we have slices if we want so called dynamic array
 * var arr [10]int : final size array , cant grow
 * assigning one array to another copies all the values
 * pass array to a function will receive copy of a array not pointer
 */
package main

import (
    "fmt"
)

func main() {

    var a=[3]int{1,3,4}
    //or
    b:=[...]int{1,3,4} //automatically calculates length of array

    //2D array
    c1:=[2][2]int{[2]int{1,2},[2]int{3,4}}
    //or
    c2:=[2][2]int{[...]int{1,2},[...]int{3,4}}
    //or
    c3:=[2][2]int{{1,2},{3,4}}

    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c1)
    fmt.Println(c2)
    fmt.Println(c3)

}

/*output
 *
 *
[1 3 4]
[1 3 4]
[[1 2] [3 4]]
[[1 2] [3 4]]
[[1 2] [3 4]]
  *
  */
