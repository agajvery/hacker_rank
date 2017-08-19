package main

import (
	"fmt"
	"strconv"
    "bufio"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    // n, _ := reader.ReadString('\n')

    // num_len, _ := strconv.ParseInt(strings.TrimRight(n, "\n"), 10, 64)
    num_len := 6
    start_value, _ := reader.ReadString('\n')
    num_start, _ := strconv.ParseInt(strings.TrimRight(start_value, "\n"), 10, 32)
    num_start = num_start + 1 
    next_value := strconv.FormatInt(num_start, 10)

	_, lucky_num := find_lucky(next_value, int(num_len))
	fmt.Printf("%v", lucky_num)
}

func find_lucky(start_value string, num_len int)(bool, int) {
	var(
        is_lucky bool
        lucky_number int
        f_sum int64
        s_sum int64
    ) 

    half := num_len / 2
// fmt.Printf("%v, %v, %v", start_value, half, num_len)
// return false, 0
	if num_len % 2 == 0 {
		i := 0
        for _, rn := range start_value {
			num, _ := strconv.ParseInt(string(rn), 10, 64)
			if i < half {
                f_sum += num
            } else {
                s_sum += num
            }
			i++
		}
        if f_sum == s_sum {
            is_lucky = true
            sum, _ := strconv.ParseInt(start_value, 10, 64)
            lucky_number = int(sum)
        }
	}

    if (is_lucky == false) {
        num_start, _ := strconv.ParseInt(start_value, 10, 32)
        num_start = num_start + 1 
        next_value := strconv.FormatInt(num_start, 10)
        // fmt.Printf("%v", next_value)
        return find_lucky(next_value, num_len)    
    }
	return is_lucky, lucky_number
}
