package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
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
	inputStrings := readInputfile()

	twoCount := 0
	threeCount := 0

	for _, str := range inputStrings {
		out := getLetterMapForString(str)
		if hasCount(out, 2) {
			twoCount++
		}
		if hasCount(out, 3) {
			threeCount++
		}
	}
	fmt.Println(twoCount * threeCount)
}

func solveTwo() {
	inputStrings := readInputfile()

	strOne, strTwo, err := getNearlyIdenticalStrings(inputStrings)
	checkErr(err)
	fmt.Println(getIdenticalCharacters(strOne, strTwo))

}

func getNearlyIdenticalStrings(inputList []string) (string, string, error) {
	for k, str := range inputList {
		for _, tocheck := range inputList[k+1:] {
			if isStringNearlyEqual(str, tocheck) {
				fmt.Println(str, tocheck)
				return str, tocheck, nil
			}
		}
	}
	return "", "", errors.New("Did not find two nearly identical strings")
}

func getIdenticalCharacters(input, other string) string {
	strToBuild := ""
	for key, c := range input {
		if string(other[key]) == string(c) {
			strToBuild += string(c)
		}
	}
	return strToBuild
}

func isStringNearlyEqual(input, toCheckAgainst string) bool {
	faultycount := 0
	for key, c := range input {
		if string(toCheckAgainst[key]) != string(c) {
			faultycount++
			if faultycount >= 2 {
				return false
			}
		}
	}
	return true
}

func hasCount(input map[string]int, count int) bool {
	for _, v := range input {
		if v == count {
			return true
		}
	}
	return false
}

func getLetterMapForString(check string) map[string]int {
	letterMap := make(map[string]int)

	for _, char := range check {
		letterMap[string(char)] = letterMap[string(char)] + 1
	}

	return letterMap
}

func readInputfile() []string {
	data, err := ioutil.ReadFile("input.txt")
	checkErr(err)

	content := strings.Split(string(data), "\n")
	var boxes []string

	for _, curr := range content {
		//We probably have a new line at the end of the file which we dont want to include
		if len(curr) > 0 {
			boxes = append(boxes, curr)
		}
	}

	return boxes
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
