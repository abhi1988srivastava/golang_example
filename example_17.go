/*
 * for loops versions
 *
 */

 package main

 import (
    "fmt"
 )

func main() {

    sum,sum1:=0,0
    for i:=0;i<10;i++ {
        sum+=i
    }

    fmt.Println("sum is ",sum)

    sum1=1

    for ;sum1<10; { //even without ";" it will work
        sum1++
    }
    fmt.Println("sum1 is ",sum1)

    //infinite loop

    /*for{

    }
*/
}
