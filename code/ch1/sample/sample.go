package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
    var i = 0
    for i = 0 ; i < 5 ; i++ {
        fmt.Println(strings.Join(os.Args[1:], " "))
    }
}
