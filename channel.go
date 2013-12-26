/*
 * channel demonstration to communicate between goroutines

 package main

 import (
   "fmt"
   "time"
 )

 func main() {
   go greet("foo")
 }

 func greet(name string) {
   time.Sleep(1 * time.Second)
   fmt.Printf("Hi %s\n", name)
 }


if we write like this, there will be no output sincee the main goroutine will exit even
before the greet goroutine has finished executing..
Solution: to manage this,we need channel to accept and receive the content. Here channel is used to
communicate between the greet goroutine and main goroutine
*/

package main

import (
  "fmt"
  "time"
)

func main() {
  ch := make(chan string)

  go greet("Abhinav", ch)

  fmt.Println(<-ch)
}

func greet(name string, ch chan string) {
  time.Sleep(1 * time.Second)
  ch <- fmt.Sprintf("Hi %s", name)
}
