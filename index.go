package main

import (
	day1 "adventofcode/day1"
	day2 "adventofcode/day2"
	util "adventofcode/util"
)

import (
	"fmt"
	"os"
	"strconv"
)

type DayFunc func([]string)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("A day must be specified")
		os.Exit(1)
	}

	dayFuncs := []DayFunc{day1.Run, day2.Run}
	day, err := strconv.Atoi(os.Args[1])
	util.Check(err)

	if day > len(dayFuncs) {
		fmt.Println("Invalid day")
		os.Exit(1)
	}

	arguments := os.Args[2:]

	dayFuncs[day-1](arguments)
}
