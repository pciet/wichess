package main

import (
	"strings"
)

// TODO: should constant SQL commands be prepared statements?
// TODO: ideally constant SQL queries would be constants instead of vars
// TODO: generate SQL symbol constants from postgres_tables.sql
// TODO: how to also simplify sql.Row.Scan?

func BuildSQLSelectQuery(selects []string, table string, whereEquals string) string {
	if len(selects) == 0 {
		panic("no columns to select")
	}
	if len(table) == 0 {
		panic("no table")
	}
	if len(whereEquals) == 0 {
		panic("no row selector")
	}

	var s strings.Builder

	s.WriteString("SELECT ")
	for i, v := range selects {
		s.WriteString(v)
		if i != len(selects)-1 {
			s.WriteString(", ")
		}
	}

	s.WriteString(" FROM ")
	s.WriteString(table)
	s.WriteString(" WHERE ")
	s.WriteString(whereEquals)
	s.WriteString("=$1;")
	return s.String()
}
