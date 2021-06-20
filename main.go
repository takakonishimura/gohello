package main

import "fmt"
import push "hello/pusheen"
// import push "hello/pusheen/baby"
// import push "github.com/takakonishimura/gopusheen/pusheen"
// import push "github.com/takakonishimura/gopusheen/pusheen/baby"

func main() {
    printHello(push.PackageName())
}

func init() {
    fmt.Println("\nexecutes during initialization...")
}
