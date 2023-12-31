package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

/*
По заданной строке, являющейся абсолютным адресом в Unix-системе, вам необходимо получить канонический адрес.

В Unix-системе "." соответсвутет текущей директории, ".." — родительской директории, при этом будем считать, что любое количество точек подряд, большее двух, соответствует директории с таким названием (состоящем из точек). "/" является разделителем вложенных директорий, причем несколько "/" подряд должны интерпретироваться как один "/".

Канонический путь должен обладать следующими свойствами:

1) всегда начинаться с одного "/"

2) любые две вложенные директории разделяются ровно одним знаком "/"

3) путь не заканчивается "/" (за исключением корневой директории, состоящего только из символа "/")

4) в каноническом пути есть только директории, т.е. нет ни одного вхождения "." или ".." как соответствия текущей или родительской директории

Формат ввода
Вводится строка с абсолютным адресом, её длина не превосходит 100.

*/

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var src string
	var res []string
	fmt.Fscan(in, &src)
	array := strings.Split(src, "/")
	for i := 0; i < len(array); i++ {
		if array[i] == "" || array[i] == "." {
			continue
		}
		if array[i] == ".." {
			if len(res) >= 1 {
				res = append(res[:len(res)-1], res[len(res):]...)
			}
			continue
		}
		res = append(res, array[i])
	}
	if len(res) == 0 {
		out.WriteString("/")
	} else {
		buff := bytes.Buffer{}
		for _, v := range res {
			buff.WriteString("/")
			buff.WriteString(v)
		}
		out.WriteString(buff.String())
	}
}
