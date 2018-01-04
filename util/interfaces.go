// Mgmt
// Copyright (C) 2013-2018+ James Shubin and the project contributors
// Written by James Shubin <james@shubin.ca> and the project contributors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package util

// Bool returns the interface value if it is a bool, and otherwise it panics.
func Bool(x interface{}) bool {
	b, ok := x.(bool)
	if !ok {
		panic("not a bool")
	}
	return b
}

// Uint returns the interface value if it is a uint, and otherwise it panics.
func Uint(x interface{}) uint {
	u, ok := x.(uint)
	if !ok {
		panic("not a uint")
	}
	return u
}
