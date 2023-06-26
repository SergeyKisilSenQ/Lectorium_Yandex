package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Вам дана последовательность измерений некоторой величины. Требуется определить, повторялась ли какое-либо число, причём расстояние между повторами не превосходило k.

Формат ввода
В первой строке задаются два числа n и k (1 ≤ n, k ≤ 105).

Во второй строке задаются n чисел, по модулю не превосходящих 109.

Формат вывода
Выведите YES, если найдется повторяющееся число и расстояние между повторами не превосходит k и NO в противном случае.

*/

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n, k int
	fmt.Fscanln(in, &n, &k)
	numsMap := make(map[int]int, n)
	src, _ := in.ReadString('\n')
	src = strings.TrimRight(src, "\r\n")
	array := make([]int, n)
	for i, v := range strings.Fields(src) {
		array[i], _ = strconv.Atoi(v)
	}
	var found bool
	var num int
	for i := 0; i < n; i++ {
		num = array[i]
		if !found {
			if v, ok := numsMap[num]; !ok {
				numsMap[num] = i
			} else if i-v <= k {
				found = true
				break
			} else {
				numsMap[num] = i
			}
		}
	}
	if found {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
	}
}
