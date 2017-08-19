package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    s, _ := reader.ReadString('\n')

    items_str, _ := reader.ReadString('\n')
    items := strings.Split(items_str, " ")
    
    fmt.Printf(s)
    fmt.Printf("%v", items)
}