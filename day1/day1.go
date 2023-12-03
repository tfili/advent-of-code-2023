package adventofcode

import util "adventofcode/util"

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func filter[T any](ss []T, test func(T) bool) (ret []T) {
    for _, s := range ss {
        if test(s) {
            ret = append(ret, s)
        }
    }
    return
}

func Run(args []string) {   
	if len(args) == 0 {
		fmt.Println("A file must be specified")
		os.Exit(1)
	} 

	file, err := os.Open(args[0])
    util.Check(err)
    defer file.Close()

	numberMappings := getMappings();

	scanner := bufio.NewScanner(file)
	result := 0
    for scanner.Scan() {
		line := scanner.Text()
		lineLength := len(line)
		values := make([]int, len(line))
		for key, value := range numberMappings {
			index := 0
			for {
				subIndex := strings.Index(line[(index):], key)
				if subIndex == -1 {
					break
				}
				index += subIndex
				values[index] = value
				index++
				if index >= lineLength {
					break
				}
			}
		}

		filterFunc := func(n int) bool { return n != 0 }
		values = filter(values, filterFunc)
		lenValues := len(values)
		if (lenValues != 0) {
 			result += values[0] * 10 + values[lenValues-1]
		}
    }

	util.Check(scanner.Err())

	fmt.Println(result)
}

func getMappings() map[string]int {
	numberMappings := map[string]int {
		"one"   : 1,
		"1"     : 1,
		"two"   : 2,
		"2"     : 2,
		"three" : 3,
		"3"     : 3,
		"four"  : 4,
		"4"     : 4,
		"five"  : 5,
		"5"     : 5,
		"six"   : 6,
		"6"     : 6,
		"seven" : 7,
		"7"     : 7,
		"eight" : 8,
		"8"     : 8,
		"nine"  : 9,
		"9"     : 9,
	}

	return numberMappings;
}
