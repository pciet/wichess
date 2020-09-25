package rules

import "testing"

type RemoveAddressSliceDuplicatesCase struct {
	In  []Address
	Out []Address
}

var RemoveAddressSliceDuplicatesCases = []RemoveAddressSliceDuplicatesCase{
	{
		[]Address{
			{0, 1}, {0, 1}, {5, 2}, {5, 2},
		},
		[]Address{
			{0, 1}, {5, 2},
		},
	},
	{
		[]Address{
			{6, 1},
		},
		[]Address{
			{6, 1},
		},
	},
	{
		[]Address{
			{5, 5}, {5, 4}, {5, 4}, {5, 5}, {5, 5}, {5, 5}, {5, 5}, {5, 4},
		},
		[]Address{
			{5, 5}, {5, 4},
		},
	},
	{
		[]Address{
			{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}, {7, 7},
		},
		[]Address{
			{7, 7}, {6, 6}, {5, 5}, {4, 4}, {3, 3}, {2, 2}, {1, 1},
		},
	},
	{
		[]Address{
			{1, 1}, {1, 1}, {1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}, {7, 7},
		},
		[]Address{
			{7, 7}, {6, 6}, {5, 5}, {4, 4}, {3, 3}, {2, 2}, {1, 1},
		},
	},
	{
		[]Address{},
		[]Address{},
	},
	{
		[]Address{{0, 1}, {0, 1}, {0, 1}, {0, 1}, {0, 1}, {0, 1}, {0, 1}, {0, 1}},
		[]Address{{0, 1}},
	},
}

func TestRemoveAddressSliceDuplicates(t *testing.T) {
	for i, c := range RemoveAddressSliceDuplicatesCases {
		out := removeAddressSliceDuplicates(c.In)
		if len(out) != len(c.Out) {
			t.Fatal("case", i, "calculated len", len(out), "expected", len(c.Out))
		}
		for _, addr := range c.Out {
			testCount := AddressSliceHasCount(c.Out, addr)
			if testCount != 1 {
				t.Fatal("case", i, "Out slice has", testCount, "dupliate", addr)
			}

			count := AddressSliceHasCount(out, addr)
			if count != 1 {
				t.Fatal("case", i, "calculated slice has", count, "of", addr)
			}
		}
	}
}

func TestSquareEven(t *testing.T) {
	evenSquares := []AddressIndex{1, 3, 5, 12, 14, 17, 23, 24, 33, 42, 46, 55, 58, 62}
	for _, s := range evenSquares {
		if s.Address().squareEven() == false {
			t.Fatal("square address index", s, "incorrectly indicated as odd")
		}
	}
	oddSquares := []AddressIndex{0, 6, 9, 13, 16, 20, 29, 31, 34, 38, 41, 43, 50, 57, 59, 63}
	for _, s := range oddSquares {
		if s.Address().squareEven() {
			t.Fatal("square address index", s, "incorrectly indicated as even")
		}
	}
}

func TestAddressToIndex(t *testing.T) {
	addresses := []Address{{0, 0}, {7, 7}, {3, 5}, {6, 7}, {6, 0}, {7, 0}, {0, 7}}
	indices := []AddressIndex{0, 63, 43, 62, 6, 7, 56}
	for i, addr := range addresses {
		if addr.Index() != indices[i] {
			t.Fatal("address", addr, "doesn't match index", indices[i])
		}
		if indices[i].Address() != addr {
			t.Fatal("index", indices[i], "doesn't match address", addr)
		}
	}
}
