package main

import "strings"

func echo1(args []string) string {
	str, sep := "", ""
	for _, arg := range args[:] {
		str += sep + arg
		sep = " "
	}
	return str
}

func echo2(args []string) string {
	return strings.Join(args[:], " ")
}

func main() {}
