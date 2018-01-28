package main

import (
    "bytes"
    "testing"
)

var writeOneParameterQueryTestInputOutput = []struct {
    adaptString string
    numOfGPUs string
    expectedOutput string
} {
    {"adapt_q_280x=", "0",  "adapt_q_280x=0&"},
    {"adapt_q_280x=", "1", "adapt_q_280x=1&adapt_280x=true&"},
}

func TestWriteOneParameterQuery(t *testing.T) {
    for _, vals := range writeOneParameterQueryTestInputOutput {
        var buffer bytes.Buffer
        writeOneParameterQuery(&buffer, vals.adaptString, vals.numOfGPUs)
        actualResult := buffer.String()
        expectedResult := vals.expectedOutput
        if expectedResult != actualResult {
            t.Fatalf("Expected: %s but got: %s\n", expectedResult, actualResult)
        }
    }
}
