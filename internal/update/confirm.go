/*
   Golang To Github Markdown Utility: gotomd
   Copyright (C) 2025 Leslie Dancsecs

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

package update

import (
	"fmt"
	"path/filepath"

	"github.com/dancsecs/szlog"
)

// Overwrite messages.
const (
	ConfirmMsg = "Confirm overwrite of: %s\n" +
		"(Y to overwrite; [N/n] to Cancel; [R/r] to review)? "
	ConfirmCancelled = "Overwrite cancelled\n\n"
	ConfirmUnknown   = "Unknown response: '%s'\n\n"
)

func confirm(fPath, oldData, newData string) (bool, error) {
	var (
		ok  bool
		err error
	)

	response := "R"
	for err == nil && response == "R" {
		szlog.Say0f(ConfirmMsg, fPath)

		if _, err = fmt.Scanln(&response); err == nil {
			switch response {
			case "Y":
				ok = true
			case "N", "n":
				szlog.Say0(ConfirmCancelled)
			case "R", "r":
				diffStr := diffFile(
					filepath.Base(fPath),
					oldData,
					newData,
				)
				szlog.Say0(diffStr, "\n")

				response = "R"
			default:
				szlog.Say0f(ConfirmUnknown, response)
				response = "R" // Ask again.
			}
		}
	}

	return ok, err //nolint:wrapcheck // Caller will wrap error.
}
