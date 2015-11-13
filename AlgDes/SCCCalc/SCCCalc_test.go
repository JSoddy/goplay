package main

import (
	"testing"
)

type leadersPairs struct {
	graph
	l []leaders
}

type insertPairs struct {
	input  []leaders
	name   int
	output []leaders
}

var (
	v1 = vertex{leader: 1}
	v2 = vertex{leader: 1}
	v3 = vertex{leader: 3}
	v4 = vertex{leader: 2}
	v5 = vertex{leader: 6}

	countLeadersData = []leadersPairs{
		{graph{[]*vertex{}, []*edge{}}, []leaders{}},
		{graph{[]*vertex{&v1}, []*edge{}}, []leaders{{1, 1}}},
		{graph{[]*vertex{&v1, &v2, &v3, &v4, &v5}, []*edge{}},
			[]leaders{{1, 2}, {2, 1}, {3, 1}, {6, 1}}},
	}
	insertLeaderData = []insertPairs{
		{[]leaders{}, 1,
			[]leaders{{1, 1}}},
		{[]leaders{{2, 1}}, 1,
			[]leaders{{1, 1}, {2, 1}}},
		{[]leaders{{1, 1}, {2, 1}}, 1,
			[]leaders{{1, 2}, {2, 1}}},
		{[]leaders{{1, 1}, {2, 1}, {7, 5}}, 2,
			[]leaders{{1, 1}, {2, 2}, {7, 5}}},
		{[]leaders{{1, 1}, {2, 2}, {7, 5}}, 11,
			[]leaders{{1, 1}, {2, 2}, {7, 5}, {11, 1}}},
	}
)

func TestCountLeaders(t *testing.T) {
	for _, pair := range countLeadersData {
		a := countLeaders(&pair.graph)
		if !sameLeaders(a, pair.l) {
			t.Error(
				"For", pair.graph,
				"expected", pair.l,
				"got", a,
			)
		}
	}
}

func TestInsertLeader(t *testing.T) {
	for _, pair := range insertLeaderData {
		a := insertLeader(pair.input, pair.name)
		if !sameLeaders(a, pair.output) {
			t.Error(
				"For", pair.input, pair.name,
				"expected", pair.output,
				"got", a,
			)
		}
	}
}

func sameLeaders(a []leaders, b []leaders) bool {
	if len(a) != len(b) {
		return false
	}
	for i, j := range a {
		if b[i] != j {
			return false
		}
	}
	return true
}
