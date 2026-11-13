package death

import (
	"os"
	"os/exec"
	"slices"
	"strings"
	"time"

	"github.com/qean32/holo/internal/callstack"
	errorlist "github.com/qean32/holo/internal/constants/error-list"
	eventlist "github.com/qean32/holo/internal/constants/event-list"
	"github.com/qean32/holo/internal/constants/literal"
	"github.com/qean32/holo/internal/model"
)

func Logger(stackErr model.StackError) error {
	if stackErr.Event.Key == eventlist.RESPONSE {
		return nil
	}
	if noLog := slices.ContainsFunc(stackErr.Event.Flags, func(item string) bool {
		return strings.TrimSpace(item) == literal.Flags.NOLOG
	}); noLog {
		return nil
	}

	return PushToFile(literal.Path.PathLog, callstack.StackErrorToString(stackErr, 0))
}

func CurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func RunScript(command string) error {
	if err := exec.Command("cmd", "/C", command).Run(); err != nil {
		return errorlist.RunCommand(command, err)
	}
	return nil
}

func Exit() {
	ClearTerminal()
	os.Exit(0)
}

func ClearTerminal() error {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return errorlist.ClearTerminal(err)
	}
	return nil
}
