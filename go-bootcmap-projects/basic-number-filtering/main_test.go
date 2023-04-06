package main

import (
	"reflect"
	"testing"
)

func TestEvenNumbers(t *testing.T) {
	// given
	numRange := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := []int{2, 4, 6, 8, 10}

	// when
	output := printEven(numRange)

	// then
	if !reflect.DeepEqual(output, want) {
		t.Errorf("wanted: %v but got: %v", want, output)
	}
}
func TestOddNumbers(t *testing.T) {
	// given
	numRange := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := []int{1, 3, 5, 7, 9}

	// when
	output := printOdd(numRange)

	// then
	if !reflect.DeepEqual(output, want) {
		t.Errorf("wanted: %v but got: %v", want, output)
	}
}

func TestFilter(t *testing.T) {

	numRange := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := []int{3, 5, 6, 7, 8, 9, 10}
	greaterThanFive := func(val int) bool { return val >= 5 }
	divideByThree := func(val int) bool { return val%3 == 0 }
	output := filterAny(numRange, greaterThanFive, divideByThree)
	if !reflect.DeepEqual(output, want) {
		t.Errorf("wanted: %v but got: %v", want, output)
	}
}

func TestWithMultipleCondtions(t *testing.T) {
	odd := func(n int) bool { return n%2 != 0 }
	even := func(n int) bool { return !odd(n) }
	greaterThan := func(n int) Condition {
		return func(m int) bool {
			return m > n
		}
	}

	greaterThan5 := greaterThan(5)
	multiplesOf := func(n int) Condition { return func(m int) bool { return m%n == 0 } }
	multiplesOf3 := multiplesOf(3)

	lessThanN := func(n int) Condition { return func(m int) bool { return m < n } }
	lessThan15 := lessThanN(15)

	tt := []struct {
		nums  []int
		conds []Condition
		want  []int
	}{

		{
			nums:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			conds: []Condition{odd, greaterThan5, multiplesOf3},
			want:  []int{9, 15},
		},
		{
			nums:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			conds: []Condition{even, lessThan15, multiplesOf3},
			want:  []int{6, 12},
		},
	}

	for _, opt := range tt {
		output := filter(opt.nums, opt.conds...)
		if !reflect.DeepEqual(output, opt.want) {
			t.Errorf("wanted: %v but got: %v", opt.want, output)
		}
	}
}
