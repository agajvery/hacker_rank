package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var scoreA, scoreB int

	reader := bufio.NewReader(os.Stdin)
	inputStrArrayA, _ := reader.ReadString('\n')
	inputStrArrayB, _ := reader.ReadString('\n')

	arrayA := strings.Split(strings.TrimRight(inputStrArrayA, "\n"), " ")
	arrayB := strings.Split(strings.TrimRight(inputStrArrayB, "\n"), " ")
	arrayLen := len(arrayA)

	for i := 0; i < arrayLen; i++ {
		a, errorA := strconv.ParseInt(arrayA[i], 10, 64)
		b, errorB := strconv.ParseInt(arrayB[i], 10, 64)

		if errorA == nil && errorB == nil {
			if a > b {
				scoreA++
			}
			if b > a {
				scoreB++
			}
		} else {
			fmt.Printf("%v %v\n", errorA, errorB)
		}
	}
	fmt.Printf("%v %v", scoreA, scoreB)
}
