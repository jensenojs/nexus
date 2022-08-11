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

package tree

import "fmt"

type OrderBy struct {
	Limit  *Value
	Orders OrderList
}

func (n OrderBy) String() string {
	var s string

	s = "SORT BY "
	if n.Limit != nil {
		s += n.Limit.String() + " "
	}
	for i := range n.Orders {
		if i > 0 {
			s += ", "
		}
		s += n.Orders[i].String()
	}
	return s
}

type OrderList []*Order

type Order struct {
	Type Direction
	E    ColunmName
}

// Direction for ordering results.
type Direction int8

// Direction values.
const (
	DefaultDirection Direction = iota
	Ascending
	Descending
)

var directionName = [...]string{
	DefaultDirection: "",
	Ascending:        "ASC",
	Descending:       "DESC",
}

func (i Direction) String() string {
	if i < 0 || i > Direction(len(directionName)-1) {
		return fmt.Sprintf("Direction(%d)", i)
	}
	return directionName[i]
}

func (n *Order) String() string {
	var s string

	s += n.E.String()
	if n.Type != DefaultDirection {
		s = " " + n.Type.String()
	}
	return s
}
