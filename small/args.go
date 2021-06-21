package main

import (
    "os"
    "fmt"
    "strings"
)

func main() {

 /* FOR LOOP
    var s, separator string
    for i := 1; i < len(os.Args); i++ {
        s += separator + os.Args[i]
        separator = " & "
    }
    fmt.Println(s)
  */

    fmt.Println(strings.Join(os.Args[1:], " & "))

 /* QUICK DEBUG
    fmt.Println(os.Args[1:])
  */
}

