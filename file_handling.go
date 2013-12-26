//file reading and writing
//stackoverflow



package main

import (
    "bufio"
    "io"
    "os"
    "io/ioutil"
)

func main() {
    // open input file
    fi, err := os.Open("example_27.go")
   /* if err != nil { panic(err) }
    // close fi on exit and check for its returned error
    defer func() {
        if err := fi.Close(); err != nil {
            panic(err)
        }
    }()
    // make a read buffer
    */
    r := bufio.NewReader(fi)

    // open output file
    fo, err := os.Create("output1.txt")
    /*if err != nil { panic(err) }
    // close fo on exit and check for its returned error
    defer func() {
        if err := fo.Close(); err != nil {
            panic(err)
        }
    }()
    // make a write buffer
    */
    w := bufio.NewWriter(fo)

    // make a buffer to keep chunks that are read
    buf := make([]byte, 1024)
    for {
        // read a chunk
        n, err := r.Read(buf)
        if err != nil && err != io.EOF { panic(err) }
        if n == 0 { break }

        // write a chunk
        if _, err := w.Write(buf[:n]); err != nil {
            panic(err)
        }
    }

    if err = w.Flush(); err != nil { panic(err) }

    //when we know, we are dealing with small files, the below code will read the complete file

    b, err := ioutil.ReadFile("output.txt")
    if err != nil { panic(err) }

    // write whole the body : 0644 filepermission
    err = ioutil.WriteFile("output2.txt", b, 0644)
    if err != nil { panic(err) }



}
