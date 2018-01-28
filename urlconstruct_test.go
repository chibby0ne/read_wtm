package main

import (
    "bytes"
    "testing"
)

func TestWriteOneParameterQueryToReturnNoAdaptTrueStringWithZeroGPUsGiven(t *testing.T) {
    var buffer bytes.Buffer
    writeOneParameterQuery(&buffer, "adapt_q_280x=", "0")
    actualResult := buffer.String()
    expectedResult := "adapt_q_280x=0&"
    if expectedResult != actualResult {
        t.Fatalf("Expected: %s but got: %s\n", expectedResult, actualResult)
    }
}

func TestWriteOneParameterQueryToReturnAdaptTrueStringWithNonZeroGPUsGiven(t *testing.T) {
    var buffer bytes.Buffer
    writeOneParameterQuery(&buffer, "adapt_q_280x=", "1")
    actualResult := buffer.String()
    expectedResult := "adapt_q_280x=1&adapt_280x=true&"
    if expectedResult != actualResult {
        t.Fatalf("Expected: %s but got: %s\n", expectedResult, actualResult)
    }
}
