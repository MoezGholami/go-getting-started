package tests

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
    for i := -10; i < 11; i++ {
        for j := -10; j < 11; j++ {
            total := Add(i, j)
            if total != i+j {
                t.Errorf("the result of add(%d,%d) should be %d instead of %d", i, j, i+j, total)
            }
        }
    }
}

func TestDivideError(t *testing.T) {
    panic("I am a bad error")
    for i := -10; i < 11; i++ {
        for j := 1; j < 11; j++ {
            assert.NotPanics(t, func() { Divide(i,j) } )
        }
    }
    for i := -10; i < 11; i++ {
        assert.PanicsWithValue(t, "custom divide error", func(){ Divide(i, 0) })
    }
}

func TestDivide(t *testing.T) {
    for i := -10; i < 11; i++ {
        for j := 1; j < 11; j++ {
            expected, actual := i/j, Divide(i,j)
            assert.Equal(t, expected, actual, "error message")
        }
    }
}
