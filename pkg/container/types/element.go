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

func (_ Bool) Size() int {
	return 1
}

func (_ Int) Size() int {
	return 4
}

func (_ Long) Size() int {
	return 8
}

func (_ Float) Size() int {
	return 4
}

func (_ Double) Size() int {
	return 8
}

func (s String) Size() int {
	return len(s)
}
