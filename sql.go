package wichess

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO: should constant SQL commands be prepared statements?
// TODO: ideally constant SQL queries would be constants instead of vars
// TODO: generate SQL symbol constants from postgres_tables.sql
// TODO: how to also simplify sql.Row.Scan?

func SQLGeneralizedWhereQuery(selects []string, table, where string) string {
	return SQLQueryImplementation(selects, table, where, false)
}

// TODO: have strings.Builder at top here to avoid the extra + concatenation?

func SQLQuery(selects []string, table, whereEquals string) string {
	return SQLQueryImplementation(selects, table, whereEquals+"=$1", false)
}

func SQLForUpdateQuery(selects []string, table, whereEquals string) string {
	return SQLQueryImplementation(selects, table, whereEquals+"=$1", true)
}

func SQLQueryImplementation(selects []string, table, where string, forUpdate bool) string {
	if table == "" {
		Panic("no table")
	}
	if where == "" {
		Panic("no row selector")
	}

	var s strings.Builder

	s.WriteString("SELECT ")
	if selects == nil {
		s.WriteString("null")
	} else {
		if len(selects) == 0 {
			Panic("no selects")
		}
		for i, v := range selects {
			if v == "" {
				Panic("empty select at index", i)
			}
			s.WriteString(v)
			if i != len(selects)-1 {
				s.WriteString(", ")
			}
		}
	}
	s.WriteString(" FROM ")
	s.WriteString(table)
	s.WriteString(" WHERE ")
	s.WriteString(where)
	if forUpdate {
		s.WriteString(" FOR UPDATE")
	}
	s.WriteRune(';')
	if DebugSQL {
		fmt.Println(s.String())
	}
	return s.String()
}

func SQLInsert(table string, inserts []string) string {
	return SQLInsertReturning(table, inserts, "")
}

func SQLInsertReturning(table string, inserts []string, returns string) string {
	if table == "" {
		Panic("no table")
	}
	if len(inserts) == 0 {
		Panic("no inserts")
	}

	var s strings.Builder

	s.WriteString("INSERT INTO ")
	s.WriteString(table)
	s.WriteString(" (")
	for i, v := range inserts {
		if v == "" {
			Panic("empty insert at index", i)
		}
		s.WriteString(v)
		if i != len(inserts)-1 {
			s.WriteString(", ")
		}
	}
	s.WriteString(") VALUES (")
	for i := 1; i <= len(inserts); i++ {
		s.WriteRune('$')
		s.WriteString(strconv.Itoa(i))
		if i != len(inserts) {
			s.WriteString(", ")
		}
	}
	if returns == "" {
		s.WriteString(");")
	} else {
		s.WriteString(") RETURNING ")
		s.WriteString(returns)
		s.WriteRune(';')
	}
	if DebugSQL {
		fmt.Println(s.String())
	}
	return s.String()
}

func SQLUpdate(table, set, whereEquals string) string {
	if table == "" {
		Panic("no table")
	}
	if set == "" {
		Panic("no key to set")
	}
	if whereEquals == "" {
		Panic("no where equals key")
	}

	var s strings.Builder

	s.WriteString("UPDATE ")
	s.WriteString(table)
	s.WriteString(" SET ")
	s.WriteString(set)
	s.WriteString(" = $1 WHERE ")
	s.WriteString(whereEquals)
	s.WriteString(" = $2;")
	if DebugSQL {
		fmt.Println(s.String())
	}
	return s.String()
}

func SQLDelete(table, whereEquals string) string {
	if table == "" {
		Panic("no table")
	}
	if whereEquals == "" {
		Panic("no where equals key")
	}

	var s strings.Builder

	s.WriteString("DELETE FROM ")
	s.WriteString(table)
	s.WriteString(" WHERE ")
	s.WriteString(whereEquals)
	s.WriteString("=$1;")
	return s.String()
}
