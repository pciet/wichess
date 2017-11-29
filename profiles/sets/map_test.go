// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"testing"
)

type MapEqualCase struct {
	Equal bool
	A     MapPathSet
	B     MapPathSet
	C     MapPathSet
	D     MapPathSet
}

var MapEqualCases = []MapEqualCase{
	{
		Equal: false,
		A:     MapPathSet{},
		B: MapPathSet{
			&Path{{0, 0}}: {},
		},
		C: MapPathSet{},
		D: MapPathSet{
			&Path{{2, 2}}: {},
		},
	},
	{
		Equal: true,
		A:     MapPathSet{},
		B:     MapPathSet{},
	},
	{
		Equal: true,
		A: MapPathSet{
			&Path{{0, 0}, {1, 1}, {2, 2}}: {},
			&Path{{1, 2}, {3, 2}}:         {},
		},
		B: MapPathSet{
			&Path{{1, 2}, {3, 2}}:         {},
			&Path{{0, 0}, {1, 1}, {2, 2}}: {},
		},
	},
	{
		Equal: true,
		A: MapPathSet{
			&Path{{1, 5}}:         {},
			&Path{{2, 6}}:         {},
			&Path{{1, 1}, {2, 2}}: {},
		},
		B: MapPathSet{
			&Path{{1, 5}}:         {},
			&Path{{1, 1}, {2, 2}}: {},
			&Path{{2, 6}}:         {},
		},
		C: MapPathSet{
			&Path{{2, 6}}:         {},
			&Path{{1, 5}}:         {},
			&Path{{1, 1}, {2, 2}}: {},
		},
	},
	{
		Equal: false,
		A: MapPathSet{
			&Path{{1, 5}, {2, 6}}:                                 {},
			&Path{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}}: {},
			&Path{{1, 1}}: {},
		},
		B: MapPathSet{
			&Path{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}}: {},
			&Path{{1, 5}, {2, 6}, {3, 7}}:                         {},
			&Path{{1, 1}}:                                         {},
		},
		C: MapPathSet{
			&Path{{1, 5}, {2, 6}}:                                 {},
			&Path{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}}: {},
			&Path{{1, 1}}: {},
		},
	},
}

func TestMapEqual(t *testing.T) {
	for i, c := range MapEqualCases {
		args := make([]MapPathSet, 0, 4)
		if c.A != nil {
			args = append(args, c.A)
		}
		if c.B != nil {
			args = append(args, c.B)
		}
		if c.C != nil {
			args = append(args, c.C)
		}
		if c.D != nil {
			args = append(args, c.D)
		}
		if MapEqual(args...) != c.Equal {
			t.Log(c)
			t.Fatalf("%v failed\n", i)
		}
	}
}

type MapAddCase struct {
	A    MapPathSet
	Item Path
	Out  MapPathSet
}

var MapAddCases = []MapAddCase{
	{
		A:    MapPathSet{},
		Item: Path{{0, 0}},
		Out: MapPathSet{
			&Path{{0, 0}}: {},
		},
	},
	{
		A: MapPathSet{
			&Path{{1, 1}}: {},
		},
		Item: Path{{2, 2}},
		Out: MapPathSet{
			&Path{{2, 2}}: {},
			&Path{{1, 1}}: {},
		},
	},
	{
		A: MapPathSet{
			&Path{{0, 0}, {0, 1}}: {},
			&Path{{0, 2}, {1, 1}}: {},
		},
		Item: Path{{0, 0}, {0, 3}, {2, 2}},
		Out: MapPathSet{
			&Path{{0, 0}, {0, 3}, {2, 2}}: {},
			&Path{{0, 0}, {0, 1}}:         {},
			&Path{{0, 2}, {1, 1}}:         {},
		},
	},
}

func TestMapAdd(t *testing.T) {
	for i, c := range MapAddCases {
		if MapEqual(MapAdd(c.A, c.Item), c.Out) == false {
			t.Log(c)
			t.Fatalf("%v: out not equal\n", i)
		}
	}
}

type MapDeleteCase struct {
	Set  MapPathSet
	Item Path
	Out  MapPathSet
}

var MapDeleteCases = []MapDeleteCase{
	{
		Set: MapPathSet{
			&Path{{0, 0}}:         {},
			&Path{{0, 1}, {1, 1}}: {},
		},
		Item: Path{{0, 0}},
		Out: MapPathSet{
			&Path{{0, 1}, {1, 1}}: {},
		},
	},
	{
		Set: MapPathSet{
			&Path{{5, 5}}: {},
		},
		Item: Path{{5, 5}},
		Out:  MapPathSet{},
	},
	{
		Set: MapPathSet{
			&Path{{5, 5}, {4, 5}, {3, 3}}: {},
			&Path{{3, 3}, {2, 2}}:         {},
			&Path{{4, 4}}:                 {},
		},
		Item: Path{{5, 5}, {4, 5}, {3, 3}},
		Out: MapPathSet{
			&Path{{3, 3}, {2, 2}}: {},
			&Path{{4, 4}}:         {},
		},
	},
}

func TestMapDelete(t *testing.T) {
	for i, c := range MapDeleteCases {
		if MapEqual(c.Out, MapDelete(c.Set, c.Item)) == false {
			t.Log(c)
			t.Fatalf("%v: out not equal\n", i)
		}
	}
}

type MapCombineCase struct {
	A   MapPathSet
	B   MapPathSet
	C   MapPathSet
	D   MapPathSet
	Out MapPathSet
}

var MapCombineCases = []MapCombineCase{
	{
		A: MapPathSet{
			&Path{{1, 1}, {2, 2}}: {},
		},
		B: MapPathSet{
			&Path{{1, 1}}: {},
		},
		Out: MapPathSet{
			&Path{{1, 1}, {2, 2}}: {},
			&Path{{1, 1}}:         {},
		},
	},
	{
		A: MapPathSet{
			&Path{{1, 1}, {2, 2}}: {},
			&Path{{1, 1}}:         {},
		},
		B: MapPathSet{
			&Path{{1, 1}}: {},
		},
		Out: MapPathSet{
			&Path{{1, 1}, {2, 2}}: {},
			&Path{{1, 1}}:         {},
			&Path{{1, 1}}:         {},
		},
	},
	{
		A: MapPathSet{
			&Path{{1, 1}, {2, 2}}: {},
		},
		B: MapPathSet{
			&Path{{1, 1}}: {},
		},
		C: MapPathSet{
			&Path{{3, 3}, {4, 4}}: {},
		},
		Out: MapPathSet{
			&Path{{3, 3}, {4, 4}}: {},
			&Path{{1, 1}, {2, 2}}: {},
			&Path{{1, 1}}:         {},
		},
	},
	{
		A: MapPathSet{
			&Path{{0, 2}, {5, 3}}: {},
		},
		B: MapPathSet{
			&Path{{1, 7}, {2, 6}}: {},
		},
		C: MapPathSet{
			&Path{{1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}}: {},
		},
		D: MapPathSet{
			&Path{{3, 3}, {4, 4}}: {},
		},
		Out: MapPathSet{
			&Path{{3, 3}, {4, 4}}:                         {},
			&Path{{1, 7}, {2, 6}}:                         {},
			&Path{{0, 2}, {5, 3}}:                         {},
			&Path{{1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}}: {},
		},
	},
}

func TestMapCombine(t *testing.T) {
	for i, c := range MapCombineCases {
		args := make([]MapPathSet, 0, 4)
		if c.A != nil {
			args = append(args, c.A)
		}
		if c.B != nil {
			args = append(args, c.B)
		}
		if c.C != nil {
			args = append(args, c.C)
		}
		if c.D != nil {
			args = append(args, c.D)
		}
		if MapEqual(c.Out, MapCombine(args...)) == false {
			t.Log(c)
			t.Fatalf("%v: out not equal\n", i)
		}
	}
}

type MapReduceCase struct {
	The MapPathSet
	Out MapPathSet
}

var MapReduceCases = []MapReduceCase{
	{
		The: MapPathSet{
			&Path{{0, 0}}: {},
			&Path{{0, 0}}: {},
		},
		Out: MapPathSet{
			&Path{{0, 0}}: {},
		},
	},
	{
		The: MapPathSet{
			&Path{{0, 0}}: {},
		},
		Out: MapPathSet{
			&Path{{0, 0}}: {},
		},
	},
	{
		The: MapPathSet{
			&Path{{1, 1}, {2, 2}}: {},
			&Path{{3, 3}, {1, 1}}: {},
			&Path{{1, 1}, {1, 1}}: {},
			&Path{{3, 3}, {1, 1}}: {},
			&Path{{1, 1}, {2, 2}}: {},
		},
		Out: MapPathSet{
			&Path{{1, 1}, {1, 1}}: {},
			&Path{{1, 1}, {2, 2}}: {},
			&Path{{3, 3}, {1, 1}}: {},
		},
	},
}

func TestMapReduce(t *testing.T) {
	for i, c := range MapReduceCases {
		if MapEqual(c.Out, MapReduce(c.The)) == false {
			t.Log(MapReduce(c.The))
			t.Log(c)
			t.Fatalf("%v: out not equal\n", i)
		}
	}
}

type MapHasCase struct {
	Has  bool
	Set  MapPathSet
	Item Path
}

var MapHasCases = []MapHasCase{
	{
		Has: true,
		Set: MapPathSet{
			&Path{{0, 0}, {1, 1}, {2, 2}}: {},
			&Path{{0, 0}}:                 {},
		},
		Item: Path{{0, 0}, {1, 1}, {2, 2}},
	},
	{
		Has: false,
		Set: MapPathSet{
			&Path{{0, 0}, {1, 1}, {2, 2}}: {},
			&Path{{0, 0}}:                 {},
		},
		Item: Path{{0, 0}, {2, 2}, {1, 1}},
	},
}

func TestMapHas(t *testing.T) {
	for i, c := range MapHasCases {
		if MapHas(c.Set, c.Item) != c.Has {
			t.Log(c)
			t.Fatalf("%v: MapHas not equal to Has\n", i)
		}
	}
}

type MapDiffCase struct {
	A   MapPathSet
	B   MapPathSet
	Out MapPathSet
}

var MapDiffCases = []MapDiffCase{
	{
		A: MapPathSet{
			&Path{{0, 0}, {1, 1}, {2, 2}, {3, 3}}: {},
			&Path{{0, 0}, {1, 5}}:                 {},
			&Path{{0, 0}, {1, 1}}:                 {},
		},
		B: MapPathSet{
			&Path{{0, 0}, {1, 1}}: {},
		},
		Out: MapPathSet{
			&Path{{0, 0}, {1, 1}, {2, 2}, {3, 3}}: {},
			&Path{{0, 0}, {1, 5}}:                 {},
		},
	},
	{
		A: MapPathSet{
			&Path{{0, 0}}: {},
		},
		B: MapPathSet{
			&Path{{0, 0}}: {},
		},
		Out: MapPathSet{},
	},
	{
		A: MapPathSet{
			&Path{{5, 5}, {4, 4}, {3, 3}}: {},
			&Path{{3, 3}, {4, 4}, {5, 5}}: {},
		},
		B: MapPathSet{
			&Path{{3, 3}, {5, 5}, {4, 4}}: {},
		},
		Out: MapPathSet{
			&Path{{5, 5}, {4, 4}, {3, 3}}: {},
			&Path{{3, 3}, {4, 4}, {5, 5}}: {},
			&Path{{3, 3}, {5, 5}, {4, 4}}: {},
		},
	},
}

func TestMapDiff(t *testing.T) {
	for i, c := range MapDiffCases {
		if MapEqual(c.Out, MapDiff(c.A, c.B)) == false {
			t.Log(c)
			t.Fatalf("%v: out doesn't match\n", i)
		}
	}
}
