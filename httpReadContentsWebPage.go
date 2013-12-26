/*
 *
 * read contents of a web page
 * html content of a web page with http.Get(); the response res
 * returned from Get has the content in a field Body, which is read with ioutil.ReadAll:
 *
 */

 package main

 import (
    "fmt"
    "net/http"
    "io/ioutil"
 )

 func main() {
    res,err:=http.Get("http://google.com")
    if err!=nil {
        fmt.Println("problem in reachin the web page")
    }
    data,err:=ioutil.ReadAll(res.Body)
    if err!=nil {
        fmt.Println("error reading body")
    }
    fmt.Println("body content is :",string(data))

 }
