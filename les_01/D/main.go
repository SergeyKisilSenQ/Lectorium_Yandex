package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var count int
	fmt.Fscanln(in, &count)

	line, _ := in.ReadString('\n')
	out.WriteString(strconv.Itoa(GetMinTime(count, line)))
}

func GetMinTime(count int, src string) int {
	lineTrim := strings.TrimSpace(src)
	lineSplit := strings.Split(lineTrim, " ")
	data := make([]int, count)
	for i := 0; i < count; i++ {
		data[i] = GetTimeFromString(lineSplit[i])
	}
	sort.Ints(data)
	if data[1] == 0 {
		return 0
	}
	res := 1440
	for i := 1; i < count; i++ {
		res = GetMin(res, GetMin(data[i]-data[i-1], 1440-data[i]+data[0]))
	}
	return res
}

func GetTimeFromString(src string) int {
	lineSplit := strings.Split(src, ":")
	h, _ := strconv.Atoi(lineSplit[0])
	m, _ := strconv.Atoi(lineSplit[1])
	return h*60 + m
}

func GetMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}
