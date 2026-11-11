package handler

import (
	"path/filepath"
	"strings"

	"atomicgo.dev/cursor"

	"main/constants"
	errorlist "main/constants/error-list"
	"main/constants/literal"
	"main/terminal"
)

func validateName(name string) error {
	if name == "" || name == "." || name == ".." {
		return errorlist.ErrForbiddenName
	}
	if strings.ContainsAny(name, `/\:*?"<>|`) {
		return errorlist.ErrForbiddenChars
	}
	return nil
}

func scriptPath(keyword string) string {
	return filepath.Join(constants.Script, keyword+literal.Extension.Script)
}

func Help() {
	cursor.StartOfLine()
	terminal.PrintASCIICenter(constants.HelpMessage, "")
}
