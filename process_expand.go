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
	"bytes"
	"os"
	"path/filepath"
	"strings"

	"github.com/dancsecs/szlog"
)

func expandMD(rPath string) error {
	var (
		err         error
		rDir, rFile string
		wDir, wFile string
		wPath       string
		fileBytes   []byte
		res         string
	)

	rDir, rFile = filepath.Split(rPath)
	wDir = rDir

	if outputDir != "." {
		wDir = outputDir
	}

	wFile = strings.TrimSuffix(
		strings.TrimPrefix(rFile, "."),
		".gtm.md",
	) + ".md"
	wPath = filepath.Join(wDir, wFile)

	szlog.Infof("Expanding %s to: %s", rPath, wPath)

	fileBytes, err = os.ReadFile(rPath) //nolint:gosec // Ok.

	if err == nil {
		fileData := string(bytes.TrimRight(fileBytes, "\n"))
		res, err = updateMarkDownDocument(rDir, fileData)
	}

	if err == nil {
		err = writeFile(wPath, res)
	}

	return err
}
