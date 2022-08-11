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

package mempool

func New() *Mempool {
	return &Mempool{}
}

func (m *Mempool) Free(_ []byte) {
}

func (m *Mempool) Alloc(size int) (ret []byte) {
	return make([]byte, size)
}

func Realloc(data []byte, size int64) int64 {
	if data == nil {
		return size
	}
	n := int64(cap(data))
	if size <= n {
		return n
	}
	newcap := n
	doublecap := n + n
	if size > doublecap {
		newcap = size
	} else {
		if len(data) < 1024 {
			newcap = doublecap
		} else {
			for 0 < newcap && newcap < size {
				newcap += newcap / 4
			}
			if newcap <= 0 {
				newcap = size
			}
		}
	}
	return newcap
}
