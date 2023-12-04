package main

import (
	"testing"
)

func TestDecipher(t *testing.T) {
    type testCase struct {
        input       string
        expectedP1  int
        expectedP2  int
    }

    testCases := []testCase{
        {"1abc2", 12, 12},
        {"pqr3stu8vwx", 38, 38},
        {"two1nine", 11, 29},
        {"eightwothree", 0, 83},
        {"abcone2threexyz", 22, 13},
        {"xtwone3four", 33, 24},
        {"4nineeightseven2", 42, 42},
        {"zoneight234", 24, 14},
        {"7pqrstsixteen", 77, 76},
    }

    for _, tc := range testCases {
        t.Run(tc.input, func(t *testing.T) {
            gotP1, _ := decipher(tc.input, false)
            if gotP1 != tc.expectedP1 {
                t.Errorf("decipher() with P1, got %v, want %v", gotP1, tc.expectedP1)
            }
            gotP2, _ := decipher(tc.input, true)
            if gotP2 != tc.expectedP2 {
                t.Errorf("decipher() with P2, got %v, want %v", gotP2, tc.expectedP2)
            }
        })
    }
}
