/*
 *
 * pointers demo
 *
 * *p++ is interpreted as(*p)++ that is first referencing and then increment the value.
 * new(T) allocates zeroed storage for a new item of type T and returns *T
 *
 * make(T,args) creates slices,maps,channels only & it returns an initialised (not zero) value
 * of type T ,not *T
 * e.g.
 * make([]int ,10,100) allocates array of 100 ints and then create a slice with length of 10 and
 * capacity of 100 and pointing to first 10 elements of the arrays
 *
 * whereas
 *
 * new([]int) returns pointer to a newly allocated zeroed slice structure i.e. pointer to a nil
 * sliced value
 */

 package main

 import (
    "fmt"
 )

 func main() {

    var p *int
    fmt.Printf("Type : %T => value : %v\n",p,p)

    var i int
    x:=&i //p point to i
    fmt.Printf("Type : %T => value : %v\n",x,x)

    p=&i
    *p=8

    fmt.Printf("Type : %T => value : %v\n",p,p)
    fmt.Printf("Type : %T => value : %v\n",*p,*p)

    var np *[]int=new ([]int)
    var mp []int=make([]int,100) //array of 100 ints

    fmt.Printf("Type : %T => value : %v\n",np,np)
    fmt.Printf("Type : %T => value : %v\n",mp,mp)

 }

 /*output
  *
  *
Type : *int => value : <nil>
Type : *int => value : 0xc080000048

  */
