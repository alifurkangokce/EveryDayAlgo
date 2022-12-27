package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func wayTooLong() {
	reader := bufio.NewReader(os.Stdin)
	tTemp, _ := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	var result string
	for i := 0; i < int(tTemp); i++ {
		s := readLine(reader)
		l := len(s)
		if len(s) > 10 {
			result += string(s[0]) + strconv.Itoa(len(s[1:(l-1)])) + string(s[l-1]) + "\n"
		} else {
			result += s + "\n"
		}
	}
	fmt.Print(result)
}
func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
