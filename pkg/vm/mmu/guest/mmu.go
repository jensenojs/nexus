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

package guest

import (
	"github.com/matrixorigin/nexus/pkg/vm/mmu"
	"github.com/matrixorigin/nexus/pkg/vm/mmu/host"
)

func New(limit int64, mmu *host.Mmu) *Mmu {
	return &Mmu{
		mmu:   mmu,
		limit: limit,
	}
}

func (m *Mmu) Size() int64 {
	return m.size
}

func (m *Mmu) HostSize() int64 {
	return m.mmu.Size()
}

func (m *Mmu) Free(size int64) {
	if size == 0 {
		return
	}
	m.size -= size
	m.mmu.Free(size)
}

func (m *Mmu) Alloc(size int64) error {
	if size == 0 {
		return nil
	}
	if m.size+size > m.limit {
		return mmu.OutOfMemory
	}
	if err := m.mmu.Alloc(size); err != nil {
		return err
	}
	m.size += size
	return nil
}
