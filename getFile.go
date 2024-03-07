/*
   Golang To Github Markdown Utility: gotomd
   Copyright (C) 2023, 2024 Leslie Dancsecs

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

package main

import (
	"os"
)

const catCmd = "cat "

func getGoFile(cmd string) (string, error) {
	var (
		fData []byte
		res   string
	)

	dir, fName, err := parseCmds(cmd)
	for i, mi := 0, len(dir); i < mi && err == nil; i++ {
		fPath := dir[i] + string(os.PathSeparator) + fName[i]
		fData, err = os.ReadFile(fPath) //nolint:gosec // Ok.

		if err == nil {
			if res != "" {
				res += "\n\n"
			}

			res += "" +
				markBashCode(catCmd+fPath) + "\n\n" +
				markGoCode(string(fData))
		}
	}

	if err == nil {
		return res, nil
	}

	return "", err
}
