package field

import (
	"bufio"
	"os"
	"strings"

	"main/lib/utils"
	"main/model"
	"main/terminal"
)

var reader = bufio.NewReader(os.Stdin)

func Field(payload model.FieldPayload) string {
	terminal.Print(utils.GetCustomMessage(payload.Prefix, payload.PrefixSRG...))
	text, _ := reader.ReadString('\n')

	if payload.Transfer {
		terminal.DownAndStart()
	}
	return strings.TrimSpace(text)
}
