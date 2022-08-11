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

package mheap

import (
	"github.com/matrixorigin/nexus/pkg/vm/mempool"
	"github.com/matrixorigin/nexus/pkg/vm/mmu/guest"
)

func New(gm *guest.Mmu) *Mheap {
	return &Mheap{
		gm: gm,
		mp: mempool.New(),
	}
}

func (m *Mheap) Size() int64 {
	return m.gm.Size()
}

func (m *Mheap) HostSize() int64 {
	return m.gm.HostSize()
}

func (m *Mheap) Free(data []byte) {
	m.mp.Free(data)
	m.gm.Free(int64(cap(data)))
}

func (m *Mheap) Alloc(size int64) ([]byte, error) {
	data := m.mp.Alloc(int(size))
	if err := m.gm.Alloc(int64(cap(data))); err != nil {
		return nil, err
	}
	return data[:size], nil
}

func (m *Mheap) Grow(old []byte, size int64) ([]byte, error) {
	data, err := m.Alloc(mempool.Realloc(old, size))
	if err != nil {
		return nil, err
	}
	copy(data, old)
	return data[:size], nil
}

func (m *Mheap) PutSels(sels []int64) {
	m.Lock()
	defer m.Unlock()
	m.ss = append(m.ss, sels)
}

func (m *Mheap) GetSels() []int64 {
	m.Lock()
	defer m.Unlock()
	if len(m.ss) == 0 {
		return make([]int64, 0, 16)
	}
	sels := m.ss[0]
	m.ss = m.ss[1:]
	return sels[:0]
}
