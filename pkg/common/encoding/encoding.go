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

package encoding

import (
	"unsafe"
)

func EncodeUint64(v uint64) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(&v)), 8)
}

func DecodeUint64(v []byte) uint64 {
	return *(*uint64)(unsafe.Pointer(&v[0]))
}

func EncodeSlice[T any](v []T, sz int) (ret []byte) {
	if len(v) > 0 {
		ret = unsafe.Slice((*byte)(unsafe.Pointer(&v[0])), cap(v)*sz)[:len(v)*sz]
	}
	return
}

func DecodeSlice[T any](v []byte, sz int) (ret []T) {
	if len(v) > 0 {
		ret = unsafe.Slice((*T)(unsafe.Pointer(&v[0])), cap(v)/sz)[:len(v)/sz]
	}
	return
}
