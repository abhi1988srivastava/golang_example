/*
 *
 * defer demonstration and anonymous functions
 * anoymous functions in map
 * defer executes in LIFO ..
 *
 */

 package main

 import (
    "fmt"
 )

 func f() {
    for i:=0;i<=5;i++ {
         defer fmt.Println(i)
    }
 }
 func main() {
    f()
    xs:=map[int]func() int{
        1: func () int {return 10},
        2: func () int {return 20},

        }

    a:=func(){
        fmt.Println("Abhinav anonymous function")
    }
    a()
    fmt.Printf("%T => %v\n", xs[1], xs[1])
	fmt.Printf("%T => %v\n", xs[1](), xs[1]())

 }

/*output
 *
5
4
3
2
1
0
Abhinav anonymous function
func() int => 0x401370
int => 10
*/
