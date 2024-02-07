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
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

func confirmOverwrite(fPath string, data string) (bool, error) {
	var ok bool
	var oldData []byte
	_, err := os.Stat(fPath)
	if errors.Is(err, os.ErrNotExist) {
		return true, nil
	}
	if err == nil {
		oldData, err = os.ReadFile(fPath) //nolint:gosec // Ok.
	}
	if err == nil && strings.TrimRight(string(oldData), "\n") == data {
		fmt.Println("No change: " + fPath)
		return false, nil
	}
	if err == nil {
		fmt.Print("Confirm overwrite of ", fPath, " (Y to overwrite)? ")
		var response string
		if _, err = fmt.Scanln(&response); err == nil {
			ok = response == "Y"
			if !ok {
				fmt.Println("overwrite cancelled")
			}
		}
	}
	return ok, err
}

func writeFile(fPath string, data string) error {
	var err error

	data = strings.ReplaceAll(data, "\t", "    ")
	okToOverwrite := forceOverwrite
	if !okToOverwrite {
		okToOverwrite, err = confirmOverwrite(fPath, data)
	}

	if err == nil && okToOverwrite {
		var f *os.File

		//nolint:gosec // Ok.
		f, err = os.OpenFile(fPath,
			os.O_TRUNC|os.O_WRONLY|os.O_CREATE,
			os.FileMode(defaultPerm),
		)
		if err == nil {
			_, err = f.WriteString(data + "\n")
			if err == nil {
				err = f.Close()
			}
		}
	}
	return err
}

func getFilesToProcess() ([]string, error) {
	var err error
	var files []os.DirEntry
	var stat os.FileInfo
	var filesToProcess []string
	var filter = ".md"

	if !cleanOnly && !replace {
		filter += ".gtm"
	}

	for i, mi := 0, flag.NArg(); i < mi && err == nil; i++ {
		stat, err = os.Stat(flag.Arg(i))
		if err == nil && stat.IsDir() {
			files, err = os.ReadDir(flag.Arg(i))
			for j, mj := 0, len(files); j < mj && err == nil; j++ {
				fName := files[j].Name()
				if strings.HasSuffix(fName, filter) {
					filesToProcess = append(filesToProcess,
						flag.Arg(i)+string(os.PathSeparator)+fName,
					)
					if verbose {
						fmt.Println("filesToProcess: ",
							flag.Arg(i)+string(os.PathSeparator)+fName,
						)
					}
				}
			}
		}

		if err == nil && !stat.IsDir() {
			if !strings.HasSuffix(stat.Name(), filter) {
				err = errors.New("file must have extension: " + filter)
			} else {
				filesToProcess = append(filesToProcess, flag.Arg(i))
				if verbose {
					fmt.Println("filesToProcess: ", flag.Arg(i))
				}
			}
		}
	}
	return filesToProcess, err
}
