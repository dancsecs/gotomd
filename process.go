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
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/dancsecs/szlog"
)

// Overwrite messages.
const (
	confirmMsg = "Confirm overwrite of: %s\n" +
		" (Y to overwrite; [N/n] to Cancel; [R/r] to review)? "
	confirmCancelled = "Overwrite cancelled\n\n"
	confirmUnknown   = "Unknown response: '%s'\n\n"
)

func askToOverwrite(fPath, oldData, newData string) (bool, error) {
	var (
		ok  bool
		err error
	)

	response := "R"
	for err == nil && response == "R" {
		szlog.Say0f(confirmMsg, fPath)

		if _, err = fmt.Scanln(&response); err == nil {
			switch response {
			case "Y":
				ok = true
			case "N", "n":
				szlog.Say0(confirmCancelled)
			case "R", "r":
				diffFile(filepath.Base(fPath), oldData, newData+"\n")

				response = "R"
			default:
				szlog.Say0f(confirmUnknown, response)
				response = "R" // Ask again.
			}
		}
	}

	return ok, err //nolint:wrapcheck // Caller will wrap error.
}

func confirmOverwrite(fPath string, data string) (bool, error) {
	var oldData []byte

	_, err := os.Stat(fPath)
	if errors.Is(err, os.ErrNotExist) {
		return true, nil
	}

	if err == nil {
		oldData, err = os.ReadFile(fPath) //nolint:gosec // Ok.
	}

	if err == nil && strings.TrimRight(string(oldData), "\n") == data {
		szlog.Say0("No change: ", fPath, "\n")

		return false, nil
	}

	return askToOverwrite(fPath, string(oldData)+"\n", data+"\n")
}

func writeFile(fPath string, data string) error {
	var err error

	data = strings.ReplaceAll(data, "\t", "    ")
	okToOverwrite := forceOverwrite

	if !okToOverwrite {
		okToOverwrite, err = confirmOverwrite(fPath, data)
	}

	if err == nil && okToOverwrite {
		var file *os.File

		//nolint:gosec // Ok.
		file, err = os.OpenFile(fPath,
			os.O_TRUNC|os.O_WRONLY|os.O_CREATE,
			os.FileMode(defaultPerm),
		)
		if err == nil {
			_, err = file.WriteString(data + "\n")
			if err == nil {
				err = file.Close()
			}
		}
	}

	return err //nolint:wrapcheck // Caller will wrap error.
}

//nolint:cyclop // Ok.
func getFilesToProcess(rawFilesToProcess []string) ([]string, error) {
	var (
		err            error
		files          []os.DirEntry
		stat           os.FileInfo
		filesToProcess []string
		filter         = ".md"
	)

	if !cleanOnly && !replace {
		filter = ".gtm" + filter
	}

	addFileToProcess := func(newFileToProcess string) {
		filesToProcess = append(filesToProcess, newFileToProcess)

		szlog.Info("filesToProcess: ", newFileToProcess)
	}

	for i, mi := 0, len(rawFilesToProcess); i < mi && err == nil; i++ {
		argPath := filepath.Clean(rawFilesToProcess[i])

		stat, err = os.Stat(argPath)
		if err == nil { //nolint:nestif // Ok file or dir logic.
			if stat.IsDir() {
				// Process a directory for all matching extensions.
				files, err = os.ReadDir(argPath)
				for j, mj := 0, len(files); j < mj && err == nil; j++ {
					fName := files[j].Name()
					if strings.HasSuffix(fName, filter) {
						addFileToProcess(filepath.Join(argPath, fName))
					}
				}
			} else {
				// Process the file (if it has the correct extension).
				if strings.HasSuffix(stat.Name(), filter) {
					addFileToProcess(filepath.Clean(argPath))
				} else {
					err = fmt.Errorf(
						"%w: expected - %s",
						ErrUnexpectedExtension, filter,
					)
				}
			}
		}
	}

	return filesToProcess, err
}
