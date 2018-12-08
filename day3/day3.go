package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type field struct {
	Squares []square
	Vectors map[int]map[int]int
}

func (f *field) addSquare(s square) {
	f.Squares = append(f.Squares, s)
}

type square struct {
	Num      int
	FromLeft int
	FromTop  int
	Width    int
	Length   int
}

const toSolve = 1

func main() {
	if toSolve == 1 {
		solveOne()
	}
	if toSolve == 2 {
		solveTwo()
	}
}

func solveOne() {
	f := readInputfile()
	f.Vectors = map[int]map[int]int{}

	for _, s := range f.Squares {
		for i := s.FromLeft; i < s.Width+s.FromLeft; i++ {
			for j := s.FromTop; j < s.Length+s.FromTop; j++ {
				if f.Vectors[i] == nil {
					f.Vectors[i] = map[int]int{}
				}
				f.Vectors[i][j] = f.Vectors[i][j] + 1
			}
		}
	}

	total := 0
	for _, maps := range f.Vectors {
		for _, count := range maps {
			if count > 1 {
				total++
			}
			fmt.Println(count)

		}
	}
	fmt.Println(total)
}

func solveTwo() {

}

func readInputfile() field {
	var theField field
	data, err := ioutil.ReadFile("input.txt")
	checkErr(err)
	content := strings.Split(string(data), "\n")

	for _, curr := range content {
		if len(curr) > 0 {
			theField.addSquare(parseLine(curr))
		}
	}

	return theField
}

/**
 * Example data:
 * #1015 @ 904,392: 19x21
 */
func parseLine(line string) square {
	//trim initial #
	lines := strings.Split(line[1:], " ")

	count, err := strconv.Atoi(lines[0])
	checkErr(err)

	fromLeft, fromTop := parseFromPart(lines[2], ":", ",")

	width, length := parseFromPart(lines[3], "", "x")
	return square{
		Num:      count,
		FromLeft: fromLeft,
		FromTop:  fromTop,
		Width:    width,
		Length:   length,
	}
}

func parseFromPart(s, trim, split string) (int, int) {
	s = strings.Trim(s, trim)
	counts := strings.Split(s, split)

	left, err := strconv.Atoi(counts[0])
	checkErr(err)
	top, err := strconv.Atoi(counts[1])
	checkErr(err)

	return left, top
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
