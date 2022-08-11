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
	"context"

	"github.com/matrixorigin/nexus/pkg/container/batch"
	"github.com/matrixorigin/nexus/pkg/vm/mheap"
)

// WaitRegister channel
type WaitRegister struct {
	ctx context.Context
	ch  chan *batch.Batch
}

// Register used in execution pipeline and shared with all operators of the same pipeline.
type Register struct {
	// InputBatch, stores the result of the previous operator.
	inputBatch *batch.Batch
	// MergeReceivers, receives result of multi previous operators from other pipelines
	// e.g. merge operator.
	mergeReceivers []*WaitRegister
}

//Limitation specifies the maximum resources that can be used in one query.
type Limitation struct {
}

// Process contains context used in query execution
// one or more pipeline will be generated for one query,
// and one pipeline has one process instance.
type Process struct {
	// Id, query id.
	id  string
	reg Register
	lim Limitation
	mp  *mheap.Mheap

	// unix timestamp
	unixTime int64

	cancel context.CancelFunc
}
