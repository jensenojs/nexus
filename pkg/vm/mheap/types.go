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
	"sync"

	"github.com/matrixorigin/nexus/pkg/vm/mempool"
	"github.com/matrixorigin/nexus/pkg/vm/mmu/guest"
)

type Mheap struct {
	sync.Mutex
	// SelectList, temporarily stores the row number list in the execution of operators
	// and it can be reused in the future execution.
	ss [][]int64
	gm *guest.Mmu
	mp *mempool.Mempool
}
