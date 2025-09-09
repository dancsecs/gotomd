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
	"strings"
)

func markGoCode(content string) string {
	return "```go\n" + strings.TrimRight(content, "\n") + "\n```"
}

func markBashCode(content string) string {
	return "```bash\n" + strings.TrimRight(content, "\n") + "\n```"
}

func getDoc(cmd string) (string, error) {
	var (
		dInfo *docInfo
		res   string
	)

	dir, action, err := parseCmds(cmd)
	for i, mi := 0, len(dir); i < mi && err == nil; i++ {
		dInfo, err = getInfo(dir[i], action[i])
		if err == nil {
			if res != "" {
				res += "\n\n"
			}

			res += dInfo.declGoLang() + "\n\n" +
				dInfo.docMarkdown()
		}
	}

	if err == nil {
		return res, nil
	}

	return "", err
}

func getDocDecl(cmd string) (string, error) {
	var (
		dInfo *docInfo
		res   string
	)

	dir, action, err := parseCmds(cmd)
	if err == nil {
		for i, mi := 0, len(dir); i < mi && err == nil; i++ {
			dInfo, err = getInfo(dir[i], action[i])
			if err == nil {
				if res != "" {
					res += "\n"
				}

				res += strings.Join(dInfo.header, "\n")
			}
		}
	}

	if err == nil {
		return markGoCode(res), nil
	}

	return "", err
}

func getDocDeclSingle(cmd string) (string, error) {
	var (
		dInfo *docInfo
		res   string
	)

	dir, action, err := parseCmds(cmd)
	if err == nil {
		for i, mi := 0, len(dir); i < mi && err == nil; i++ {
			dInfo, err = getInfo(dir[i], action[i])
			if err == nil {
				if res != "" {
					res += "\n"
				}

				res += dInfo.oneLine()
			}
		}
	}

	if err == nil {
		return markGoCode(res), nil
	}

	return "", err
}

func getDocDeclNatural(cmd string) (string, error) {
	var (
		dInfo *docInfo
		res   string
	)

	dir, action, err := parseCmds(cmd)
	if err == nil {
		for i, mi := 0, len(dir); i < mi && err == nil; i++ {
			dInfo, err = getInfo(dir[i], action[i])
			if err == nil {
				if res != "" {
					res += "\n\n"
				}

				res += dInfo.naturalComments() + "\n"
				res += dInfo.oneLine()
			}
		}
	}

	if err == nil {
		return markGoCode(res), nil
	}

	return "", err
}

func getDocDeclConstantBlock(cmd string) (string, error) {
	var (
		dInfo *docInfo
		res   string
	)

	dir, action, err := parseCmds(cmd)
	if err == nil {
		for i, mi := 0, len(dir); i < mi && err == nil; i++ {
			dInfo, err = getInfo(dir[i], action[i])
			if err == nil {
				if res != "" {
					res += "\n\n"
				}

				res += dInfo.constantBlock() + "\n"
			}
		}
	}

	if err == nil {
		return markGoCode(res), nil
	}

	return "", err
}
