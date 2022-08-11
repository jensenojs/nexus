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

package types

import (
	"fmt"
	"testing"
)

type typesTestCase struct {
	oid T
	e   Element
}

var tcs []typesTestCase

func init() {
	tcs = append(tcs, newTestCase(T_bool, Bool(true)))
	tcs = append(tcs, newTestCase(T_int32, Int(0)))
	tcs = append(tcs, newTestCase(T_int64, Long(0)))
	tcs = append(tcs, newTestCase(T_float32, Float(0)))
	tcs = append(tcs, newTestCase(T_float64, Double(0)))
	tcs = append(tcs, newTestCase(T_string, String("x")))
}

func TestSize(t *testing.T) {
	for _, tc := range tcs {
		tc.e.Size()
	}
}

func TestNew(t *testing.T) {
	for _, tc := range tcs {
		New(tc.oid)
	}
}

func TestTypeSize(t *testing.T) {
	for _, tc := range tcs {
		fmt.Printf("%v\n", TypeSize(tc.oid))
	}
}

func TestString(t *testing.T) {
	for _, tc := range tcs {
		fmt.Printf("%v\n", tc.oid.String())
	}
}

func TestTypeString(t *testing.T) {
	for _, tc := range tcs {
		fmt.Printf("%v\n", New(tc.oid).String())
	}
}

func newTestCase(oid T, e Element) typesTestCase {
	return typesTestCase{
		e:   e,
		oid: oid,
	}
}
