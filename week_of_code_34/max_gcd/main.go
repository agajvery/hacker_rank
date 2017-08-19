package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
    "math"
)

func main(){
    var (
        g int64
        max_g int64
        max_sum int64
        input_str string
    )
    
    reader := bufio.NewReader(os.Stdin)
    _, _ = reader.ReadString('\n')
    
    input_str, _ = reader.ReadString('\n')
    input_str = strings.TrimRight(input_str, "\n")
    m_array := strings.Split(input_str, " ")

    input_str, _ = reader.ReadString('\n')
    input_str = strings.TrimRight(input_str, "\n")
    n_array := strings.Split(input_str, " ")

    // fmt.Printf("%v\n %v", m_array, n_array)
    // return

    for _, m_str := range m_array {
        for _, n_str := range n_array {
            m, _ := strconv.ParseInt(m_str, 10, 64)
            n, _ := strconv.ParseInt(n_str, 10, 64)
            
            g = gcd(m, n)
            if g > max_g {
                max_g = g
                max_sum = m + n
            }
            // fmt.Printf("M = %v, N = %v, G = %v\n", m, n, g)
        }
    }
    fmt.Printf("%v", max_sum)
}

func gcd_bin(m uint64, n uint64) int64 {
    if m == 0 && n == 0 {
        return 0
    }
    if m == 0 {
        return n
    }
    return gcd_1(m, n % m)        
}

func gcd(m int64, n int64) int64 {
    if m < 0 {
        m = m * -1
    }
    if n < 0 {
        n = n * -1
    }
    if m == 0 && n == 0 {
        return 0
    }

    if m == 0 || n == 0 {
        return m + n
    }

    if m > n {
        m = m % n
    } else {
        n = n % m
    }
    return gcd(m, n)
}