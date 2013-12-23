package main

import (
    "fmt"
)

func compare(a,b []byte) int {
    var val int
    for i:=0;i<len(a) && i< len(b);i++ {
        switch{
            case a[i]>b[i]:
                val=1
            case a[i]<b[i]:
                val=-1
            default: val=0
        }
    }
    return val
}

func main() {
    a:=[]byte{1,2,3,4,5}
    b:=[]byte{3,5,6,7,8}
    c:=[]byte{5,4,3,2,1}
    fmt.Println(compare(a,b)) //-1
    fmt.Println(compare(a,c)) //1

}
