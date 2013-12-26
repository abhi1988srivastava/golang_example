/*
 * lets user enter number and prints
 * their mean ,median in a tabular form
 * exercise 15.6
 */

 package main

 import (
    "fmt"
    "net/http"
 )

 const form=`<html><body><form action="#" method="post" value="bar">
             <h2>Statistics</h2><br>
             <h4>Computes basic statistics for a given list of numbers</h4><br>
             <h4>Numbers (comma or space-separated)</h4><br>
             <input type="text" value="in"/>
             <input type="submit" value="Calculate"/>
             </form></body></html>`

 func showForm(w http.ResponseWriter,req *http.Request) {
    w.Header().Set("Content-Type","text/html")
    switch req.Method  {
    case "GET":fmt.Fprint(w,form);
    case "POST" : //msg:=fmt.Fprint(w,req.FormValue("in"))
                  //fmt.Println(msg)
    }


 }
 func main() {
    http.HandleFunc("/formShow",showForm)
    //http.HandleFunc("/showData",showData)
    if err:=http.ListenAndServe(":8080",nil);err!=nil {
        panic(err)
    }
 }
