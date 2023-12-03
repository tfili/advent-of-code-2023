package adventofcode

import util "adventofcode/util"

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	re := regexp.MustCompile(`\D`)
	scanner := bufio.NewScanner(file)
	result := 0
    for scanner.Scan() {
		numbersArray := strings.Split(re.ReplaceAllString(scanner.Text(), ""), "")
		numbersLen := len(numbersArray)
		if numbersLen > 0 {
			n, err := strconv.Atoi(numbersArray[0] + numbersArray[numbersLen-1])
			util.Check(err)
			result += n
		}
    }

	util.Check(scanner.Err())

	fmt.Println(result)
}
