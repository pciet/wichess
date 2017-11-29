// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"testing"
)

type SliceEqualCase struct {
	Equal bool
	A     SlicePathSet
	B     SlicePathSet
	C     SlicePathSet
	D     SlicePathSet
}

var SliceEqualCases = []SliceEqualCase{
	{
		Equal: false,
		A:     SlicePathSet{},
		B: SlicePathSet{
			{{0, 0}},
		},
		C: SlicePathSet{},
		D: SlicePathSet{
			{{2, 2}},
		},
	},
	{
		Equal: true,
		A:     SlicePathSet{},
		B:     SlicePathSet{},
	},
	{
		Equal: true,
		A: SlicePathSet{
			{{0, 0}, {1, 1}, {2, 2}},
			{{1, 2}, {3, 2}},
		},
		B: SlicePathSet{
			{{1, 2}, {3, 2}},
			{{0, 0}, {1, 1}, {2, 2}},
		},
	},
	{
		Equal: true,
		A: SlicePathSet{
			{{1, 5}},
			{{2, 6}},
			{{1, 1}, {2, 2}},
		},
		B: SlicePathSet{
			{{1, 5}},
			{{1, 1}, {2, 2}},
			{{2, 6}},
		},
		C: SlicePathSet{
			{{2, 6}},
			{{1, 5}},
			{{1, 1}, {2, 2}},
		},
	},
	{
		Equal: false,
		A: SlicePathSet{
			{{1, 5}, {2, 6}},
			{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}},
			{{1, 1}},
		},
		B: SlicePathSet{
			{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}},
			{{1, 5}, {2, 6}, {3, 7}},
			{{1, 1}},
		},
		C: SlicePathSet{
			{{1, 5}, {2, 6}},
			{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}},
			{{1, 1}},
		},
	},
}

func TestSliceEqual(t *testing.T) {
	for i, c := range SliceEqualCases {
		args := make([]SlicePathSet, 0, 4)
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
		if SliceEqual(args...) != c.Equal {
			t.Log(c)
			t.Fatalf("%v failed\n", i)
		}
	}
}

type SliceAddCase struct {
	A    SlicePathSet
	Item Path
	Out  SlicePathSet
}

var SliceAddCases = []SliceAddCase{
	{
		A:    SlicePathSet{},
		Item: Path{{0, 0}},
		Out: SlicePathSet{
			{{0, 0}},
		},
	},
	{
		A: SlicePathSet{
			{{1, 1}},
		},
		Item: Path{{2, 2}},
		Out: SlicePathSet{
			{{2, 2}},
			{{1, 1}},
		},
	},
	{
		A: SlicePathSet{
			{{0, 0}, {0, 1}},
			{{0, 2}, {1, 1}},
		},
		Item: Path{{0, 0}, {0, 3}, {2, 2}},
		Out: SlicePathSet{
			{{0, 0}, {0, 3}, {2, 2}},
			{{0, 0}, {0, 1}},
			{{0, 2}, {1, 1}},
		},
	},
}

func TestSliceAdd(t *testing.T) {
	for i, c := range SliceAddCases {
		if SliceEqual(SliceAdd(c.A, c.Item), c.Out) == false {
			t.Log(c)
			t.Fatalf("%v: out not equal\n", i)
		}
	}
}

type SliceDeleteCase struct {
	Set  SlicePathSet
	Item Path
	Out  SlicePathSet
}

var SliceDeleteCases = []SliceDeleteCase{
	{
		Set: SlicePathSet{
			{{0, 0}},
			{{0, 1}, {1, 1}},
		},
		Item: Path{{0, 0}},
		Out: SlicePathSet{
			{{0, 1}, {1, 1}},
		},
	},
	{
		Set: SlicePathSet{
			{{5, 5}},
		},
		Item: Path{{5, 5}},
		Out:  SlicePathSet{},
	},
	{
		Set: SlicePathSet{
			{{5, 5}, {4, 5}, {3, 3}},
			{{3, 3}, {2, 2}},
			{{4, 4}},
		},
		Item: Path{{5, 5}, {4, 5}, {3, 3}},
		Out: SlicePathSet{
			{{3, 3}, {2, 2}},
			{{4, 4}},
		},
	},
}

func TestSliceDelete(t *testing.T) {
	for i, c := range SliceDeleteCases {
		if SliceEqual(c.Out, SliceDelete(c.Set, c.Item)) == false {
			t.Log(c)
			t.Fatalf("%v: out not equal\n", i)
		}
	}
}

type SliceCombineCase struct {
	A   SlicePathSet
	B   SlicePathSet
	C   SlicePathSet
	D   SlicePathSet
	Out SlicePathSet
}

var SliceCombineCases = []SliceCombineCase{
	{
		A: SlicePathSet{
			{{1, 1}, {2, 2}},
		},
		B: SlicePathSet{
			{{1, 1}},
		},
		Out: SlicePathSet{
			{{1, 1}, {2, 2}},
			{{1, 1}},
		},
	},
	{
		A: SlicePathSet{
			{{1, 1}, {2, 2}},
			{{1, 1}},
		},
		B: SlicePathSet{
			{{1, 1}},
		},
		Out: SlicePathSet{
			{{1, 1}, {2, 2}},
			{{1, 1}},
			{{1, 1}},
		},
	},
	{
		A: SlicePathSet{
			{{1, 1}, {2, 2}},
		},
		B: SlicePathSet{
			{{1, 1}},
		},
		C: SlicePathSet{
			{{3, 3}, {4, 4}},
		},
		Out: SlicePathSet{
			{{3, 3}, {4, 4}},
			{{1, 1}, {2, 2}},
			{{1, 1}},
		},
	},
	{
		A: SlicePathSet{
			{{0, 2}, {5, 3}},
		},
		B: SlicePathSet{
			{{1, 7}, {2, 6}},
		},
		C: SlicePathSet{
			{{1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}},
		},
		D: SlicePathSet{
			{{3, 3}, {4, 4}},
		},
		Out: SlicePathSet{
			{{3, 3}, {4, 4}},
			{{1, 7}, {2, 6}},
			{{0, 2}, {5, 3}},
			{{1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}},
		},
	},
}

func TestSliceCombine(t *testing.T) {
	for i, c := range SliceCombineCases {
		args := make([]SlicePathSet, 0, 4)
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
		if SliceEqual(c.Out, SliceCombine(args...)) == false {
			t.Log(c)
			t.Fatalf("%v: out not equal\n", i)
		}
	}
}

type SliceReduceCase struct {
	The SlicePathSet
	Out SlicePathSet
}

var SliceReduceCases = []SliceReduceCase{
	{
		The: SlicePathSet{
			{{0, 0}},
			{{0, 0}},
		},
		Out: SlicePathSet{
			{{0, 0}},
		},
	},
	{
		The: SlicePathSet{
			{{0, 0}},
		},
		Out: SlicePathSet{
			{{0, 0}},
		},
	},
	{
		The: SlicePathSet{
			{{1, 1}, {2, 2}},
			{{3, 3}, {1, 1}},
			{{1, 1}, {1, 1}},
			{{3, 3}, {1, 1}},
			{{1, 1}, {2, 2}},
		},
		Out: SlicePathSet{
			{{1, 1}, {1, 1}},
			{{1, 1}, {2, 2}},
			{{3, 3}, {1, 1}},
		},
	},
}

func TestSliceReduce(t *testing.T) {
	for i, c := range SliceReduceCases {
		if SliceEqual(c.Out, SliceReduce(c.The)) == false {
			t.Log(SliceReduce(c.The))
			t.Log(c)
			t.Fatalf("%v: out not equal\n", i)
		}
	}
}

type SliceHasCase struct {
	Has  bool
	Set  SlicePathSet
	Item Path
}

var SliceHasCases = []SliceHasCase{
	{
		Has: true,
		Set: SlicePathSet{
			{{0, 0}, {1, 1}, {2, 2}},
			{{0, 0}},
		},
		Item: Path{{0, 0}, {1, 1}, {2, 2}},
	},
	{
		Has: false,
		Set: SlicePathSet{
			{{0, 0}, {1, 1}, {2, 2}},
			{{0, 0}},
		},
		Item: Path{{0, 0}, {2, 2}, {1, 1}},
	},
}

func TestSliceHas(t *testing.T) {
	for i, c := range SliceHasCases {
		if SliceHas(c.Set, c.Item) != c.Has {
			t.Log(c)
			t.Fatalf("%v: SliceHas not equal to Has\n", i)
		}
	}
}

type SliceDiffCase struct {
	A   SlicePathSet
	B   SlicePathSet
	Out SlicePathSet
}

var SliceDiffCases = []SliceDiffCase{
	{
		A: SlicePathSet{
			{{0, 0}, {1, 1}, {2, 2}, {3, 3}},
			{{0, 0}, {1, 5}},
			{{0, 0}, {1, 1}},
		},
		B: SlicePathSet{
			{{0, 0}, {1, 1}},
		},
		Out: SlicePathSet{
			{{0, 0}, {1, 1}, {2, 2}, {3, 3}},
			{{0, 0}, {1, 5}},
		},
	},
	{
		A: SlicePathSet{
			{{0, 0}},
		},
		B: SlicePathSet{
			{{0, 0}},
		},
		Out: SlicePathSet{},
	},
	{
		A: SlicePathSet{
			{{5, 5}, {4, 4}, {3, 3}},
			{{3, 3}, {4, 4}, {5, 5}},
		},
		B: SlicePathSet{
			{{3, 3}, {5, 5}, {4, 4}},
		},
		Out: SlicePathSet{
			{{5, 5}, {4, 4}, {3, 3}},
			{{3, 3}, {4, 4}, {5, 5}},
			{{3, 3}, {5, 5}, {4, 4}},
		},
	},
}

func TestSliceDiff(t *testing.T) {
	for i, c := range SliceDiffCases {
		if SliceEqual(c.Out, SliceDiff(c.A, c.B)) == false {
			t.Log(c)
			t.Fatalf("%v: out doesn't match\n", i)
		}
	}
}
