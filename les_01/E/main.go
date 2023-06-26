package main

import "fmt"

func main() {
	var scr string
	fmt.Scan(&scr)
	res := []byte(scr)
	if len(scr) == 1 {
		fmt.Println(" ")
	} else {
		fmt.Println(BrokenPoli(scr, res))
	}
}

func BrokenPoli(scr string, res []byte) string {
	for i := 0; i < len(scr)/2; i++ {
		if scr[i] != 'a' {
			res[i] = 'a'
			return string(res)
		}
	}
	res[len(res)-1] = 'b'
	return string(res)
}
