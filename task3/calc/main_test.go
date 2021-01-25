package main

import (
	"math"
	"testing"
)

func Test_add (t *testing.T) {
	testsTable := []struct {
		gotA float64
		gotB float64
		want float64
	}{
		{
			5.0, 6.0, 11.0,
		},
		{
			5.0, 4.5, 9.5,
		},
		{
			5.5, 6.6, 12.1,
		},
		{
			7.0, 0.0, 7.0,
		},
	}
	for _, testCase := range testsTable {
		result := add(testCase.gotA,testCase.gotB)
		t.Logf ("The result is %.1f, and want: %.1f", result, testCase.want)
		if result != testCase.want {
			t.Errorf("The error occurs, got: %.1f, want: %.1f", result, testCase.want)
		}
	}
}

func Test_sub (t *testing.T) {
	testsTable := []struct {
		gotA float64
		gotB float64
		want float64
	}{
		{
			5.0, 6.0, -1.0,
		},
		{
			5.0, 4.5, 0.5,
		},
		{
			5.5, 6.0, -0.5,
		},
		{
			7.0, 0.0, 7.0,
		},
	}
	for _, testCase := range testsTable {
		result := sub(testCase.gotA,testCase.gotB)
		t.Logf ("The result is %.1f, and want: %.1f", result, testCase.want)
		if result != testCase.want {
			t.Errorf("The error occurs, got: %.1f, want: %.1f", result, testCase.want)
		}
	}
}

func Test_mul (t *testing.T) {
	testsTable := []struct {
		gotA float64
		gotB float64
		want float64
	}{
		{
			5.0, -6.0, -30.0,
		},
		{
			5.0, 4.5, 22.5,
		},
		{
			5.5, 6.0, 33.0,
		},
		{
			7.0, 0.0, 0.0,
		},
	}
	for _, testCase := range testsTable {
		result := mul(testCase.gotA,testCase.gotB)
		t.Logf ("The result is %.1f, and want: %.1f", result, testCase.want)
		if result != testCase.want {
			t.Errorf("The error occurs, got: %.1f, want: %.1f", result, testCase.want)
		}
	}
}

func Test_div (t *testing.T) {
	testsTable := []struct {
		gotA float64
		gotB float64
		want float64
	}{
		{
			5.0, 7.0, 0.7,
		},
		{
			5.0, 0.5, 10.0,
		},
		{
			5.5, 4.0, 1.4,
		},
		{
			7.0, 0.0, math.Inf(0),
		},
	}
	for _, testCase := range testsTable {
		result := div(testCase.gotA,testCase.gotB)
		t.Logf ("The result is %.1f, and want: %.1f", result, testCase.want)
		if math.Floor(result * 100/100) != math.Floor(testCase.want*100/100) {
			t.Errorf("The error occurs, got: %.1f, want: %.1f", result, testCase.want)
		}
	}
}