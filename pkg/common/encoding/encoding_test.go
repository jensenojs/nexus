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

package encoding

import (
	"fmt"
	"testing"

	"github.com/matrixorigin/nexus/pkg/container/types"
)

func TestEncode(t *testing.T) {
	vs := make([]types.Long, 10)
	for i := 0; i < 10; i++ {
		vs[i] = types.Long(i)
	}
	data := EncodeSlice(vs, 8)
	fmt.Printf("data: %v\n", data)
	rs := DecodeSlice[types.Long](data, 8)
	fmt.Printf("rs: %v\n", rs)
}
