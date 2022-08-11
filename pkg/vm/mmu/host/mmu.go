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

package host

import (
	"sync/atomic"

	"github.com/matrixorigin/nexus/pkg/vm/mmu"
)

func New(limit int64) *Mmu {
	return &Mmu{
		limit: limit,
	}
}

func (m *Mmu) Size() int64 {
	return atomic.LoadInt64(&m.size)
}

func (m *Mmu) Free(size int64) {
	atomic.AddInt64(&m.size, size*-1)
}

func (m *Mmu) Alloc(size int64) error {
	if atomic.LoadInt64(&m.size)+size > m.limit {
		return mmu.OutOfMemory
	}
	for v := atomic.LoadInt64(&m.size); !atomic.CompareAndSwapInt64(&m.size, v, v+size); v = atomic.LoadInt64(&m.size) {
	}
	return nil
}
