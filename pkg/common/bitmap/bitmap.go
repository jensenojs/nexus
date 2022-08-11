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

package bitmap

import (
	"bytes"
	"fmt"
	"math/bits"

	"github.com/matrixorigin/nexus/pkg/common/encoding"
)

func New(n int) *Bitmap {
	return &Bitmap{
		len:  n,
		data: make([]uint64, (n-1)/64+1),
	}
}

func (n *Bitmap) Clear() {
	for i := range n.data {
		n.data[i] = 0
	}
}

func (n *Bitmap) Len() int {
	return n.len
}

func (n *Bitmap) Size() int {
	return len(n.data) * 8
}

// IsEmpty returns true if no bit in the Bitmap is set, otherwise it will return false.
func (n *Bitmap) IsEmpty() bool {
	for i := 0; i < len(n.data); i++ {
		if n.data[i] != 0 {
			return false
		}
	}
	return true
}

func (n *Bitmap) Add(row uint64) {
	n.data[row>>6] |= 1 << (row & 0x3F)
}

func (n *Bitmap) AddMany(rows []uint64) {
	for _, row := range rows {
		n.data[row>>6] |= 1 << (row & 0x3F)
	}
}

func (n *Bitmap) Remove(row uint64) {
	n.data[row>>6] &^= (uint64(1) << (row & 0x3F))
}

// Contains returns true if the row is contained in the Bitmap
func (n *Bitmap) Contains(row uint64) bool {
	return (n.data[row>>6] & (1 << (row & 0x3F))) != 0
}

func (n *Bitmap) AddRange(start, end uint64) {
	if start >= end {
		return
	}
	i, j := start>>6, (end-1)>>6
	if i == j {
		n.data[i] |= (^uint64(0) << uint(start&0x3F)) & (^uint64(0) >> (uint(-end) & 0x3F))
		return
	}
	n.data[i] |= (^uint64(0) << uint(start&0x3F))
	for k := i + 1; k < j; k++ {
		n.data[k] = ^uint64(0)
	}
	n.data[j] |= (^uint64(0) >> (uint(-end) & 0x3F))
}

func (n *Bitmap) RemoveRange(start, end uint64) {
	if start >= end {
		return
	}
	i, j := start>>6, (end-1)>>6
	if i == j {
		n.data[i] &= ^((^uint64(0) << uint(start&0x3F)) & (^uint64(0) >> (uint(-end) % 0x3F)))
		return
	}
	n.data[i] &= ^(^uint64(0) << uint(start&0x3F))
	for k := i + 1; k < j; k++ {
		n.data[k] = 0
	}
	n.data[j] &= ^(^uint64(0) >> (uint(-end) & 0x3F))
}

func (n *Bitmap) Or(m *Bitmap) {
	n.TryExpand(m)
	for i := 0; i < len(n.data); i++ {
		n.data[i] |= m.data[i]
	}
}

func (n *Bitmap) TryExpand(m *Bitmap) {
	if n.len < m.len {
		n.Expand(m.len)
	}
}

func (n *Bitmap) TryExpandWithSize(size int) {
	if n.len < size {
		n.Expand(size)
	}
}

func (n *Bitmap) Expand(size int) {
	data := make([]uint64, (size-1)/64+1)
	copy(data, n.data)
	n.len = size
	n.data = data
}

func (n *Bitmap) Filter(sels []int64) *Bitmap {
	m := New(n.len)
	for i, sel := range sels {
		if n.Contains(uint64(sel)) {
			m.Add(uint64(i))
		}
	}
	return m
}

func (n *Bitmap) Count() int {
	var cnt int

	for i := 0; i < len(n.data); i++ {
		cnt += bits.OnesCount64(n.data[i])
	}
	return cnt
}

func (n *Bitmap) ToArray() []uint64 {
	var rows []uint64

	start := uint64(0)
	for i := 0; i < len(n.data); i++ {
		bit := n.data[i]
		for bit != 0 {
			t := bit & -bit
			rows = append(rows, start+uint64(bits.OnesCount64(t-1)))
			bit ^= t
		}
		start += 64
	}
	return rows
}

func (n *Bitmap) Show() []byte {
	var buf bytes.Buffer

	buf.Write(encoding.EncodeUint64(uint64(n.len)))
	buf.Write(encoding.EncodeUint64(uint64(len(n.data) * 8)))
	buf.Write(encoding.EncodeSlice(n.data, 8))
	return buf.Bytes()
}

func (n *Bitmap) Read(data []byte) {
	n.len = int(encoding.DecodeUint64(data[:8]))
	data = data[8:]
	size := int(encoding.DecodeUint64(data[:8]))
	data = data[8:]
	n.data = encoding.DecodeSlice[uint64](data[:size], 8)
}

func (n *Bitmap) String() string {
	return fmt.Sprintf("%v", n.ToArray())
}
