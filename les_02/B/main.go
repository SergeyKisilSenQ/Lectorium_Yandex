package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Есть 10 вертикальных палочек, пронумерованных от 0 до 9 и n колец трёх цветов красного «R», зеленого «G» и голубого «B», которые на надеты на палочки. Ваша задача по строке
s, кодирующей, где находится каждое из колец определить количество палочек, на которое надеты кольца всех трёх цветов.
Строка представляет из себя последовательность чётной длины, где на нечётных позициях (1,3,5 и т.д.) написан цвет кольца, а на чётных позициях (
2,4,6 и т.д.) — номер палочки, на которую надето кольцо. Таким образом, кольцо номер i имеет цвет s2 i −1 и находится на палочке номер
s2i.
Например, строка «R2G1R1B2G2» кодирует 5 колец:
Первое кольцо имеет красный цвет и находится на палочке 2;
Второе кольцо имеет зеленый цвет и находится на палочке 1;
Третье кольцо имеет красный цвет и находится на палочке 1;
Четвертое кольцо имеет синий цвет и находится на палочке 2;
Пятое кольцо имеет зеленый цвет и находится на палочке 2;
Формат ввода
Первая строка входных данных  — это непустая строка четной длины
s (1≤|s|≤1000), состоящая из символов «R», «G» или «B» и цифр от
0 до 9.
Формат вывода
Выведите единственное целое число — количество палочек, на которых имеются кольца всех трёх цветов.
*/

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
