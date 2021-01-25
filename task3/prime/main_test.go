package main

import "testing"

func Test_isPrime(t *testing.T) {
	testsTable := []struct {
		get int
		want bool
	}{
		{	
			15,false,
		},
		{
			4,false,
		},
		{
			31,true,
		},
	}
	for _, testCase := range testsTable {
		result := isPrime(testCase.get)
		t.Logf("Calling isPrime(%d), result: %t", testCase.get, result)
		if result != testCase.want {
			t.Errorf("Incorrect result, Expect %t, got %t",
					testCase.want,result)
		}
	}
}
