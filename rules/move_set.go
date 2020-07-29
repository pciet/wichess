package rules

import "fmt"

func MovesAddressSlice(the []MoveSet) []Address {
	out := make([]Address, 0, len(the)*2)
	for _, ms := range the {
	LOOP:
		for _, move := range ms.Moves {
			for _, outmove := range out {
				if outmove == move {
					continue LOOP
				}
			}
			out = append(out, move)
		}
	}
	return out
}

func MoveSetSliceHasMove(slice []MoveSet, m Move) bool {
	for _, moveset := range slice {
		for _, to := range moveset.Moves {
			move := Move{moveset.From, to}
			if move == m {
				return true
			}
		}
	}
	return false
}

func RemoveDuplicateMoveSetMoves(slice []MoveSet) []MoveSet {
	for i, ms := range slice {
		slice[i].Moves = RemoveAddressSliceDuplicates(ms.Moves)
	}
	return slice
}

func MoveSetSliceAdd(slice []MoveSet, at Address, to Address) []MoveSet {
	for i, ms := range slice {
		if ms.From != at {
			continue
		}
		slice[i].Moves = append(slice[i].Moves, to)
		return slice
	}
	return append(slice, MoveSet{at, []Address{to}})
}

func (a MoveSet) RemoveMove(to Address) MoveSet {
	index := -1
	for i, move := range a.Moves {
		if to == move {
			index = i
			break
		}
	}
	if index == -1 {
		Panic("didn't find", to, "in", a)
	}
	// assuming here that the order of Moves doesn't matter
	a.Moves[index] = a.Moves[len(a.Moves)-1]
	a.Moves = a.Moves[:len(a.Moves)-1]
	return a
}

func RemoveMoveSet(ms []MoveSet, from Address) []MoveSet {
	index := -1
	for i, moveset := range ms {
		if moveset.From == from {
			index = i
			break
		}
	}
	if index == -1 {
		Panic("didn't find", from, "MoveSet in", ms)
	}
	ms[index] = ms[len(ms)-1]
	return ms[:len(ms)-1]
}

func (a MoveSet) String() string { return a.From.String() + ":" + fmt.Sprint(a.Moves) }
