package main

import (
	"testing"
)

var testInputOutput = []struct {
	input          map[string]float64
	expectedOutput PairList
}{
	{
		map[string]float64{
			"Bitcoin":  15000,
			"Ripple":   1,
			"Ethereum": 1000,
		}, PairList{
			Pair{"Bitcoin", 15000},
			Pair{"Ethereum", 1000},
			Pair{"Ripple", 1},
		},
	},
}

func testEqualityPairList(a, b PairList) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestSortMapByValue(t *testing.T) {
	for _, value := range testInputOutput {
		actualResult := SortMapByValue(value.input)
		expectedResult := value.expectedOutput
		if !testEqualityPairList(expectedResult, actualResult) {
			t.Fatalf("Expected: %v but got %v\n", expectedResult, actualResult)
		}
	}
}
