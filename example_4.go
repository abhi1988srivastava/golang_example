package main

import (
    "fmt"
)


func main() {

    //string is immutable

    var s string="hello"
    //var multilineString string="starting" //not allowed like this :error or "starting";
    //+"ending"                             +"ending";
    var multilineString string="starting"+
    "ending"
    var multilineString1 string=`"starting"
    "ending"
    `
    fmt.Println(multilineString)
    fmt.Println(multilineString1)
    //s[0]='c' //not allowed
    c:=[]rune(s) // ascii value of each character
    fmt.Println(c)
    c[0]='c' //changed allowed in slice and rune
    s1:=string(c)
    fmt.Println("original string-->",s)
    fmt.Println("modified string-->",s1)
}


/*output
 *
 *
startingending
"starting"
    "ending"

[104 101 108 108 111]
original string--> hello
modified string--> cello
 */
