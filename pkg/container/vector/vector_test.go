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
	"testing"

	"github.com/matrixorigin/nexus/pkg/container/types"
	"github.com/matrixorigin/nexus/pkg/vm/mheap"
	"github.com/matrixorigin/nexus/pkg/vm/mmu/guest"
	"github.com/matrixorigin/nexus/pkg/vm/mmu/host"
	"github.com/stretchr/testify/require"
)

func TestAppend(t *testing.T) {
	hm := host.New(1 << 20)
	gm := guest.New(1<<20, hm)
	m := mheap.New(gm)
	vx := New[types.Long](types.New(types.T_int64))
	fmt.Printf("type: %v\n", vx.Type())
	vx.Append(types.Long(1), m)
	vx.Append(types.Long(2), m)
	vx.Append(types.Long(3), m)
	fmt.Printf("%v\n", vx)
	fmt.Printf("vx: %v: %v\n", vx.col, vx.data)
	vx.Reset()
	vx.Append(types.Long(3), m)
	vx.Append(types.Long(1), m)
	vx.SetLength(2)
	vx.Append(types.Long(2), m)
	fmt.Printf("vx: %v: %v\n", vx.col, vx.data)
	fmt.Printf("length: %v\n", vx.Length())
	vy, err := vx.Dup(m)
	require.NoError(t, err)
	fmt.Printf("vy: %v\n", vy)
	vy.Shrink([]int64{0, 1, 2})
	err = vx.UnionOne(vy, 0, m)
	require.NoError(t, err)
	err = vx.UnionNull(vy, 0, m)
	require.NoError(t, err)
	err = vx.UnionBatch(vy, 0, 1, []uint8{1}, m)
	require.NoError(t, err)
	err = vy.Shuffle([]int64{1}, m)
	require.NoError(t, err)

	vx.Free(m)
	vy.Free(m)
	require.Equal(t, int64(0), m.Size())
}

func TestAppendStr(t *testing.T) {
	hm := host.New(1 << 20)
	gm := guest.New(1<<20, hm)
	m := mheap.New(gm)
	vx := New[types.String](types.New(types.T_string))
	fmt.Printf("type: %v\n", vx.Type())
	vx.Append(types.String("1"), m)
	vx.Append(types.String("2"), m)
	vx.Append(types.String("3"), m)
	fmt.Printf("%v\n", vx)
	fmt.Printf("vx: %v: %v\n", vx.col, vx.data)
	vx.Reset()
	vx.Append(types.String("3"), m)
	vx.Append(types.String("1"), m)
	vx.SetLength(2)
	vx.Append(types.String("2"), m)
	fmt.Printf("vx: %v: %v\n", vx.col, vx.data)
	fmt.Printf("length: %v\n", vx.Length())
	vy, err := vx.Dup(m)
	require.NoError(t, err)
	fmt.Printf("vy: %v\n", vy)
	vy.Shrink([]int64{0, 1, 2})
	err = vx.UnionOne(vy, 0, m)
	require.NoError(t, err)
	err = vx.UnionNull(vy, 0, m)
	require.NoError(t, err)
	err = vx.UnionBatch(vy, 0, 1, []uint8{1}, m)
	require.NoError(t, err)
	err = vy.Shuffle([]int64{1}, m)
	require.NoError(t, err)

	vx.Free(m)
	vy.Free(m)
	require.Equal(t, int64(0), m.Size())
}
