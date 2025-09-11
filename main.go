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
Package gotomd maintains GitHub-style README.md files by embedding Go
documentation, source code, test output, and command output directly
from the Go codebase. This ensures that program documentation is kept
in one place—the Go code itself—while keeping the README automatically
up to date.

## How it works

gotomd processes a Markdown file in one of two ways:

 1. **Template mode** — You maintain a separate template file named `*.gtm.md`.
 2. **In-place mode** — You insert directives directly into an existing
    `*.md` file, and gotomd replaces them in place.

Directives are written inside HTML-style comments:

```html
<!---
gotomd::ACTION::PARAMETERS
-->
```

When processing the file, gotomd replaces each directive with the
corresponding generated content, enclosed in matching "Bgn" and "End"
markers.

## Actions

Each directive's ACTION determines what is included:

```html
<!---
gotomd::doc::./relativeDirectory/goObject
-->
```
Runs go doc on the specified object from the given directory.

The parameter format is: ./path/to/package/ObjectName

A special object name package includes the package-level comments.

```html
<!---
gotomd::docConstGrp::./relativeDirectory/goConstName ListOfConstNames
-->
```
Runs go doc on the specified constant block(s) from the given directory.

```html
<!---
gotomd::dcls::./relativeDirectory/declaredObject ListOfDeclaredGoObjects
-->
```

Inserts each listed declaration as a single line, regardless of
how it is declared in the source. No comments are included.

Example: functions, methods, constants.

```html
<!---
gotomd::dcln::./relativeDirectory/declaredObject ListOfDeclaredGoObjects
-->
```

Inserts each listed declaration exactly as in the source, including
any leading comments.

```html
<!---
gotomd::dcl::./relativeDirectory/declaredObject ListOfDeclaredGoObjects
-->
```

Inserts each listed declaration exactly as in the source, but without
comments.

```html
<!---
gotomd::tst::goTest::./relativeDirectory/testName
-->
```

Runs go test in the given directory, targeting the specified test(s)
or package, and includes the output.

```html
<!---
gotomd::file::./relativeDirectory/fName
-->
```

Inserts the contents of the specified file into a fenced code block.

```html
<!---
gotomd::run::./relativeDirectory [args ...]
-->
```

Runs go run on the package in the given directory (assumes main)
with the provided arguments, including the output.

## Output Markers

Generated content is wrapped between markers in the target file:

	const sztestBgnPrefix = sztestPrefix + "Bgn::"
	const sztestEndPrefix = sztestPrefix + "End::"

Additionally, an auto-generated section header is prefixed with:

	const szAutoPrefix = sztestPrefix + "Auto::"

This header is followed by a blank line. If operating in template mode
(not in-place), a DO NOT MODIFY warning is also inserted.

# Dedication

This project is dedicated to Reem.
Your brilliance, courage, and quiet strength continue to inspire me.
Every line is written in gratitude for the light and hope you brought into my
life.

NOTE: Documentation reviewed and polished with the assistance of ChatGPT from
OpenAI.
*/
package main

import (
	"fmt"
	"os"
)

const license = `
Golang To Github Markdown: gotomd.
Copyright (C) 2023, 2024  Leslie Dancsecs

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
`

func processFiles(filesToProcess []string) error {
	var err error

	for i, mi := 0, len(filesToProcess); i < mi && err == nil; i++ {
		clearPackageCache()

		switch {
		case cleanOnly:
			err = cleanMD(filesToProcess[i])
		case replace:
			err = replaceMDInPlace(filesToProcess[i])
		default:
			err = expandMD(filesToProcess[i])
		}
	}

	return err
}

func main() {
	var (
		err            error
		origWd         string
		filesToProcess []string
		usage          string
	)

	// Restore original working directory on exit.
	origWd, err = os.Getwd()
	if err == nil {
		defer func() {
			_ = os.Chdir(origWd)
		}()
	}

	filesToProcess, usage, err = processArgs()

	if showLicense {
		fmt.Print(license) //nolint:forbidigo  // Ok.
	}

	if showHelp {
		fmt.Println(usage) //nolint:forbidigo  // Ok.
	}

	if err == nil && buildUsage != "" {
		err = usageUpdate(buildUsage)
	}

	if err == nil {
		filesToProcess, err = getFilesToProcess(filesToProcess)
	}

	if err == nil {
		err = processFiles(filesToProcess)
	}

	if err != nil {
		panic(err)
	}
}
