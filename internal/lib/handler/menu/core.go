package menu

import (
	"path/filepath"

	"github.com/qean32/holo/internal/constants"
	errorlist "github.com/qean32/holo/internal/constants/error-list"
	eventlist "github.com/qean32/holo/internal/constants/event-list"
	"github.com/qean32/holo/internal/constants/literal"
	"github.com/qean32/holo/internal/lib/array"
	"github.com/qean32/holo/internal/lib/death"
	"github.com/qean32/holo/internal/lib/global"
	"github.com/qean32/holo/internal/lib/utils"
	"github.com/qean32/holo/internal/model"
	"github.com/qean32/holo/internal/terminal/list"
)

func MainMenu() {
	list.List(model.ListPayload{
		Options: List(),
		Title:   "Главное меню: ",
	})
}

func ScriptList(keyEvent string) error {
	return getScriptList(literal.Titles.ListScript,
		func(value string) model.Event {
			return model.Event{
				Key:     keyEvent,
				KeyWord: utils.TrimBatExtension(value),
			}
		},
	)
}

func getScriptList(title string, makeEvent func(name string) model.Event) error {
	dir := filepath.Join(constants.Root, constants.Script)
	commands, err := death.ListFilesByExt(dir, literal.Extension.Script)
	if err != nil {
		return errorlist.Read(dir, err)
	}

	list.List(model.ListPayload{
		Options: array.Mapped(commands,
			func(value string) model.Option {
				return model.Option{
					Message: utils.TrimBatExtension(value),
					Event:   makeEvent(value),
				}
			}),
		Title: title,
	})
	return nil
}

func List() []model.Option {
	items := []model.Option{
		{Message: "Ручной ввод", Event: model.Event{Key: eventlist.MANUAL}},
		{Message: "Помощь", Event: model.Event{Key: eventlist.HELP}},
		{Message: "Генерация мастер ключа", Event: model.Event{Key: eventlist.INWORK}},
		{Message: "Выход", Event: model.Event{Key: eventlist.STOP}},
	}

	if global.ScriptEnabled() {
		scriptItems := []model.Option{
			{Message: "Запустить команду", Event: model.Event{Key: eventlist.LISTSCRIPT}},
			{Message: "Добавить команду", Event: model.Event{Key: eventlist.QCREATESCRIPT}},
			{Message: "Удалить команду", Event: model.Event{Key: eventlist.LISTSCRIPTREMOVE}},
		}
		items = append(scriptItems, items...)
	}

	return items
}
