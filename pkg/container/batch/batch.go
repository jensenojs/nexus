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

package batch

import (
	"sync/atomic"

	"github.com/matrixorigin/nexus/pkg/container/vector"
	"github.com/matrixorigin/nexus/pkg/vm/mheap"
)

func New(n int, m *mheap.Mheap) *Batch {
	return &Batch{
		cnt:  1,
		vecs: make([]vector.AnyVector, n),
	}
}

func (bat *Batch) GetVectors() []vector.AnyVector {
	return bat.vecs
}

func (bat *Batch) GetVector(pos int) vector.AnyVector {
	return bat.vecs[pos]
}

func (bat *Batch) InBatch(vec vector.AnyVector) bool {
	for i := range bat.vecs {
		if vec == bat.vecs[i] {
			return true
		}
	}
	return false
}

func (bat *Batch) Length() int {
	return len(bat.zs)
}

func (bat *Batch) SetLength(n int) {
	for _, vec := range bat.vecs {
		vec.SetLength(n)
	}
	bat.zs = bat.zs[:n]
}

func (bat *Batch) Shrink(sels []int64) {
	mp := make(map[vector.AnyVector]uint8)
	for _, vec := range bat.vecs {
		if _, ok := mp[vec]; ok {
			continue
		}
		mp[vec]++
		vec.Shrink(sels)
	}
	vs := bat.zs
	for i, sel := range sels {
		vs[i] = vs[sel]
	}
	bat.zs = bat.zs[:len(sels)]
}

func (bat *Batch) Free(m *mheap.Mheap) {
	if atomic.AddInt64(&bat.cnt, -1) != 0 {
		return
	}
	for _, vec := range bat.vecs {
		if vec != nil {
			vec.Free(m)
		}
	}
	m.PutSels(bat.zs)
	bat.zs = nil
	bat.vecs = nil
}

func (bat *Batch) SetZs(n int, m *mheap.Mheap) {
	if cap(bat.zs) < n {
		m.PutSels(bat.zs)
		bat.zs = make([]int64, n)
	}
	bat.zs = bat.zs[:n]
	for i := 0; i < n; i++ {
		bat.zs[i] = 1
	}
}
