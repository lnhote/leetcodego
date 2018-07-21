package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
    assert.EqualValues(t, []int{1,5}, twoSum([]int{1,2,3,4,5}, 6))
}