package main

import (
	"fmt"
	"os"
	"strings"
)

func getLeet(key string) (string, bool) {
	leetspeak := map[string]string{
		"A": "4",
		"B": "8",
		"C": "(",
		"E": "3",
		"G": "6",
		"H": "#",
		"I": "1",
		"J": "_",
		"L": "1_",
		"O": "0",
		"Q": "0_",
		"S": "5",
		"T": "7",
		"V": "/",
		"W": "//",
		"X": "><",
		"Y": "`/",
		"Z": "2",
	}


	val, exists := leetspeak[key]
	return val, exists
}


func main() {

	args := os.Args[1:]
	res := []string{}
	for _, arg := range args {
		buf := []string{}
		for _, c := range arg {
			symbol, ok := getLeet(strings.ToUpper(string(c)))
			if ok {
				buf = append(buf,symbol)
			} else {
				buf = append(buf,string(c))
			}
		}
		res = append(res,strings.Join(buf,""))
	}
	fmt.Println(strings.Join(res,""))
}
