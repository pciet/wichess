// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import ()

type SlicePathSet []Path

func SliceAdd(the SlicePathSet, add Path) SlicePathSet {
	the = append(the, add)
	return the
}

func SliceDelete(the SlicePathSet, remove Path) SlicePathSet {
	out := make(SlicePathSet, 0, len(the))
	for _, item := range the {
		if item.Equal(remove) {
			continue
		}
		out = append(out, item)
	}
	return out
}

func SliceCombine(sets ...SlicePathSet) SlicePathSet {
	out := make(SlicePathSet, 0, len(sets[0])*len(sets))
	for _, set := range sets {
		for _, item := range set {
			out = append(out, item)
		}
	}
	return out
}

func SliceReduce(the SlicePathSet) SlicePathSet {
	out := make(SlicePathSet, 0, len(the))
	for _, item := range the {
		if SliceHas(out, item) {
			continue
		}
		out = append(out, item)
	}
	return out
}

func SliceHas(the SlicePathSet, item Path) bool {
	for _, path := range the {
		if path.Equal(item) {
			return true
		}
	}
	return false
}

func SliceEqual(sets ...SlicePathSet) bool {
	first := sets[0]
	length := len(first)
	for _, set := range sets {
		if length != len(set) {
			return false
		}
		// since slices cannot be compared there will be an unnecessary check on first
		for _, item := range first {
			if SliceHas(set, item) == false {
				return false
			}
		}
	}
	return true
}

func SliceDiff(a SlicePathSet, b SlicePathSet) SlicePathSet {
	out := make(SlicePathSet, 0, len(a))
	for _, item := range a {
		if SliceHas(b, item) == false {
			out = append(out, item)
		}
	}
	for _, item := range b {
		if SliceHas(a, item) == false {
			out = append(out, item)
		}
	}
	return out
}
