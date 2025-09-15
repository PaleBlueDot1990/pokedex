package pokecli

import "testing"

type testCase struct {
	input    string 
	expected []string
}

func TestCleanInput(t *testing.T) {
	testCases := []testCase{
		{
			input:     "foo bar",
			expected: []string{"foo", "bar"},
		},
		{
			input:     "   spaced   out   ",
			expected: []string{"spaced", "out"},
		},
		{
			input:     "one   two    three",
			expected: []string{"one", "two", "three"},
		},
		{
			input:     "    leading",
			expected: []string{"leading"},
		},
		{
			input:     "trailing    ",
			expected: []string{"trailing"},
		},
		{
			input:     "   surrounded   ",
			expected: []string{"surrounded"},
		},
		{
			input:     "multiple     spaces     between",
			expected: []string{"multiple", "spaces", "between"},
		},
		{
			input:     "singleword",
			expected: []string{"singleword"},
		},
		{
			input:     "    ",
			expected: []string{},
		},
	}

	for _, c := range testCases {
		actual := CleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Test case failed- Input: %s, Actual: %v, Expected: %v\n", c.input, actual, c.expected)
			t.Fail()
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i] 
			if word != expectedWord {
				t.Errorf("Test case failed- Input: %s, Actual: %v, Expected: %v\n", c.input, actual, c.expected)
				t.Fail()
			}
		}
	}
}