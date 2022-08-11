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

package vector

import (
	"github.com/matrixorigin/nexus/pkg/common/bitmap"
	"github.com/matrixorigin/nexus/pkg/container/types"
	"github.com/matrixorigin/nexus/pkg/vm/mheap"
)

// Vector represent a column
type Vector[T types.Element] struct {
	// col represent the decoding column data
	col []T
	// Data represent the encoding column data
	data []byte
	// Type represent the type of column
	typ types.Type
	nsp *bitmap.Bitmap

	// Const used for const vector (a vector with a lot of rows of a same const value)
	isConst  bool
	constant struct {
		len int
	}

	// used for array and string
	array struct {
		offsets []uint64
		lengths []uint64
	}
}

// Vector represent a memory column
type AnyVector interface {
	Reset()
	Length() int
	SetLength(n int)
	Type() types.Type
	NewNulls(int)
	Shrink([]int64)
	Free(*mheap.Mheap)
	Nulls() *bitmap.Bitmap
	Realloc(size int, m *mheap.Mheap) error
}
