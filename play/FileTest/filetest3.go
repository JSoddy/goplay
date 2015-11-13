package main

import (
	"fmt"
	//"os"
	//"bufio"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	b, err := ioutil.ReadFile("D:\\James\\Downloads\\kargerMinCut.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	lines := strings.Split(string(b), "\n")

	arr := make([][]int, 0, len(lines))

	for _, v := range lines {
		textints := strings.Split(v, "\t")
		subarr := make([]int, 0, len(textints))
		for _, w := range textints {
			n, err := strconv.Atoi(w)
			if err != nil {
				fmt.Println("Error yo: ", err)
				continue
			}
			subarr = append(subarr, n)
		}
		arr = append(arr, subarr)
	}
	for _, v := range arr {
		for _, w := range v {
			fmt.Print(w, "\t")
		}
		fmt.Print("\r\n")
	}
}
