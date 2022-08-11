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
	"fmt"

	"github.com/matrixorigin/nexus/pkg/common/bitmap"
	"github.com/matrixorigin/nexus/pkg/common/encoding"
	"github.com/matrixorigin/nexus/pkg/container/types"
	"github.com/matrixorigin/nexus/pkg/vm/mheap"
)

func New[T types.Element](typ types.Type) *Vector[T] {
	return &Vector[T]{
		typ:  typ,
		col:  []T{},
		data: []byte{},
		nsp:  bitmap.New(0),
	}
}

func NewVector[T types.Element](typ types.Type, nsp *bitmap.Bitmap, col []T, data []byte) *Vector[T] {
	return &Vector[T]{
		typ:  typ,
		col:  col,
		nsp:  nsp,
		data: data,
	}
}

func NewConstVector[T types.Element](typ types.Type, col []T, len int) *Vector[T] {
	return &Vector[T]{
		typ:     typ,
		col:     col,
		isConst: true,
		constant: struct {
			len int
		}{
			len: len,
		},
		nsp: bitmap.New(1),
	}
}

func (v *Vector[T]) Reset() {
	v.col = v.col[:0]
	if !v.isConst {
		v.data = v.data[:0]
	}
	if _, ok := (any)(v).(*Vector[types.String]); ok {
		v.array.offsets = v.array.offsets[:0]
		v.array.lengths = v.array.lengths[:0]
	}
}

func (v *Vector[T]) IsConst() bool {
	return v.isConst
}

func (v *Vector[T]) Column() []T {
	return v.col
}

func (v *Vector[T]) Length() int {
	return len(v.col)
}

func (v *Vector[T]) SetLength(n int) {
	if v.isConst {
		v.constant.len = n
		return
	}
	switch (any)(v).(type) {
	case *Vector[types.String]:
		v.array.offsets = v.array.offsets[:n]
		v.array.lengths = v.array.lengths[:n]
		v.data = v.data[:v.array.offsets[n-1]+v.array.lengths[n-1]]
	default:
		v.data = v.data[:n*v.typ.TypeSize()]
	}
	v.col = v.col[:n]
}

func (v *Vector[T]) Type() types.Type {
	return v.typ
}

func (v *Vector[T]) Nulls() *bitmap.Bitmap {
	return v.nsp
}

func (v *Vector[T]) NewNulls(n int) {
	v.nsp = bitmap.New(n)
}

func (v *Vector[T]) Free(m *mheap.Mheap) {
	if v.data != nil {
		m.Free(v.data)
	}
}

func (v *Vector[T]) Realloc(size int, m *mheap.Mheap) error {
	oldLen := len(v.data)
	data, err := m.Grow(v.data, int64(cap(v.data)+size))
	if err != nil {
		return err
	}
	m.Free(v.data)
	v.data = data[:oldLen]
	switch vec := (any)(v).(type) {
	case *Vector[types.String]:
		vec.col = vec.col[:0]
		for i, off := range vec.array.offsets {
			vec.col = append(vec.col, vec.data[off:off+vec.array.lengths[i]])
		}
	default:
		v.col = encoding.DecodeSlice[T](v.data[:len(data)], size)[:oldLen/size]
	}
	return nil
}

func (v *Vector[T]) Append(w T, m *mheap.Mheap) error {
	switch vec := (any)(v).(type) {
	case *Vector[types.String]:
		wv, _ := (any)(w).(types.String)
		n := len(v.data)
		if n+w.Size() >= cap(v.data) {
			if err := v.Realloc(n+w.Size()-cap(v.data)+1, m); err != nil {
				return err
			}
		}
		vec.array.lengths = append(vec.array.lengths, uint64(len(wv)))
		vec.array.offsets = append(vec.array.offsets, uint64(len(v.data)))
		size := len(vec.data)
		vec.data = append(vec.data, wv...)
		vec.col = append(vec.col, vec.data[size:size+len(wv)])
	default:
		n := len(v.col)
		if n+1 >= cap(v.col) {
			if err := v.Realloc(w.Size(), m); err != nil {
				return err
			}
		}
		v.col = append(v.col, w)
		v.data = v.data[:(n+1)*w.Size()]
	}
	return nil
}

func (v *Vector[T]) Dup(m *mheap.Mheap) (*Vector[T], error) {
	w := New[T](v.typ)
	if len(v.col) == 0 {
		return w, nil
	}
	if v.isConst {
		w.col = append(w.col, v.col...)
		if len(v.array.lengths) > 0 {
			w.array.lengths = append(w.array.lengths, v.array.lengths...)
			w.array.offsets = append(w.array.offsets, v.array.offsets...)
		}
		return w, nil
	}
	if err := w.Realloc(len(v.data), m); err != nil {
		return nil, err
	}
	w.data = w.data[:len(v.data)]
	copy(w.data, v.data)
	switch vec := (any)(v).(type) {
	case *Vector[types.String]:
		wv := (any)(w).(*Vector[types.String])
		wv.col = make([]types.String, len(vec.col))
		w.array.lengths = make([]uint64, len(v.array.lengths))
		w.array.offsets = make([]uint64, len(v.array.offsets))
		copy(w.array.lengths, v.array.lengths)
		copy(w.array.offsets, v.array.offsets)
		for i, o := range w.array.offsets {
			wv.col[i] = w.data[o : o+w.array.lengths[i]]
		}
	default:
		size := v.col[0].Size()
		w.col = encoding.DecodeSlice[T](w.data, size)[:len(w.data)/size]
	}
	return w, nil
}

func (v *Vector[T]) Shrink(sels []int64) {
	if v.isConst || len(v.col) == 0 {
		return
	}
	size := v.col[0].Size()
	switch (any)(v).(type) {
	case *Vector[types.String]:
		for i, sel := range sels {
			v.array.offsets[i] = v.array.offsets[sel]
			v.array.lengths[i] = v.array.lengths[sel]
		}
		v.array.offsets = v.array.offsets[:len(sels)]
		v.array.lengths = v.array.lengths[:len(sels)]
	default:
		v.data = v.data[:len(sels)*size]
	}
	for i, sel := range sels {
		v.col[i] = v.col[sel]
	}
	v.col = v.col[:len(sels)]
	v.nsp = v.nsp.Filter(sels)
}

func (v *Vector[T]) Shuffle(sels []int64, m *mheap.Mheap) error {
	size := v.col[0].Size()
	switch vec := (any)(v).(type) {
	case *Vector[types.String]:
		maxSize := uint64(0)
		for _, sel := range sels {
			if size := v.array.offsets[sel] + v.array.lengths[sel]; size > maxSize {
				maxSize = size
			}
		}
		data, err := m.Alloc(int64(maxSize))
		if err != nil {
			return err
		}
		o := uint64(0)
		data = data[:0]
		os := make([]uint64, len(sels))
		ns := make([]uint64, len(sels))
		for i, sel := range sels {
			data = append(data, vec.col[sel]...)
			os[i] = o
			ns[i] = uint64(len(vec.col[sel]))
			o += ns[i]
		}
		m.Free(v.data)
		v.data = data
		vec.col = vec.col[:len(sels)]
		for i, o := range os {
			vec.col[i] = v.data[o : o+ns[i]]
		}
		copy(v.array.offsets, os)
		copy(v.array.lengths, ns)
	default:
		ws := make([]T, len(v.col))
		for i, sel := range sels {
			ws[i] = v.col[sel]
		}
		v.col = v.col[:len(sels)]
		v.data = v.data[:len(sels)*size]
		copy(v.col, ws)
	}
	v.nsp = v.nsp.Filter(sels)
	return nil
}

func (v *Vector[T]) UnionOne(w *Vector[T], sel int64, m *mheap.Mheap) error {
	return v.Append(w.col[sel], m)
}

func (v *Vector[T]) UnionNull(_ *Vector[T], _ int64, m *mheap.Mheap) error {
	var val T

	if err := v.Append(val, m); err != nil {
		return err
	}
	return nil
}

func (v *Vector[T]) UnionBatch(w *Vector[T], offset int64, cnt int, flags []uint8, m *mheap.Mheap) error {
	switch vec := (any)(v).(type) {
	case *Vector[types.String]:
		wv := (any)(w).(*Vector[types.String])
		incSize := 0
		for i, flg := range flags {
			if flg > 0 {
				incSize += int(vec.array.lengths[int(offset)+i])
			}
		}
		n := len(v.data)
		if n+incSize >= cap(v.data) {
			if err := v.Realloc(n+incSize-cap(v.data)+1, m); err != nil {
				return err
			}
		}
		o := uint64(len(v.data))
		for i, flg := range flags {
			if flg > 0 {
				from := wv.col[int(offset)+i]
				v.array.offsets = append(v.array.offsets, o)
				v.array.lengths = append(v.array.lengths, uint64(len(from)))
				v.data = append(v.data, from...)
				vec.col = append(vec.col, v.data[o:o+uint64(len(from))])
				o += uint64(len(from))
			}
		}
	default:
		n := len(v.col)
		if n+cnt >= cap(v.col) {
			if err := v.Realloc(n+cnt-cap(v.col)+1, m); err != nil {
				return err
			}
		}
		v.col = v.col[:n+cnt]
		for i, j := 0, n; i < len(flags); i++ {
			if flags[i] > 0 {
				v.col[j] = w.col[int(offset)+i]
				j++
			}
		}
	}
	return nil
}

func (v *Vector[T]) String() string {
	return fmt.Sprintf("%v-%v", v.col, v.nsp)
}
