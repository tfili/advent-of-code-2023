package adventofcode

import util "adventofcode/util"

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run(args []string) {   
	if len(args) == 0 {
		fmt.Println("A file must be specified")
		os.Exit(1)
	} 

	file, err := os.Open(args[0])
    util.Check(err)
    defer file.Close()

	scanner := bufio.NewScanner(file)
	result := 0
    for scanner.Scan() { 
		result += parseLine(scanner.Text())
    }

	util.Check(scanner.Err())

	fmt.Println(result)
}

func parseLine(line string) int {
	if len(line) == 0 {
		return 0
	}
	parts := strings.Split(line, ":")
	_, data := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])

	result := map[string]int {
		"red": 0,
		"green": 0,
		"blue": 0,
	}

	for _, set := range strings.Split(data, ";") {
		for _, entry := range strings.Split(strings.TrimSpace(set), ",") {
			parts = strings.Split(strings.TrimSpace(entry), " ")
			count, err := strconv.Atoi(strings.TrimSpace(parts[0]))
			util.Check(err)
			color := strings.TrimSpace(parts[1])
			result[color] = util.Max(result[color], count)
		}
	}

    return result["red"] * result["green"] * result["blue"]
}
