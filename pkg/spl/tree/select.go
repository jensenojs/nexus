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

type CommandStatement interface {
	Statement
	commandStatement()
}

func (*From) commandStatement()    {}
func (*Eval) commandStatement()    {}
func (*Join) commandStatement()    {}
func (*Limit) commandStatement()   {}
func (*Stats) commandStatement()   {}
func (*Where) commandStatement()   {}
func (*Fields) commandStatement()  {}
func (*OrderBy) commandStatement() {}
func (*Search) commandStatement()  {}

type Select struct {
	Cs []CommandStatement
}

func (n *Select) String() string {
	var s string

	for i, c := range n.Cs {
		if i > 0 {
			s += " | "
		}
		s += c.String()
	}
	return s
}
