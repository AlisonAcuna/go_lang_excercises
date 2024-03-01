package main

import "testing"

func TestStringChop(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput string
	}{
		{"Hello World", "ello Worl"},
		{"Heffalumps", "effalump"},
		{"Woosels", "oosel"},
	}

	for _, test := range tests {
		actualOutput := StringChop(test.input)
		if actualOutput != test.expectedOutput {
			t.Errorf("StringChop(%s) = %s, want %s", test.input, actualOutput, test.expectedOutput)
		}
	}
}

func TestCheckLength(t *testing.T) {
	tests := []struct {
		input          []int
		expectedOutput bool
	}{
		{[]int{5, 5, 5, 5, 5}, true},
		{[]int{1, 9, 5, 8, 2}, true},
		{[]int{-6, 15, 4}, true},
		{[]int{1, 2, 3, 4}, false},
	}

	for _, test := range tests {
		actualOutput := CheckLength(test.input)
		if actualOutput != test.expectedOutput {
			t.Errorf("CheckLength(%q) = %t, want %t", test.input, actualOutput, test.expectedOutput)
		}
	}
}

func TestMultiDivision(t *testing.T) {
	tests := []struct {
		testId         string
		inputA         int
		inputB         int
		inputC         int
		expectedOutput bool
	}{
		{"1", 1, 2, 3, false},
		{"2", 69, 15, 9, false},
		{"3", 9, 2, 52, false},
		{"4", 5, 2, 3, false},
		{"5", 5, 2, 1, true},
		{"6", 261, 2, 1, true},
		{"7", 22, 2, 22, true},
		{"8", 69, 12, 3, true},
	}

	for _, test := range tests {
		actualOutput := MultiDivision(test.inputA, test.inputB, test.inputC)
		if actualOutput != test.expectedOutput {
			t.Errorf("MultiDivision test %q = %t, want %t", test.testId, actualOutput, test.expectedOutput)
		}
	}
}

func TestCreateMultiples(t *testing.T) {
	tests := []struct {
		testId         string
		inputNum       int
		inputLenght    int
		expectedOutput []int
	}{
		{"1", 7, 5, []int{7, 14, 21, 28, 35}},
		{"2", 12, 10, []int{12, 24, 36, 48, 60, 72, 84, 96, 108, 120}},
		{"3", 17, 7, []int{17, 34, 51, 68, 85, 102, 119}},
		{"4", 630, 14, []int{630, 1260, 1890, 2520, 3150, 3780, 4410, 5040, 5670, 6300, 6930, 7560, 8190, 8820}},
		{"5", 140, 3, []int{140, 280, 420}},
		{"6", 7, 8, []int{7, 14, 21, 28, 35, 42, 49, 56}},
		{"7", 11, 21, []int{11, 22, 33, 44, 55, 66, 77, 88, 99, 110, 121, 132, 143, 154, 165, 176, 187, 198, 209, 220, 231}},
	}
	for _, test := range tests {
		actualOutput := CreateMultiples(test.inputNum, test.inputLenght)

		for i := 0; i < len(actualOutput); i++ {
			if actualOutput[i] != test.expectedOutput[i] {
				t.Errorf("CreateMultiples test %q = %d, want %d", test.testId, actualOutput[i], test.expectedOutput[i])
			}
		}
		if len(actualOutput) != len(test.expectedOutput) {
			t.Errorf("CreateMultiples test %q = %d, want %d", test.testId, len(actualOutput), len(test.expectedOutput))
		}
	}
}
