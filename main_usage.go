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

---

gotomd
Golang to 'github' markdown.

Usage: gotomd [-v | --verbose ...] [-c | --clean] [-r | --replace]
    [-l | --license] [-h | --help] [-f | --force] [-z | --colorize]
    [-o | --output dir] [-p | --permission perm] [path ...]

  - [-v | --verbose ...]</br>
    Provide more information when processing.


  - [-c | --clean]</br>
    Reverse operation and remove generated markdown (Cannot be used with the
    [-r | --replace] option).


  - [-r | --replace]</br>
    Replace the *.MD in place (Cannot be used with the [-c | --clean] option).


  - [-l | --license]</br>
    Display license before program exits.


  - [-h | --help]</br>
    Display program usage information.


  - [-f | --force]</br>
    Do not confirm overwrite of destination.


  - [-z | --colorize]</br>
    Colorize go test output.


  - [-o | --output dir]</br>
    Direct all output to the specified directory.


  - [-p | --permission perm]</br>
    Permissions to use when creating new file (can only set RW bits).


  - [path ...]</br>
    A specific gotomd file template with the extension '*.gtm.md' or a directory
    which will be searched for all matching template '*.gtm.md' files.   It
    defaults to the current directory: '.'
*/
package main
