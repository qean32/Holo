package questionnaire

import (
	"slices"

	"github.com/qean32/holo/internal/constants"
)

func IsAcceptableYea(answer string) bool {
	return slices.Contains(constants.AcceptableYeaOrNot, answer)
}
