package main

import "testing"

func TestGreet(t *testing.T){
	type testcase struct {
		input int
		want string
	}

	testcases := []testcase{
		{input: 0, want: "โย่วๆๆๆ"},
		{input: 1, want: "Hello SIR"},
		{input: 2, want: "Hello World - you lazy ass default"},

	}
	for _, tc := range testcases {
		got := greet(Language(tc.input))
		if got != tc.want {
			t.Errorf("want %s got %s", got, tc.want)
		}
	}
}
