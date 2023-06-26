package main

import (
	"bufio"
	"os"
	"strings"
)

/*
С целью экономии чернил в картридже принтера было принято решение укоротить некоторые слова в тексте. Для этого был составлен словарь слов, до которых можно сокращать более длинные слова. Слово из текста можно сократить, если в словаре найдется слово, являющееся началом слова из текста. Например, если в списке есть слово "лом", то слова из текста "ломбард", "ломоносов" и другие слова, начинающиеся на "лом", можно сократить до "лом".

Если слово из текста можно сократить до нескольких слов из словаря, то следует сокращать его до самого короткого слова.

Формат ввода
В первой строке через пробел вводятся слова из словаря, слова состоят из маленьких латинских букв. Гарантируется, что словарь не пуст и количество слов в словаре не превышет 1000, а длина слов — 100 символов.

Во второй строке через пробел вводятся слова текста (они также состоят только из маленьких латинских букв). Количество слов в тексте не превосходит 105, а суммарное количество букв в них — 106.

Формат вывода
Выведите текст, в котором осуществлены замены.
*/
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
