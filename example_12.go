/*map in golang
 *
 */

 package main

 import (
    "fmt"
 )


func main() {
    monthdays:=map[string]int{
        "Jan" :3,
        "Feb" :4,
    }

    //or

    months:=make(map[string]int)
    months["Jan"]=3
    months["Feb"]=4

    if monthdays["Jan"]==months["Jan"] {
        fmt.Println("I am Jan");
    }

    v,ok:=monthdays["Feb"]
    if !ok{
        fmt.Println("No such Key playaa")
    }else{
        fmt.Println("I am Feb :",v)
    }



}
