package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const money = 1000

func main() {
	in := bufio.NewReader(os.Stdin)
	src, _ := in.ReadString('\n')
	src = strings.TrimRight(src, "\r\n")
	n, _ := strconv.Atoi(src)
	array := make([]float64, n)
	src, _ = in.ReadString('\n')
	src = strings.TrimRight(src, "\r\n")
	inArray := strings.Split(src, " ")
	for i, s := range inArray {
		array[i], _ = strconv.ParseFloat(s, 64)
	}
	if n == 1 {
		fmt.Println("0 0")
		return
	}
	l, r := 0, 1
	buyPrice := l
	sellPrice := r
	for i := 1; i < n; i++ {
		if money/array[i] > money/array[l] {
			l = i
			r = i
		}
		if array[i]*(money/array[l]) > array[r]*(money/array[l]) {
			r = i
		}
		if array[r]*(money/array[l]) > array[sellPrice]*(money/array[buyPrice]) {
			buyPrice = l
			sellPrice = r
		}
	}

	if array[sellPrice]-array[buyPrice] <= 0 {
		fmt.Println("0 0")
		return
	}
	fmt.Println(buyPrice+1, sellPrice+1)
}
