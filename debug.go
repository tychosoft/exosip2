// Copyright (C) 2023 David Sugar, Tycho Softworks
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
//go:build debug

package exosip2

type EVT_TYPE string

func IsDebug() bool {
	return true
}

const (
	EVT_IDLE     EVT_TYPE = "idle"
	EVT_STARTUP  EVT_TYPE = "startup"
	EVT_SHUTDOWN EVT_TYPE = "shutdown"
	EVT_INVALID  EVT_TYPE = "invalid"
	EVT_REGISTER EVT_TYPE = "register"
	EVT_MESSAGE  EVT_TYPE = "message"
)
