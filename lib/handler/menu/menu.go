package menu

import (
	"path/filepath"

	"main/constants"
	errorlist "main/constants/error-list"
	eventlist "main/constants/event-list"
	"main/constants/literal"
	"main/lib/array"
	"main/lib/death"
	"main/lib/global"
	"main/lib/utils"
	"main/model"
	"main/terminal/list"
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
