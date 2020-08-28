package wichess

import (
	"database/sql"
	"strconv"

	"github.com/lib/pq"
	"github.com/pciet/wichess/piece"
	"github.com/pciet/wichess/rules"
)

// FirstAvailable returns the first empty array index, or -1 if the array is full.
func (t Captures) FirstAvailable() int {
	for i, k := range t {
		if k == piece.NoKind {
			return i
		}
	}
	return -1
}

func AddGameCapture(tx *sql.Tx, in GameIdentifier, of rules.Orientation, k piece.Kind) {
	cs := LoadGameCaptures(tx, in, of, true)
	i := 0
	for ; i < 15; i++ {
		if cs[i] == piece.NoKind {
			break
		}
	}
	if i == 15 {
		Panic("more than 15 pieces taken in", in, "by", of)
	}
	var capturesKey string
	if of == rules.White {
		capturesKey = GamesWhiteCaptures
	} else if of == rules.Black {
		capturesKey = GamesBlackCaptures
	} else {
		Panic("orientation", of, "not white or black")
	}
	arrStr := capturesKey + "[" + strconv.Itoa(i) + "]"
	u := SQLUpdate(GamesTable, arrStr, GamesIdentifier)

	SQLExecRow(tx, u, k, in)
}

func LoadGameCaptures(tx *sql.Tx,
	in GameIdentifier, of rules.Orientation, forUpdate bool) Captures {

	var q string
	if of == rules.White {
		if forUpdate {
			q = GamesWhiteCapturesForUpdateQuery
		} else {
			q = GamesWhiteCapturesQuery
		}
	} else if of == rules.Black {
		if forUpdate {
			q = GamesBlackCapturesForUpdateQuery
		} else {
			q = GamesBlackCapturesQuery
		}
	} else {
		Panic("orientation", of, "not white or black")
	}

	var values []sql.NullInt64
	err := tx.QueryRow(q, in).Scan(pq.Array(&values))
	if err != nil {
		DebugPrintln(q, in)
		Panic(err)
	}

	if len(values) != 15 {
		Panic(in, "bad captures length", len(values))
	}

	var cs Captures
	for i, v := range values {
		if v.Valid == false {
			Panic(in, "sql null at", i)
		}
		cs[i] = piece.Kind(v.Int64)
	}

	return cs
}
