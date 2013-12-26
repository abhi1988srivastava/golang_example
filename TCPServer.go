/*
 *
 * server using TCPServer
 */

 package main

 import (
    "fmt"
    "flag"
    "net"
 )

 func main() {
    flag.Parse()
    if flag.NArg()!=2 {
        fmt.Println("Please give two command line arguments : host and port")
    }
    hostport:=fmt.Sprintf("%s:%s",flag.Arg(0),flag.Arg(1))
    servadrr,err:=net.ResolveTCPAddr("tcp",hostport)
    if err!=nil {
        fmt.Println("error parsing hostport")
    }
    listener,err:=net.ListenTCP("tcp",servadrr)
    if err!=nil {
        fmt.Println("error listening on the given TCPAddr")
    }

    for{
        conn,err:=listener.Accept()
        if err!=nil {
            fmt.Println("not able to accept connections")
        }
        go processRequest(conn)
    }
 }

 func processRequest(conn net.Conn) {
    connFrom:=conn.RemoteAddr().String()
    fmt.Println("Connection from :",connFrom)
    str:="Say something to me"
    _,err:=conn.Write([]byte(str))
    if err!=nil {
       fmt.Println("error in writing to client")
    }
    buf:=make([]byte,4096)
    _,err1:=conn.Read(buf)
    if err1!=nil {
        fmt.Println("error in reading from client")
    }
    fmt.Println("client says :",string(buf))
 }


