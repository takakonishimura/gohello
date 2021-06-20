package main

import (
    "os"
    "fmt"
)

func main() {
    var s, separator string
    for i := 1; i < len(os.Args); i++ {
        s += separator + os.Args[i]
        separator = " & "
    }
    fmt.Println(s)
}

