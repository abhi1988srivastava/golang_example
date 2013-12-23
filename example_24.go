/*
 * COMMND LINE ARGUMENTS
 *
 */

package main

import (
 "fmt"
 "os"
)
func main() {
 if len(os.Args) < 2 { // length of argument slice
 os.Exit(1)
 }
 for i := 1; i < len(os.Args); i++ {
 fmt.Printf("arg %d: %s\n", i, os.Args[i])
 }
} // falling off end == os.Exit(0)
