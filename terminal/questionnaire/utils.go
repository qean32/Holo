package questionnaire

import (
	"slices"

	"main/constants"
)

func IsAcceptableYea(answer string) bool {
	return slices.Contains(constants.AcceptableYeaOrNot, answer)
}
