// Copyright (C) 2021 nexus.
//
// This file is part of nexus
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package lex

// GetKeywordID returns the lex id of the SQL keyword k or IDENT if k is
// not a keyword.
func GetKeywordID(k string) int32 {
	// The previous implementation generated a map that did a string ->
	// id lookup. Various ideas were benchmarked and the implementation below
	// was the fastest of those, between 3% and 10% faster (at parsing, so the
	// scanning speedup is even more) than the map implementation.
	switch k {
	case "and":
		return AND
	case "as":
		return AS
	case "asc":
		return ASC
	case "bool":
		return BOOL
	case "by":
		return BY
	case "cast":
		return CAST
	case "desc":
		return DESC
	case "end":
		return END
	case "eval":
		return EVAL
	case "endtime":
		return ENDTIME
	case "false":
		return FALSE
	case "fields":
		return FIELDS
	case "float":
		return FLOAT
	case "doube":
		return DOUBLE
	case "from":
		return FROM
	case "int":
		return INT
	case "join":
		return JOIN
	case "long":
		return LONG
	case "limit":
		return LIMIT
	case "not":
		return NOT
	case "or":
		return OR
	case "sort":
		return ORDER
	case "start":
		return START
	case "stats":
		return STATS
	case "string":
		return STRING
	case "true":
		return TRUE
	case "type":
		return TYPE
	case "where":
		return WHERE
	case "search":
		return SEARCH
	case "starttime":
		return STARTTIME
	default:
		return IDENT
	}
}
