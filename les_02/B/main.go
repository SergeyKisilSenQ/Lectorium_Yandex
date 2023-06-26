package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var count int
	resMap := make(map[byte][]byte)
	src, _ := in.ReadString('\n')
	src = strings.TrimSpace(strings.TrimRight(src, "\r\n"))
	r := 1
	for i := 0; i < len(src); i += 2 {
		resMap[src[i+r]] = append(resMap[src[i+r]], src[i])
	}
	for _, bytes := range resMap {
		if len(bytes) >= 3 {
			if Check(bytes) {
				count++
			}
		}
	}
	fmt.Fprint(out, count)
}

func Check(array []byte) bool {
	var red, green, blue bool
	for i := 0; i < len(array); i++ {
		switch array[i] {
		case 'R':
			red = true
		case 'G':
			green = true
		case 'B':
			blue = true
		}
	}
	if red && green && blue {
		return true
	}
	return false
}
