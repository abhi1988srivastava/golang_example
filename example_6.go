package main

import (
    "fmt"
)

func main() {
    var a1 string="abhinav"
    var a=[]rune(a1)
    for i,j:=0,len(a)-1;i<j; i,j=i+1,j-1{
        a[i],a[j]=a[j],a[i]
    }
    fmt.Println("original string",a1)
    fmt.Println("new string ",string(a))
}


/*output
 *
 *
 *
original string abhinav
new string  vanihba
 * /
