/*fibonacci,factorial
 * merge sorting,
 *
 *
 */

 package main

 import (
    "fmt"
 )

func fact(x int) int{
     if x==1 {
        return 1
    }
    return x*fact(x-1)

}

func fiba(num int) int{
    if num<=1 {
        return num
    }else{
        return fiba(num-1)+fiba(num-2)
    }
}

func mergeSort(data []int, first int , n int) {
    var n1,n2 int
     if n1>1 {
        n1=n/2
        n2=n-n1

        mergeSort(data, first, n1)      // Sort data[first] through data[first+n1-1]
        mergeSort(data, first + n1, n2) // Sort data[first+n1] to the end

        // Merge the two sorted halves.
        merge(data, first, n1, n2)
              }

}

func merge(data []int,first int,n1,n2 int) {
    var temp []int=make([]int,(n1+n2))
    copied,copied1,copied2:=0,0,0
    //var i int

    for ;copied1<n1 && copied2<n2; {
        if data[first+copied1] <data[first+n1+copied2] {
            temp[copied+1]=data[first+(copied1+1)]

        }else{
            temp[copied+1] = data[first + n1 + (copied2+1)]
      }

    }

    for ;copied1<n1; {
        temp[copied+1] = data[first + (copied1+1)]
       }
       for ;copied2<n2; {
            temp[copied+1] = data[first + n1 + (copied2+1)]
        }
    for i:=0;i<n1+n2;i++ {
        data[first + i] = temp[i];
    }
}
func main() {
    f:=fact(4)
    fmt.Println("factorial : ",f)
    for i:=1;i<=4;i++ {

    fmt.Println("fibanocci : ",fiba(i))

    }
    a:=[]int{10,1,33,12,99,100}
     mergeSort(a,1,len(a)-2)
     fmt.Println(a)

}
