/*
   Golang To Github Markdown Utility: gotomd
   Copyright (C) 2023-2025 Leslie Dancsecs

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package errs

import "errors"

// Exported errors.
var (
	ErrInvalidDefPerm     = errors.New("invalid default perm")
	ErrInvalidOutputDir   = errors.New("invalid output directory")
	ErrInvalidTemplate    = errors.New("invalid template")
	ErrInvalidArgument    = errors.New("invalid argument")
	ErrUnknownObject      = errors.New("unknown package object")
	ErrInvalidPackage     = errors.New("invalid package")
	ErrInvalidRelativeDir = errors.New("invalid relative directory")
	ErrInvalidDirectory   = errors.New("invalid directory")
	ErrMissingAction      = errors.New("missing action")
	ErrNoTestToRun        = errors.New("no tests to run")
	ErrNoPackageToRun     = errors.New("no package to run")
	ErrUnknownCommand     = errors.New("unknown command")
	ErrUnknownTag         = errors.New("unknown tag")
)
