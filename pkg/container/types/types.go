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

package types

import (
	"fmt"
)

type T uint8

const (
	// any family
	T_any T = iota

	// bool family
	T_bool

	// integer family
	T_int32
	T_int64

	// float family
	T_float32
	T_float64

	// string family
	T_string
)

type Element interface {
	Size() int // return the size of space  the Element need
}

type Type struct {
	Oid  T
	Size int32 // e.g. int32.Size = 4, int64.Size = 8, string.Size = 24(SliceHeader size)
}

type Bool bool
type Int int32
type Long int64
type Float float32
type Double float64

type String []byte

type Ints interface {
	Int | Long
}

type Floats interface {
	Float | Double
}

type Number interface {
	Ints | Floats
}

type Generic interface {
	Ints | Floats
}

type All interface {
	Bool | Ints | Floats | String
}

func New(oid T) Type {
	return Type{Oid: oid, Size: int32(TypeSize(oid))}
}

func TypeSize(oid T) int {
	switch oid {
	case T_bool:
		return 1
	case T_int32, T_float32:
		return 4
	case T_int64, T_float64:
		return 8
	case T_string:
		return 24
	}
	return -1
}

func (t Type) TypeSize() int {
	return TypeSize(t.Oid)
}

func (t Type) String() string {
	return t.Oid.String()
}

func (t T) String() string {
	switch t {
	case T_bool:
		return "BOOL"
	case T_int32:
		return "INT"
	case T_int64:
		return "LONG"
	case T_float32:
		return "FLOAT"
	case T_float64:
		return "DOUBLE"
	case T_string:
		return "STRING"
	}
	return fmt.Sprintf("unexpected type: %d", t)
}
