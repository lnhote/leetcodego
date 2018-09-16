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
    fmt.Printf(format, args...)
}

func Equal(expected, actual interface{}, msgAndArgs ...interface{}) bool {
    return assert.Equal(t, expected, actual, msgAndArgs...)
}

func True(actual bool, msgAndArgs ...interface{}) bool {
    return assert.True(t, actual, msgAndArgs...)
}

func False(actual bool, msgAndArgs ...interface{}) bool {
    return assert.False(t, actual, msgAndArgs...)
}

func EqualValues(expected, actual interface{}, msgAndArgs ...interface{}) bool {
    return assert.EqualValues(t, expected, actual, msgAndArgs...)
}