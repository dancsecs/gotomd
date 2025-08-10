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

package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"unicode"

	"github.com/dancsecs/szlog"
)

func updatePackageLine(current, i int, lines []string) int {
	if strings.HasPrefix(lines[i], "package ") {
		if current != -1 {
			szlog.Warn("multiple package delimiters.")
		}

		current = i // We have a new package section.
	}

	return current
}

func updateUsageLine(
	current int, usageName string, i int, lines []string,
) (int, string) {
	newUsageLine := -1
	numLines := len(lines)

	re := regexp.MustCompile(`^\# Usage: (.*)$`)

	if lines[i] == "/*" && (numLines-i) > 1 {
		matches := re.FindStringSubmatch(lines[i+1])
		if len(matches) == 2 { //nolint:mnd // Length if found.
			newUsageLine = i
			usageName = matches[1]

			for newUsageLine > 0 && lines[newUsageLine-1] == "" {
				newUsageLine--
			}
		}
	}

	if newUsageLine != -1 {
		if current != -1 {
			szlog.Warn("multiple Usage delimiters.")
		}

		current = newUsageLine // We have n new usage section.
	}

	return current, usageName
}

func findDelimiters(lines []string) (int, int, string) {
	packageLine := -1
	usageLine := -1
	usageName := ""

	for i := range lines {
		packageLine = updatePackageLine(packageLine, i, lines)
		usageLine, usageName = updateUsageLine(usageLine, usageName, i, lines)
	}

	return packageLine, usageLine, usageName
}

func parseOldFile(
	usagePackageFile string,
) ([]string, int, string, int, error) {
	oldFileData, err := os.ReadFile(usagePackageFile) //nolint:gosec // Ok.
	if err != nil {
		return nil, -1, "", -1, err //nolint:wrapcheck // Ok.
	}

	oldLines := []string(nil)

	if len(oldFileData) < 1 {
		szlog.Warnf("blank usage file: '%s'", usagePackageFile)
	} else {
		oldLines = strings.Split(
			strings.TrimLeftFunc(
				string(oldFileData),
				unicode.IsSpace,
			),
			"\n",
		)

		for i, l := range oldLines {
			oldLines[i] = strings.TrimRightFunc(l, unicode.IsSpace)
		}

		if len(oldLines) == 1 && oldLines[0] == "" {
			oldLines = nil
		}
	}

	packageLine, usageLine, usageName := findDelimiters(oldLines)

	if usageLine == -1 {
		szlog.Warnf("no previous usage found in: '%s'", usagePackageFile)
	}

	if packageLine == -1 {
		szlog.Warnf("package header not found in: '%s'", usagePackageFile)
	}

	return oldLines, usageLine, usageName, packageLine, nil
}

func preUsageLines(usageLine, packageLine int, lines []string) []string {
	mi := len(lines)
	if usageLine > -1 && usageLine < mi {
		mi = usageLine
	}

	if packageLine > -1 && packageLine < mi {
		mi = packageLine
	}

	return lines[:mi]
}

func parseNewUsage(isFirst bool, usageName string, lines []string) []string {
	newLines := make([]string, 0, len(lines)+4) //nolint:mnd // Extra lines.

	if !isFirst {
		newLines = append(newLines, "")
	}

	// Add new lines
	newLines = append(newLines,
		"/*",
		"# Usage: "+usageName,
	)
	newLines = append(newLines, lines...)
	newLines = append(newLines, "*/")

	return newLines
}

func usageUpdate(usgPkgFile string) error {
	var (
		oldLines       []string
		usageStdinData []byte
		newLines       []string
		packageLine    int
		usageLine      int
		usageName      string
		err            error
	)

	oldLines, usageLine, usageName, packageLine, err = parseOldFile(usgPkgFile)

	if err == nil {
		newLines = preUsageLines(usageLine, packageLine, oldLines)

		// Now read replacement usage from os.Stdin.
		usageStdinData, err = io.ReadAll(os.Stdin)
	}

	if err == nil {
		newLines = append(
			newLines,
			parseNewUsage(
				len(newLines) == 0, // Is First line.
				usageName,
				strings.Split(string(usageStdinData), "\n"),
			)...,
		)

		if packageLine > -1 {
			for i, mi := packageLine, len(oldLines); i < mi; i++ {
				newLines = append(newLines, oldLines[i])
			}
		}
	}

	if err == nil {
		err = os.WriteFile(
			usgPkgFile,
			[]byte(strings.Join(newLines, "\n")),
			0, // Use existing permissions.
		)
	}

	if err == nil {
		return nil
	}

	return fmt.Errorf("could not read file: %w", err)
}
