package main

import (
    "fmt"
)


func main() {

    var name string="abhinav"
    if len(name)!=0{ //if name!=nil wont work since nil comparsion with pointers and arrays
        fmt.Println("Inside if")
    }else{
        fmt.Println("Inside else")
    }

    //this is same

    if name:="abhinav"; len(name)!=0 {
        fmt.Println("Inside if")
    }else{
        fmt.Println("Inside else")
    }


}

//samne is allowed with switch statements : optional initialization
