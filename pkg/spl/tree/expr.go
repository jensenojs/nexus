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

import (
	"fmt"
	"go/constant"

	"github.com/matrixorigin/nexus/pkg/container/types"
)

type ExprStatement interface {
	Statement
	exprStatement()
}

type Value struct {
	Isnull bool
	Value  constant.Value
}

type FuncExpr struct {
	Name string
	Args ExprStatements
}

type TypeExpr struct {
	Typ types.Type
}

type ParenExpr struct {
	E ExprStatement
}

type ExprStatements []ExprStatement

func (*Value) exprStatement()         {}
func (*FuncExpr) exprStatement()      {}
func (*TypeExpr) exprStatement()      {}
func (*ParenExpr) exprStatement()     {}
func (ExprStatements) exprStatement() {}
func (ColunmName) exprStatement()     {}
func (ColunmNameList) exprStatement() {}

func (e *Value) String() string {
	if e.Isnull {
		return "NULL"
	}
	return e.Value.String()
}

func (e *FuncExpr) String() string {
	return fmt.Sprintf("%s(%s)", e.Name, e.Args)
}

func (e *TypeExpr) String() string {
	return e.Typ.String()
}

func (e *ParenExpr) String() string {
	return fmt.Sprintf("(%s)", e.E)
}

func (es ExprStatements) String() string {
	var s string

	for i := range es {
		if i > 0 {
			s += ", "
		}
		s += es[i].String()
	}
	return s
}
