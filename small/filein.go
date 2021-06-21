package main

import(
    "bufio"
    "os"
    "fmt"
    "log"
    "strings"
)

func main() {
    counts := make(map[string]int)
    files := os.Args[1:]
    if len(files) == 0 {
        log.Fatal("no files provided")
    }

    for _, filename := range files {
        f, err := os.Open("sample_files/" + filename)
        if err != nil {
            fmt.Fprintf(os.Stderr, "opening file failed with error %v\n", err)
            continue
        }
        countLines(f, counts)
        f.Close()
    }

    for word, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, word)
        }
    }
}

func countLines(f *os.File, counts map[string]int) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        tokens := strings.Split(input.Text(), " ")
        for _, tkn := range tokens {
            counts[tkn]++
        }
    }
} 

