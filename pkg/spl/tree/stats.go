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

import "fmt"

type Stats struct {
	Ss StatList
	By ColunmNameList
}

type StatList []Stat

type Stat struct {
	As Name
	F  Name // func name
	A  Name // attribute name
}

func (n *Stats) String() string {
	s := fmt.Sprintf("STATS ")
	s += n.Ss.String()
	if len(n.By) > 0 {
		s += " " + n.By.String()
	}
	return s
}

func (ns StatList) String() string {
	var s string

	for i, n := range ns {
		if i > 0 {
			s += ", "
		}
		s += n.String()
	}
	return s
}

func (n Stat) String() string {
	return fmt.Sprintf("%s(%s) As %s", n.F, n.A, n.As)
}
