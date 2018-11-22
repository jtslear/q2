package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	initial       []int
	desiredResult []int
}

var td []Test

func TestTheQuestion(t *testing.T) {
	td = append(td, Test{
		[]int{1, 2, 3, 4, 5},
		[]int{120, 60, 40, 30, 24},
	})

	td = append(td, Test{
		[]int{3, 2, 1},
		[]int{2, 3, 6},
	})

	td = append(td, Test{
		[]int{2, 2, 2},
		[]int{4, 4, 4},
	})

	for _, v := range td {
		result := theQuestion(v.initial)

		assert.Equalf(
			t,
			v.desiredResult,
			result,
			"Expected %v, got %v\n", v.desiredResult, result,
		)
	}
}
