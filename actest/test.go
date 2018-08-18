package actest

import (
    "fmt"
    "github.com/stretchr/testify/assert"
)

type SimpleTest struct {

}

var (
    t = SimpleTest{}
)

func (t SimpleTest) Errorf(format string, args ...interface{}) {
    fmt.Printf(format, args)
}

func Equal(expected, actual interface{}, msgAndArgs ...interface{}) bool {
    return assert.Equal(t, expected, actual, msgAndArgs)
}