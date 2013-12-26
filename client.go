//client file

package main

import (
    "fmt"
    "net"
    "os"
    //"bufio"
    //"strings"
    "text/scanner"
)

func main() {
    conn,err:=net.Dial("tcp","localhost:50000")
    if err!=nil {
        fmt.Println("error in connecting to server")
        return
    }

    //inputReader:=bufio.NewReader(os.Stdin)
    //scanner:=bufio.NewScanner(os.Stdin)
    //fmt.Println("Say name to server")
    //clientName,_:=inputReader.ReadString('\n')
    var s scanner.Scanner
    s.Init(os.Stdin)
    tok:=s.Scan()
    //var clientName string
    for tok!=scanner.EOF {
        tok=s.Scan()
        fmt.Println(tok)
        //clitName=+str
    }
    //clientName:=scanner.Text()
//    trimmedClient := strings.Trim(clientName, "\r\n")
    //trimmedClient := strings.Trim(tok, "\r\n")
    trimmedClient:=tok

    for  {
        fmt.Println("Sending data to server,type Q to quit")
        //input,_:=inputReader.ReadString('\n')
        //input:=scanner.Text()
        //var input string
        tok1:=s.Scan()
        for tok1!=scanner.EOF {
           //str1:=scanner.Text()
           //input=+str1
           tok1=s.Scan()
        }
        //trimmedInput := strings.Trim(tok1, "\r\n")
        trimmedInput:=tok1
        if string(trimmedInput)=="Q" {
            fmt.Println("No more to write")
            return
        }
        _,err:=conn.Write([]byte(string(trimmedClient)+" says :"+string(trimmedInput)))
        if err!=nil {
            fmt.Println("error writing to server")
            return
        }
    }
}
