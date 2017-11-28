// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import ()

type MapPathSet map[*Path]struct{}

func MapAdd(the MapPathSet, add Path) MapPathSet {
	the[&add] = struct{}{}
	return the
}

func MapDelete(the MapPathSet, remove Path) MapPathSet {
	out := make(MapPathSet)
	for item, _ := range the {
		if (*item).Equal(remove) {
			continue
		}
		out[item] = struct{}{}
	}
	return out
}

func MapCombine(sets ...MapPathSet) MapPathSet {
	out := make(MapPathSet)
	for _, set := range sets {
		for item, _ := range set {
			out[item] = struct{}{}
		}
	}
	return out
}

func MapReduce(the MapPathSet) MapPathSet {
	out := make(MapPathSet)
	for item, _ := range the {
		if MapHas(out, *item) {
			continue
		}
		out[item] = struct{}{}
	}
	return out
}

func MapHas(the MapPathSet, item Path) bool {
	for path, _ := range the {
		if (*path).Equal(item) {
			return true
		}
	}
	return false
}

func MapEqual(sets ...MapPathSet) bool {
	first := sets[0]
	length := len(first)
	for _, set := range sets {
		if length != len(set) {
			return false
		}
		for item, _ := range first {
			if MapHas(set, *item) == false {
				return false
			}
		}
	}
	return true
}

func MapDiff(a MapPathSet, b MapPathSet) MapPathSet {
	out := make(MapPathSet)
	for item, _ := range a {
		if MapHas(b, *item) == false {
			out[item] = struct{}{}
		}
	}
	for item, _ := range b {
		if MapHas(a, *item) == false {
			out[item] = struct{}{}
		}
	}
	return out
}
