package main

import (
	"log"
	"strconv"
	"strings"
)

// TODO: should constant SQL commands be prepared statements?
// TODO: ideally constant SQL queries would be constants instead of vars
// TODO: generate SQL symbol constants from postgres_tables.sql
// TODO: how to also simplify sql.Row.Scan?

func BuildSQLGeneralizedWhereQuery(selects []string, table string, where string) string {
	return BuildSQLQueryImplementation(selects, table, where, false)
}

// TODO: have strings.Builder at top here to avoid the extra + concatenation?

func BuildSQLQuery(selects []string, table string, whereEquals string) string {
	return BuildSQLQueryImplementation(selects, table, whereEquals+"=$1", false)
}

func BuildSQLForUpdateQuery(selects []string, table string, whereEquals string) string {
	return BuildSQLQueryImplementation(selects, table, whereEquals+"=$1", true)
}

func BuildSQLQueryImplementation(selects []string, table string, where string, forUpdate bool) string {
	if table == "" {
		log.Panic("no table")
	}
	if where == "" {
		log.Panic("no row selector")
	}

	var s strings.Builder

	s.WriteString("SELECT ")
	if selects == nil {
		s.WriteString("null")
	} else {
		if len(selects) == 0 {
			log.Panic("no selects")
		}
		for i, v := range selects {
			if v == "" {
				log.Panicln("empty select at index", i)
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
		log.Println("SQL Query:", s.String())
	}
	return s.String()
}

func BuildSQLInsert(table string, inserts []string) string {
	if table == "" {
		log.Panic("no table")
	}
	if len(inserts) == 0 {
		log.Panic("no inserts")
	}

	var s strings.Builder

	s.WriteString("INSERT INTO ")
	s.WriteString(table)
	s.WriteString(" (")
	for i, v := range inserts {
		if v == "" {
			log.Panicln("empty insert at index", i)
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
	s.WriteString(");")
	if DebugSQL {
		log.Println("SQL Insert:", s.String())
	}
	return s.String()
}

func BuildSQLUpdate(table string, set string, whereEquals string) string {
	if table == "" {
		log.Panic("no table")
	}
	if set == "" {
		log.Panic("no key to set")
	}
	if whereEquals == "" {
		log.Panic("no where equals key")
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
		log.Println("SQL Update:", s.String())
	}
	return s.String()
}
