package main

import (
	"bufio"
	"os"
	"strings"
)

const maxLeightWord = 1000

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	src, _ := in.ReadString('\n')
	src = strings.TrimRight(src, "\r\n")
	dictionary := AddToMap(src)
	src, _ = in.ReadString('\n')
	src = strings.TrimRight(src, "\r\n")
	array := strings.Fields(src)
	res := make([]string, len(array))
	for i, word := range array {
		if ok, newWord := Check(word, dictionary); ok {
			res[i] = newWord
		} else {
			res[i] = word
		}
	}
	out.WriteString(strings.Join(res, " "))
}

func AddToMap(src string) map[string]struct{} {
	array := strings.Fields(src)
	dictionary := make(map[string]struct{})
	for _, word := range array {
		dictionary[word] = struct{}{}
	}
	return dictionary
}

func Check(word string, dictionary map[string]struct{}) (bool, string) {
	r := 1
	for i := 0; i < len(word); i++ {
		if i == maxLeightWord {
			break
		}
		if _, ok := dictionary[word[0:r]]; ok {
			return true, word[0:r]
		}
		r++
	}
	return false, ""
}
