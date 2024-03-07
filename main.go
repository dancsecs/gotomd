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

//nolint:lll // Ok.
/*
The gotomd utility provides for the maintenance of github README.MD style
pages by permitting go files, go documentation and go test output to be
included by reference into the github README.md file directly from the Go
code permitting program documentation to be maintained in one place (the Go
code.)

It can use a template file (```*.md.gtm```) or can maintain a ```*.md``` file
in place.

Usage of gotomd [-c | -r] [-fvl] [-p perm] [-o outDir] [-U file] [-u uint] path [path...]

The flags are:

  -c
      Reverse operation and remove generated markdown (Cannot be used
      with the -r option). Files with a .md extension are expected with
      a an .md.gtm file being produced.
  -f
      Do not confirm overwrite of destination.
  -l
      Display license before program exits.
  -o string
      Direct all output to the specified directory. (default ".")
  -p int
      Permissions to use when creating new file (can only set RW
      bits). (default 420)
  -r
      Replace the *.MD in place (Cannot be used with the -c flag).
  -U  string
      Collect cpu profile data into named file.
  -u  uint
      Number of iterations to run when collecting cpu profile information.
  -v
      Provide more information when processing.

Directives are placed into the ```*.md.gtm``` file (or directly into the
```*.md``` document if the replace in place option is given.  These directives
are embedded into HTML style comments.

```html
<!--- gotomd::ACTION::PARAMETERS -->
```

where ACTION can be one of the following:

- gotomd::doc::./relativeDirectory/goObject

    Run the go doc command on the object listed from the directory
    specified.  The PARAMETER is formatted with the relative directory up
    to the last directory separator before the end of the string and the
    desired object.  A special object package returns the package
    comments.

- gotomd::dcls::./relativeDirectory/declaredObject ListOfDeclaredGoObjects

    Pull out the declaration for the object and include as a single-line
    regardless of how declared in the go code.  The Parameter is a list of
    go functions, methods and constants (more object coming) to be included
    in a go code block. No comments are included.

- gotomd::dcln::./relativeDirectory/declaredObject ListOfDeclaredGoObjects

    Pull the declaration and include exactly as declared in the go
    source including leading comments.

- gotomd::dcl::./relativeDirectory/declaredObject ListOfDeclaredGoObjects

    Pull the declaration and include exactly as declared in the go
    source.  No Comments are included.

- gotomd::tst::goTest::./relativeDirectory/testName

    Run go test with the tests listed (or package) to run all tests and
    included the output.

- gotomd::file::./relativeDirectory/fName

    Include the specified file in a code block.


When expanded in the target file the content will be framed by similar
comments prefixed with 'Bgn' and 'End' as:

const sztestBgnPrefix = sztestPrefix + "Bgn::"
const sztestEndPrefix = sztestPrefix + "End::"

A header prefixed with

const szAutoPrefix = sztestPrefix + "Auto::"

and a blank line following will be inserted into the output file.  If the
action is not "replace in place" then an addition **DO NOT MODIFY**
warning is included.
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

func main() {
	var (
		err            error
		origWd         string
		filesToProcess []string
	)

	// Restore original working directory on exit.
	origWd, err = os.Getwd()
	if err == nil {
		defer func() {
			_ = os.Chdir(origWd)
		}()
	}

	processArgs()

	if showLicense {
		fmt.Print(license)
	}

	filesToProcess, err = getFilesToProcess()

	for i, mi := 0, len(filesToProcess); i < mi && err == nil; i++ {
		switch {
		case cleanOnly:
			//   err = cleanMD(filesToProcess[i])
			err = cleanMD(filesToProcess[i])
		case replace:
			err = replaceMDInPlace(filesToProcess[i])
		default:
			err = expandMD(filesToProcess[i])
		}
	}

	if err != nil {
		panic(err)
	}
}
