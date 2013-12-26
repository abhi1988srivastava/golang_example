/*
 *
 * net/http demonstration
 *
 */

 package main

 import (
    "fmt"
    "net/http"
 )

 func HandleServer(w http.ResponseWriter,req *http.Request) {
    fmt.Println("Insid handleServer function:")
    fmt.Fprint(w,"hello"+req.URL.Path[1:])
 }
 func main() {
    /*
     * here / = http://localhost:8080. So if the request come to root (/), then the function HandleServer
     * is executed
     *
     * use http.ListenAndServeTLS() if use want https connection
     */

    http.HandleFunc("/",HandleServer)
    res:=http.ListenAndServe("localhost:8080",nil)
    if res!=nil {
        fmt.Println("problem with listening at 8080")
        return
    }

 }
