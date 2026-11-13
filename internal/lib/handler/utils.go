package handler

import (
	"path/filepath"
	"strings"

	"atomicgo.dev/cursor"

	"github.com/qean32/holo/internal/constants"
	errorlist "github.com/qean32/holo/internal/constants/error-list"
	"github.com/qean32/holo/internal/constants/literal"
	"github.com/qean32/holo/internal/terminal"
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
