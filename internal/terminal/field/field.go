package field

import (
	"bufio"
	"os"
	"strings"

	"github.com/qean32/holo/internal/lib/utils"
	"github.com/qean32/holo/internal/model"
	"github.com/qean32/holo/internal/terminal"
)

var reader = bufio.NewReader(os.Stdin)

func Field(payload model.FieldPayload) string {
	terminal.Print(utils.GetCustomMessage(payload.Prefix, payload.PrefixSRG...))
	text, _ := reader.ReadString('\n')

	if payload.Transfer {
		terminal.DownAndStart()
	}
	return strings.ToLower(strings.TrimSpace(text))
}
