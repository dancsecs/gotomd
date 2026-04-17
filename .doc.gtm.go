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

/*
<!--- gotomd::inline-run::./. --help -->

<!--- gotomd::snippet::./.directives.sds.md # START SNIPPET -->

# Dedication

This project is dedicated to Reem. Your brilliance, courage, and quiet
strength continue to inspire me. Every line is written in gratitude for the
light and hope you brought into my life.

NOTE: Documentation reviewed and polished with the assistance of ChatGPT from
OpenAI.
*/
package ////main

import (
	"strings"

	"github.com/dancsecs/szlog"
)

const copyrightMessage = `
//<<<! gotomd::file::~/.copyright >>
`

// Copyright writes the copyright message to os.Stdout.
func Copyright() {
	szlog.Say0(strings.Trim(copyrightMessage, " \t\n") + "\n")
}
