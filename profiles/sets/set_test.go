// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"testing"
)

// following patterns described at https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go

var (
	BaseChessSliceSet = SlicePathSet{
		{{0, 1}, {0, 2}},
		{{1, 1}, {1, 2}},
		{{2, 1}, {2, 2}},
		{{3, 1}, {3, 2}},
		{{4, 1}, {4, 2}},
		{{5, 1}, {5, 2}},
		{{6, 1}, {6, 2}},
		{{7, 1}, {7, 2}},
		{{1, 1}, {1, 2}, {0, 2}},
		{{1, 1}, {1, 2}, {2, 2}},
		{{6, 1}, {6, 2}, {5, 2}},
		{{6, 1}, {6, 2}, {7, 2}},
		{{0, 5}, {0, 4}},
		{{1, 5}, {1, 4}},
		{{2, 5}, {2, 4}},
		{{3, 5}, {3, 4}},
		{{4, 5}, {4, 4}},
		{{5, 5}, {5, 4}},
		{{6, 5}, {6, 4}},
		{{7, 5}, {7, 4}},
		{{1, 6}, {1, 5}, {0, 5}},
		{{1, 6}, {1, 5}, {2, 5}},
		{{6, 6}, {6, 5}, {5, 5}},
		{{6, 6}, {6, 5}, {7, 5}},
	}
	BaseChessMapSet = MapPathSet{
		&Path{{0, 1}, {0, 2}}:         {},
		&Path{{1, 1}, {1, 2}}:         {},
		&Path{{2, 1}, {2, 2}}:         {},
		&Path{{3, 1}, {3, 2}}:         {},
		&Path{{4, 1}, {4, 2}}:         {},
		&Path{{5, 1}, {5, 2}}:         {},
		&Path{{6, 1}, {6, 2}}:         {},
		&Path{{7, 1}, {7, 2}}:         {},
		&Path{{1, 1}, {1, 2}, {0, 2}}: {},
		&Path{{1, 1}, {1, 2}, {2, 2}}: {},
		&Path{{6, 1}, {6, 2}, {5, 2}}: {},
		&Path{{6, 1}, {6, 2}, {7, 2}}: {},
		&Path{{0, 5}, {0, 4}}:         {},
		&Path{{1, 5}, {1, 4}}:         {},
		&Path{{2, 5}, {2, 4}}:         {},
		&Path{{3, 5}, {3, 4}}:         {},
		&Path{{4, 5}, {4, 4}}:         {},
		&Path{{5, 5}, {5, 4}}:         {},
		&Path{{6, 5}, {6, 4}}:         {},
		&Path{{7, 5}, {7, 4}}:         {},
		&Path{{1, 6}, {1, 5}, {0, 5}}: {},
		&Path{{1, 6}, {1, 5}, {2, 5}}: {},
		&Path{{6, 6}, {6, 5}, {5, 5}}: {},
		&Path{{6, 6}, {6, 5}, {7, 5}}: {},
	}
	QueenOnRightPath      = Path{{6, 4}, {5, 4}, {3, 4}, {2, 4}, {1, 4}, {0, 4}}
	DeletePath            = Path{{6, 6}, {6, 5}, {5, 5}}
	QueenInMiddleSliceSet = SlicePathSet{
		{{4, 2}, {4, 1}, {4, 0}},
		{{3, 3}, {2, 3}, {1, 3}, {0, 3}},
		{{4, 4}, {4, 5}, {4, 6}, {4, 7}},
		{{5, 3}, {6, 3}, {7, 3}},
		{{3, 2}, {2, 1}, {1, 0}},
		{{5, 2}, {6, 1}, {7, 0}},
		{{3, 4}, {2, 5}, {1, 6}, {0, 7}},
		{{5, 4}, {6, 5}, {7, 6}},
	}
	QueenInMiddleMapSet = MapPathSet{
		&Path{{4, 2}, {4, 1}, {4, 0}}:         {},
		&Path{{3, 3}, {2, 3}, {1, 3}, {0, 3}}: {},
		&Path{{4, 4}, {4, 5}, {4, 6}, {4, 7}}: {},
		&Path{{5, 3}, {6, 3}, {7, 3}}:         {},
		&Path{{3, 2}, {2, 1}, {1, 0}}:         {},
		&Path{{5, 2}, {6, 1}, {7, 0}}:         {},
		&Path{{3, 4}, {2, 5}, {1, 6}, {0, 7}}: {},
		&Path{{5, 4}, {6, 5}, {7, 6}}:         {},
	}
	BaseChessSliceSetDoubled = SlicePathSet{
		{{0, 1}, {0, 2}},
		{{1, 1}, {1, 2}},
		{{2, 1}, {2, 2}},
		{{3, 1}, {3, 2}},
		{{4, 1}, {4, 2}},
		{{5, 1}, {5, 2}},
		{{6, 1}, {6, 2}},
		{{7, 1}, {7, 2}},
		{{1, 1}, {1, 2}, {0, 2}},
		{{1, 1}, {1, 2}, {2, 2}},
		{{6, 1}, {6, 2}, {5, 2}},
		{{6, 1}, {6, 2}, {7, 2}},
		{{0, 5}, {0, 4}},
		{{1, 5}, {1, 4}},
		{{2, 5}, {2, 4}},
		{{3, 5}, {3, 4}},
		{{4, 5}, {4, 4}},
		{{5, 5}, {5, 4}},
		{{6, 5}, {6, 4}},
		{{7, 5}, {7, 4}},
		{{1, 6}, {1, 5}, {0, 5}},
		{{1, 6}, {1, 5}, {2, 5}},
		{{6, 6}, {6, 5}, {5, 5}},
		{{6, 6}, {6, 5}, {7, 5}},
		{{0, 1}, {0, 2}},
		{{1, 1}, {1, 2}},
		{{2, 1}, {2, 2}},
		{{3, 1}, {3, 2}},
		{{4, 1}, {4, 2}},
		{{5, 1}, {5, 2}},
		{{6, 1}, {6, 2}},
		{{7, 1}, {7, 2}},
		{{1, 1}, {1, 2}, {0, 2}},
		{{1, 1}, {1, 2}, {2, 2}},
		{{6, 1}, {6, 2}, {5, 2}},
		{{6, 1}, {6, 2}, {7, 2}},
		{{0, 5}, {0, 4}},
		{{1, 5}, {1, 4}},
		{{2, 5}, {2, 4}},
		{{3, 5}, {3, 4}},
		{{4, 5}, {4, 4}},
		{{5, 5}, {5, 4}},
		{{6, 5}, {6, 4}},
		{{7, 5}, {7, 4}},
		{{1, 6}, {1, 5}, {0, 5}},
		{{1, 6}, {1, 5}, {2, 5}},
		{{6, 6}, {6, 5}, {5, 5}},
		{{6, 6}, {6, 5}, {7, 5}},
	}
	BaseChessMapSetDoubled = MapPathSet{
		&Path{{0, 1}, {0, 2}}:         {},
		&Path{{1, 1}, {1, 2}}:         {},
		&Path{{2, 1}, {2, 2}}:         {},
		&Path{{3, 1}, {3, 2}}:         {},
		&Path{{4, 1}, {4, 2}}:         {},
		&Path{{5, 1}, {5, 2}}:         {},
		&Path{{6, 1}, {6, 2}}:         {},
		&Path{{7, 1}, {7, 2}}:         {},
		&Path{{1, 1}, {1, 2}, {0, 2}}: {},
		&Path{{1, 1}, {1, 2}, {2, 2}}: {},
		&Path{{6, 1}, {6, 2}, {5, 2}}: {},
		&Path{{6, 1}, {6, 2}, {7, 2}}: {},
		&Path{{0, 5}, {0, 4}}:         {},
		&Path{{1, 5}, {1, 4}}:         {},
		&Path{{2, 5}, {2, 4}}:         {},
		&Path{{3, 5}, {3, 4}}:         {},
		&Path{{4, 5}, {4, 4}}:         {},
		&Path{{5, 5}, {5, 4}}:         {},
		&Path{{6, 5}, {6, 4}}:         {},
		&Path{{7, 5}, {7, 4}}:         {},
		&Path{{1, 6}, {1, 5}, {0, 5}}: {},
		&Path{{1, 6}, {1, 5}, {2, 5}}: {},
		&Path{{6, 6}, {6, 5}, {5, 5}}: {},
		&Path{{6, 6}, {6, 5}, {7, 5}}: {},
		&Path{{0, 1}, {0, 2}}:         {},
		&Path{{1, 1}, {1, 2}}:         {},
		&Path{{2, 1}, {2, 2}}:         {},
		&Path{{3, 1}, {3, 2}}:         {},
		&Path{{4, 1}, {4, 2}}:         {},
		&Path{{5, 1}, {5, 2}}:         {},
		&Path{{6, 1}, {6, 2}}:         {},
		&Path{{7, 1}, {7, 2}}:         {},
		&Path{{1, 1}, {1, 2}, {0, 2}}: {},
		&Path{{1, 1}, {1, 2}, {2, 2}}: {},
		&Path{{6, 1}, {6, 2}, {5, 2}}: {},
		&Path{{6, 1}, {6, 2}, {7, 2}}: {},
		&Path{{0, 5}, {0, 4}}:         {},
		&Path{{1, 5}, {1, 4}}:         {},
		&Path{{2, 5}, {2, 4}}:         {},
		&Path{{3, 5}, {3, 4}}:         {},
		&Path{{4, 5}, {4, 4}}:         {},
		&Path{{5, 5}, {5, 4}}:         {},
		&Path{{6, 5}, {6, 4}}:         {},
		&Path{{7, 5}, {7, 4}}:         {},
		&Path{{1, 6}, {1, 5}, {0, 5}}: {},
		&Path{{1, 6}, {1, 5}, {2, 5}}: {},
		&Path{{6, 6}, {6, 5}, {5, 5}}: {},
		&Path{{6, 6}, {6, 5}, {7, 5}}: {},
	}
)

var sliceAddResult SlicePathSet

func BenchmarkSliceAdd(b *testing.B) {
	var out SlicePathSet
	for n := 0; n < b.N; n++ {
		out = SliceAdd(BaseChessSliceSet, QueenOnRightPath)
	}
	sliceAddResult = out
}

var mapAddResult MapPathSet

func BenchmarkMapAdd(b *testing.B) {
	var out MapPathSet
	for n := 0; n < b.N; n++ {
		out = MapAdd(BaseChessMapSet, QueenOnRightPath)
	}
	mapAddResult = out
}

var sliceDeleteResult SlicePathSet

func BenchmarkSliceDelete(b *testing.B) {
	var out SlicePathSet
	for n := 0; n < b.N; n++ {
		out = SliceDelete(BaseChessSliceSet, DeletePath)
	}
	sliceDeleteResult = out
}

var mapDeleteResult MapPathSet

func BenchmarkMapDelete(b *testing.B) {
	var out MapPathSet
	for n := 0; n < b.N; n++ {
		out = MapDelete(BaseChessMapSet, DeletePath)
	}
	mapDeleteResult = out
}

var sliceCombineResult SlicePathSet

func BenchmarkSliceCombine(b *testing.B) {
	var out SlicePathSet
	for n := 0; n < b.N; n++ {
		out = SliceCombine(BaseChessSliceSet, QueenInMiddleSliceSet)
	}
	sliceCombineResult = out
}

var mapCombineResult MapPathSet

func BenchmarkMapCombine(b *testing.B) {
	var out MapPathSet
	for n := 0; n < b.N; n++ {
		out = MapCombine(BaseChessMapSet, QueenInMiddleMapSet)
	}
	mapCombineResult = out
}

var sliceReduceResult SlicePathSet

func BenchmarkSliceReduce(b *testing.B) {
	var out SlicePathSet
	for n := 0; n < b.N; n++ {
		out = SliceReduce(BaseChessSliceSetDoubled)
	}
	sliceReduceResult = out
}

var mapReduceResult MapPathSet

func BenchmarkMapReduce(b *testing.B) {
	var out MapPathSet
	for n := 0; n < b.N; n++ {
		out = MapReduce(BaseChessMapSetDoubled)
	}
	mapReduceResult = out
}

var sliceHasResult bool

func BenchmarkSliceHas(b *testing.B) {
	var out bool
	for n := 0; n < b.N; n++ {
		out = SliceHas(BaseChessSliceSet, QueenOnRightPath)
	}
	sliceHasResult = out
}

var mapHasResult bool

func BenchmarkMapHas(b *testing.B) {
	var out bool
	for n := 0; n < b.N; n++ {
		out = MapHas(BaseChessMapSet, QueenOnRightPath)
	}
	mapHasResult = out
}

var sliceEqualResult bool

func BenchmarkSliceEqual(b *testing.B) {
	var out bool
	for n := 0; n < b.N; n++ {
		out = SliceEqual(BaseChessSliceSet, BaseChessSliceSetDoubled)
	}
	sliceEqualResult = out
}

var mapEqualResult bool

func BenchmarkMapEqual(b *testing.B) {
	var out bool
	for n := 0; n < b.N; n++ {
		out = MapEqual(BaseChessMapSet, BaseChessMapSetDoubled)
	}
	mapEqualResult = out
}

var sliceDiffResult SlicePathSet

func BenchmarkSliceDiff(b *testing.B) {
	var out SlicePathSet
	for n := 0; n < b.N; n++ {
		out = SliceDiff(BaseChessSliceSet, QueenInMiddleSliceSet)
	}
	sliceDiffResult = out
}

var mapDiffResult MapPathSet

func BenchmarkMapDiff(b *testing.B) {
	var out MapPathSet
	for n := 0; n < b.N; n++ {
		out = MapDiff(BaseChessMapSet, QueenInMiddleMapSet)
	}
	mapDiffResult = out
}
