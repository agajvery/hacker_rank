package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var sum int64
	var i int64

	reader := bufio.NewReader(os.Stdin)
	inputStrLen, _ := reader.ReadString('\n')
	arrayLen, _ := strconv.ParseInt(strings.TrimRight(inputStrLen, "\n"), 10, 32)
	inputStrArray, _ := reader.ReadString('\n')

	values := strings.Split(inputStrArray, " ")
	for ; i < arrayLen; i++ {
		value, _ := strconv.ParseInt(values[i], 10, 64)
		sum += value
	}
	fmt.Println(sum)
}
