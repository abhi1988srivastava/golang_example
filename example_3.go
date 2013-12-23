package main

import (
    "fmt"
)


func main() {
    var s string="hello"
    fmt.Println(s)
    for index,content := range s{
        fmt.Println(string(s[index]),string(content)) //line 1
        fmt.Println(s[index],content) //line 2
    }

}

/*output line 2:

ascii value will be printed

104 104
111 111
..
..
..


output line 1:

as expected each char in different line
*/
