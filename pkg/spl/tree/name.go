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

package tree

type Name string

func (n Name) String() string {
	return string(n)
}

type NameList []Name

func (n NameList) String() string {
	var s string

	for i := range n {
		if i > 0 {
			s += ", "
		}
		s += n[i].String()
	}
	return s
}

type ColunmName struct {
	Path Name
}

func (n ColunmName) String() string {
	return n.Path.String()
}

type ColunmNameList []ColunmName

func (n ColunmNameList) String() string {
	var s string

	for i := range n {
		if i > 0 {
			s += ", "
		}
		s += n[i].String()
	}
	return s
}
