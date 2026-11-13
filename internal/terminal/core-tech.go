package terminal

import (
	"strings"

	"github.com/qean32/holo/internal/constants/sgr"
	"github.com/qean32/holo/internal/lib/utils"
	"github.com/qean32/holo/internal/model"
)

func PrintTechInfo(event model.Event) {
	Println(utils.GetCustomMessage("Технический вывод:", sgr.UNDERLINE, sgr.GREEN) + " " + formatEvent(event))
}

func formatEvent(event model.Event) string {
	parts := []string{
		formatField("Key", event.Key),
		formatField("KeyWord", event.KeyWord),
		formatField("Payload", event.Payload),
		formatField("Password", event.Password),
		formatField("DateTime", event.DateTime),
		"Flags=" + formatFlags(event.Flags),
		"SubEvent=" + formatSubEvent(event.SubEvent),
	}
	return strings.Join(parts, " ")
}

func formatField(name, value string) string {
	label := utils.GetCustomMessage(name+"=", sgr.DIM)
	if value == "" {
		return label + utils.GetCustomMessage("<empty>", sgr.DIM)
	}
	return label + utils.GetCustomMessage(value, sgr.GREEN)
}

func formatFlags(flags []string) string {
	if len(flags) == 0 {
		return utils.GetCustomMessage("<empty>", sgr.DIM)
	}
	parts := make([]string, len(flags))
	for i, flag := range flags {
		parts[i] = utils.GetCustomMessage(flag, sgr.GREEN)
	}
	sep := utils.GetCustomMessage(",", sgr.DIM)
	return utils.GetCustomMessage("[", sgr.DIM) + strings.Join(parts, sep) + utils.GetCustomMessage("]", sgr.DIM)
}

func formatSubEvent(sub *model.Event) string {
	if sub == nil {
		return utils.GetCustomMessage("<nil>", sgr.DIM)
	}
	return utils.GetCustomMessage("{", sgr.DIM) + formatEvent(*sub) + utils.GetCustomMessage("}", sgr.DIM)
}
