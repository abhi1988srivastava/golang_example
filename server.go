/*
 * client server demonstration using net package
 */

 //server file

 package main

 import (
    "fmt"
    "net"
 )

 func main() {
    fmt.Println("Starting server")
    listener,err:=net.Listen("tcp","localhost:50000")
    if err!=nil {
        fmt.Println("could start the server at this port or host")
        return
    }

    for  {
        conn,err:=listener.Accept()
        fmt.Println("connection :",conn)
        if err!=nil {
            fmt.Println("error listening to client :")
            return
        }
        //spawning for every client request
        go processRequest(conn)
    }

 }

 func processRequest(conn net.Conn) {
    fmt.Println("going to read clients data")
    buf:=make([]byte,1024)
    _,err:=conn.Read(buf)
    if err!=nil {
        fmt.Println("error in reading from client")
    }
    fmt.Println("Client said :",string(buf))
 }
