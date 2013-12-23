/*
 *
 * structure and addressing
 */

 package main

 import (
    "fmt"
 )

 type vertex struct {
    x,y int
 }

func main() {
 p:=vertex{1,2} //x=1,y=2
 q:=&p //address of p
 q.x=1e9 //change will reflect in p too :1e9 = 1000000000
 fmt.Println(p)
 r:=p //no change will be reflected since contains value not address
 r.x=10
 fmt.Println(p)
 s:=&vertex{3,4} //has type *vertex, construct a pointer to s vertex{x:1} y would be 0 and vertex{} here x and y 0
 fmt.Println(s)
 fmt.Println(*s)

 t:= *s
 t.x=16
 fmt.Println(s)
 fmt.Println(t)

 var v *vertex=new (vertex)
 //or
 v1:=new(vertex)

 v1.x,v1.y=11,9
 v.x,v.y=20,19

 fmt.Println(v1)
 fmt.Println(v)



}

/*output
 *
{1000000000 2}
{1000000000 2}
&{3 4}
{3 4}
&{3 4}
{16 4}
&{11 9}
&{20 19}
 *
 */
