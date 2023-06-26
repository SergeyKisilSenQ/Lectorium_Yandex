package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

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
