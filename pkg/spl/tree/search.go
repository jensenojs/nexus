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

type Search struct {
	Expr string
	Time *SearchTime
}

type SearchTime struct {
	Start string
	End   string
}

func (s *Search) String() string {
	switch {
	case s.Time != nil && len(s.Expr) == 0:
		return fmt.Sprintf("SEARCH starttime=%s, endtime=%s", s.Time.Start, s.Time.End)
	case s.Time == nil && len(s.Expr) != 0:
		return fmt.Sprintf("SEARCH %s", s.Expr)
	case s.Time == nil && len(s.Expr) == 0:
		return fmt.Sprintf("SEARCH")
	default:
		return fmt.Sprintf("SEARCH starttime=%s, endtime=%s %s", s.Time.Start, s.Time.End, s.Expr)
	}
}
