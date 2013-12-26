/*
 * creating a web page with form submission
 */

 package main

 import (
    "fmt"
    "net/http"

 )

 //using # to be on the same page
 const form=`<html><body><form action="#" method="post" name="bar">
              <input type="text" name="in"/>
              <input type="submit" name="Submit"/>
              </form></body></html>`

 func simpleServer(w http.ResponseWriter,req *http.Request) {
    fmt.Fprint(w,"<h1>I am Abhinav</h1>")
 }

 func formServer(w http.ResponseWriter,req *http.Request) {
    w.Header().Set("Content-Type","text/html")
    switch req.Method {
    case "GET": fmt.Fprint(w,form);
    case "POST" : //we need to call ParseForm before we can extract data like
                  //req.ParseForm()
                  //fmt.Fprint(w,req.Form["in"][0])
                  fmt.Fprint(w,req.FormValue("in"))
                  fmt.Println(msg)

                  fmt.Fprint(w," Am I not sexy? :P")

    }
 }
 func main() {
    http.HandleFunc("/test1",simpleServer)
    http.HandleFunc("/test2",formServer)
    if err:=http.ListenAndServe(":8080",nil);err!=nil {
        panic(err)
    }

 }
