package main

import (
	"./cdjm"
	"fmt"
	"strconv"
	"strings"
)

func serializeKeyboard(keyboardToLetters map[string]string) (letterToPressedKeys map[string]string) {
	//change keybord map into, "letter" to "pressed keys"

	letterToPressedKeys = make(map[string]string)

	for k, v := range nums {
		chars := strings.Split(v, "")
		for i, char := range chars { //her bir letter'ı ele alalım
			press := ""
			for a := 0; a < i+1; a++ { //letter kaçıncı sıradaysa o kadar o tuşa basılacak
				press = press + k
			}
			letterToPressedKeys[char] = press
		}

	}
	//example of desired map:
	//map[p:7 i:444 t:8 m:6 e:33 a:2 y:999  :0 q:77 n:66 u:88 j:5 b:22 z:9999 f:333 o:666 r:777 k:55 c:222 v:888 g:4 h:44 w:9 l:555 s:7777 d:3 x:99]

	return
}

func hmnToSeq(humanStr string, lttrs map[string]string) (sequence string) {
	humanStrArr := strings.Split(humanStr, "")
	for i, v := range humanStrArr {
		if i > 0 && lttrs[humanStrArr[i]][0] == lttrs[humanStrArr[i-1]][0] { //bir önceki letter'ı yazmak için de aynı tuşu mu kullandık?
			sequence = sequence + " " + lttrs[v]
		} else {
			sequence = sequence + lttrs[v]
		}
	}
	return
}

//Easier to define keyboards
var nums = map[string]string{
	"0": " ",
	"1": "",
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func main() {

	lttrs := serializeKeyboard(nums)
	//fmt.Println(strings.Split(nums["1"], ""))

	inArr, _ := cdjm.Input("input/C-large-practice.in")
	size, _ := strconv.Atoi(inArr[0])
	//a == a[0:] // its a reminder how it includes number on the left
	inArr = inArr[1:]

	for i := 0; i < size; i++ {
		inArr[i] = hmnToSeq(inArr[i], lttrs)

		fmt.Println(inArr[i])
	}

	cdjm.Output("output/now.out", inArr)
}
