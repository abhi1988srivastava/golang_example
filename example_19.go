/*
 * map and structure demo
 *
 */

package main

 import (
    "fmt"
 )
 type vertex struct {
    x,y int
 }

 var m map[string]vertex

 func main() {
    m=make(map[string]vertex)
    m["A"]=vertex{1,2}

    //or
    var m1=map[string]vertex{
        "A":vertex{2,3},
        "B":vertex{4,5},
    }

    //or
    var m2=map[string]vertex{
        "A":{6,7},
        "B":{8,9},
    }

    fmt.Println(m)
    fmt.Println(m1)
    fmt.Println(m2)
 }

/*output
 *
 *
map[A:{1 2}]
map[A:{2 3} B:{4 5}]
map[A:{6 7} B:{8 9}]
 */
