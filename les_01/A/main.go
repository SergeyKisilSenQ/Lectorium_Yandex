package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n int
	d := 1
	fmt.Fscan(in, &n)
	res := 1 + (n-1)*d
	out.WriteString(strconv.Itoa(res))
}
