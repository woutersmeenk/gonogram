package gonogram

import (
	"reflect"
	"testing"
)

func TestMostLeftSolution(t *testing.T) {
	var tests = []struct {
		clues    []clue
		lineStr  string
		expected []int
	}{
		{[]clue{1, 2}, "....", []int{0, 2}},
		{[]clue{1, 2}, ".-..", []int{0, 2}},
		{[]clue{1, 2}, "..-..", []int{0, 3}},
		{[]clue{1, 2}, "-....", []int{1, 3}},
		{[]clue{1, 2}, "+...", []int{0, 2}},
		{[]clue{1, 2}, ".+...", []int{1, 3}},
		{[]clue{1, 2}, ".+.-.", []int{1, -1}},
	}

	for _, test := range tests {
		line := parseLine(test.lineStr)
		actual := line.mostLeftSolution(test.clues)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("clues: %v line: %v expected: %v actual: %v", test.clues, test.lineStr, test.expected, actual)
		}
	}
}

func TestMostRightSolution(t *testing.T) {
	var tests = []struct {
		clues    []clue
		lineStr  string
		expected []int
	}{
		{[]clue{1, 2}, "....", []int{0, 2}},
		{[]clue{1, 2}, ".-..", []int{0, 2}},
		{[]clue{1, 2}, "..-..", []int{1, 3}},
		{[]clue{1, 2}, ".-...", []int{0, 3}},
		{[]clue{1, 2}, "+...", []int{0, 2}},
		{[]clue{1, 2}, ".+...", []int{1, 3}},
		{[]clue{1, 2}, ".+.-.", []int{1, -1}},
	}

	for _, test := range tests {
		line := parseLine(test.lineStr)
		actual := line.mostRightSolution(test.clues)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("clues: %v line: %v expected: %v actual: %v", test.clues, test.lineStr, test.expected, actual)
		}
	}
}

func parseLine(lineStr string) (line line) {
	line = make([]cell, len(lineStr))
	for i, c := range lineStr {
		switch c {
		case '.':
			line[i] = unknown
		case '-':
			line[i] = white
		case '+':
			line[i] = black
		default:
			panic("Unknown char")
		}
	}
	return line
}
