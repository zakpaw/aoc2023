package main

import (
	"bufio"
	"fmt"
	"os"
    "strings"
	"strconv"
	"unicode"
)

func decipher(value string, replaceTextDigits bool) (int, error) {
    if replaceTextDigits {
        digitsMap := strings.NewReplacer(
            "zero", "zero0zero",
            "one", "one1one",
            "two", "two2two",
            "three", "three3three",
            "four", "four4four",
            "five", "five5five",
            "six", "six6six",
            "seven", "seven7seven",
            "eight", "eight8eight",
            "nine", "nine9nine",
        )
        value = digitsMap.Replace(value)
    }

    firstDigitIndex := strings.IndexFunc(value, unicode.IsDigit)
    if firstDigitIndex == -1 {
        return 0, fmt.Errorf("no digits found in the input")
    }
    firstDigit := string(value[firstDigitIndex])
    lastDigit := value[strings.LastIndexFunc(value, unicode.IsDigit)]
    return strconv.Atoi(firstDigit + string(lastDigit))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var resultP1, resultP2 int
	for scanner.Scan() {
		line := scanner.Text()
		numP1, _ := decipher(line, false)
		numP2, _ := decipher(line, true)
		resultP1 += numP1
		resultP2 += numP2
	}
    fmt.Printf("part1: %d\npart2: %d\n", resultP1, resultP2)
}
