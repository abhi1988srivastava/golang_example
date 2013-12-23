/*
 * slice we are talking about now
 *
 * slice cant grow beyond capacity ..to manage this create a new slice and use the append function
 *
 * similar to array but can grow when we add new elements
 * different from array since it is pointer to array i.e. if we assign one slice to another, both refers
 * to the same underlying array
 *
 * make keyword is used to allocate mem for built in types (maps ,slices and channels)
 * new keyword is used to allocate mem for UDT
 */


package main

import (
    "fmt"
)

func main() {
    s1:=make([]int,5) //make slice of size 5 but it can grow as shown below
    s0:=[]int{1,2,4,5,6,7}
    s1=append(s0,3)
    s2:=make([]string,5)
    //s2=append("abhi","mali","adi") //invalid 1st argument must be a slice and must be of type of string
    s21:=[]string{"i am"}
    s2=append(s21,"abhi","mali","adi")
    s2=append(s21)
    fmt.Println(s0) //[1 2 4 5 6 7]
    fmt.Println(s1) //[1 2 4 5 6 7 3]
    fmt.Println(s2) //
}
