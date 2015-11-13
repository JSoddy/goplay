package math

import "testing"

type testpair struct {
	values  []float64
	average float64
	min     float64
	max     float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5, 1, 2},
	{[]float64{1, 1, 1, 1, 1, 1}, 1, 1, 1},
	{[]float64{-1, 1}, 0, -1, 1},
	{[]float64{}, 0, 0, 0},
	{[]float64{1.111, 8, 77, 0, -5}, 16.2222, -5, 77},
}

func TestAverage(t *testing.T) {
	for _, pair := range tests {
		a := Average(pair.values)
		if a != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", a,
			)
		}
	}
}

func TestMin(t *testing.T) {
	for _, pair := range tests {
		v := Min(pair.values)
		if v != pair.min {
			t.Error(
				"For", pair.values,
				"expected", pair.min,
				"got", v,
			)
		}
	}
}

func TestMax(t *testing.T) {
	for _, pair := range tests {
		v := Max(pair.values)
		if v != pair.max {
			t.Error(
				"For", pair.values,
				"expected", pair.max,
				"got", v,
			)
		}
	}
}
