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

package process

import (
	"github.com/matrixorigin/nexus/pkg/container/batch"
	"github.com/matrixorigin/nexus/pkg/vm/mheap"
)

// New creates a new Process.
// A process stores the execution context.
func New(m *mheap.Mheap) *Process {
	return &Process{
		mp: m,
	}
}

func (proc *Process) GetMheap() *mheap.Mheap {
	return proc.mp
}

func (proc *Process) InputBatch() *batch.Batch {
	return proc.reg.inputBatch
}

func (proc *Process) SetInputBatch(bat *batch.Batch) {
	proc.reg.inputBatch = bat
}
