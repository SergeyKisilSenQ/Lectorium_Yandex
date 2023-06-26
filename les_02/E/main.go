package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
Дан массив a из n чисел. Найдите минимальное количество чисел, после удаления которых попарная разность оставшихся чисел по модулю не будет превышать
1, то есть после удаления ни одно число не должно отличаться от какого-либо другого более чем на 1.
Формат ввода
Первая строка содержит одно целое число
n (1≤n≤2⋅105) — количество элементов массива a.
Вторая строка содержит n целых чисел a1,a2,…,an (0≤ai≤105) — элементы массива a.

Формат вывода
Выведите одно число — ответ на задачу.
*/
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n, maxNum, count int
	fmt.Fscanln(in, &n)
	src, _ := in.ReadString('\n')
	src = strings.TrimRight(src, "\r\n")
	array := strings.Fields(src)
	numArray := make([]int, n)
	for i := 0; i < n; i++ {
		numArray[i], _ = strconv.Atoi(array[i])
	}
	if n == 1 {
		fmt.Fprint(out, 0)
		return
	}
	sort.Ints(numArray)
	l := 0
	count = 1
	for i := 1; i < n; i++ {
		if numArray[i]-numArray[l] <= 1 {
			count++
			if maxNum < count {
				maxNum = count
			}
		} else {
			if maxNum < count {
				maxNum = count
			}
			l++
			if numArray[i]-numArray[l] <= 1 {
				count = len(numArray[l : i+1])
			} else {
				count = 1
			}
		}
	}
	fmt.Fprint(out, n-maxNum)
}
