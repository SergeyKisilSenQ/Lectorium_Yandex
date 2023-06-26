package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n, checkNum, majority int
	fmt.Fscanln(in, &n)
	if n%2 == 0 {
		checkNum = n / 2
	} else {
		checkNum = (n / 2) + 1
	}
	src, _ := in.ReadString('\n')
	src = strings.TrimRight(src, "\r\n")
	array := strings.Fields(src)
	numMaps := make(map[int]int)
	for i := 0; i < len(array); i++ {
		num, _ := strconv.Atoi(array[i])
		numMaps[num]++
		if numMaps[num] >= checkNum {
			if majority < numMaps[num] {
				majority = num
			}
		}
	}
	fmt.Fprint(out, majority)
}
