package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const toSolve = 2

func main() {

	if toSolve == 1 {
		solveOne()
	}
	if toSolve == 2 {
		solveTwo()
	}

}

func solveOne() {
	start := 0
	freqs := readInputfile()
	for _, sum := range freqs {
		start += sum
	}

	fmt.Println(start)
}

func solveTwo() {
	start := 0
	freqs := readInputfile()
	sumsMap := map[int]bool{0: true}
	for {
		for _, sum := range freqs {
			start += sum
			if _, ok := sumsMap[start]; ok {
				fmt.Println(start)
				os.Exit(0)
			}
			sumsMap[start] = true
		}
		fmt.Println("Current sum", start)
	}
}

func readInputfile() []int {
	data, err := ioutil.ReadFile("input.txt")
	checkErr(err)
	content := strings.Split(string(data), "\n")
	var freqs []int
	for _, curr := range content {
		num, err := strconv.Atoi(curr)
		if err == nil {
			freqs = append(freqs, num)
		}
	}

	return freqs
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
