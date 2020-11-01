package main

/*
   author: aditya.deepak@swiggy.in
*/

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func scanInput(s string, v ...interface{}) {
	_, _ = fmt.Fscanf(reader, s, v...)
}

func printResult(s string, v ...interface{}) {
	_, _ = fmt.Fprintf(writer, s, v...)
}

func getSequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// entry point
func main() {
	defer writer.Flush()

}
