/*
 * lets user enter number and prints
 * their mean ,median in a tabular form
 * exercise 15.6
 */

 package main

 import (
    "fmt"
    "net/http"
    "strings"
    "strconv"
 )

 const form=`<html><body><form action="#" method="post" value="bar">
             <h2>Statistics</h2><br>
             <h4>Computes basic statistics for a given list of numbers</h4><br>
             <h4>Numbers (comma-separated)</h4><br>
             <input type="text" value="in"/>
             <input type="submit" value="Calculate"/>
             </form></body></html>`

 //const result=`<html><body>`


 func showForm(w http.ResponseWriter,req *http.Request) {
    w.Header().Set("Content-Type","text/html")
    switch req.Method  {
    case "GET":fmt.Fprint(w,form);
    case "POST" : //msg:=fmt.Fprint(w,req.FormValue("in"))
                  //fmt.Println(msg)
                  msg:=req.FormValue("in")
                  arr:=strings.Split(msg,",")
                  //fmt.Fprintf(w,arr,"<h2>%s</h2>")
                  count:=len(arr)
                  sum:=0
                  for _,val:=range arr {
                    i,_:=strconv.Atoi(val)
                    fmt.Fprintf(w,val,"<h2>%s</h2>")
                    sum+=i
                  }
                  mean:=sum/count
                  fmt.Fprintf(w,"<h1>Result</h1> <table border='1'> <tr> <td> Numbers </td> <td>%d</td></tr> <tr><td>Mean</td> <td>%d</td></tr>",msg,mean)

    }


 }
 func main() {
    http.HandleFunc("/formShow",showForm)
    //http.HandleFunc("/showData",showData)
    if err:=http.ListenAndServe(":8080",nil);err!=nil {
        panic(err)
    }
 }
