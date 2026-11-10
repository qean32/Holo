package handler

import (
	"atomicgo.dev/cursor"

	"main/constants"
	eventlist "main/constants/event-list"
	"main/constants/literal"
	"main/lib/global"
	"main/model"
	"main/terminal"
	"main/terminal/list"
)

func Menu() []model.Option {
	items := []model.Option{
		{Message: "Ручной ввод", Event: model.Event{Key: eventlist.MANUAL}},
		{Message: "Помощь", Event: model.Event{Key: eventlist.HELP}},
		{Message: "Генерация мастер ключа", Event: model.Event{Key: eventlist.INWORK}},
		{Message: "Выход", Event: model.Event{Key: eventlist.STOP}},
	}

	if global.ScriptEnabled() {
		scriptItems := []model.Option{
			{Message: "Запустить команду", Event: model.Event{Key: eventlist.LISTSCRIPT}},
			{Message: "Добавить команду", Event: model.Event{Key: eventlist.MANUCREATESCRIPT}},
			{Message: "Удалить команду", Event: model.Event{Key: eventlist.LISTSCRIPTREMOVE}},
		}
		items = append(scriptItems, items...)
	}

	return items
}

func runMenu() {
	list.List(model.ListPayload{
		Options: Menu(),
		Title:   literal.Titles.Menu,
	})
}

func help() {
	cursor.StartOfLine()
	terminal.PrintASCIICenter(constants.HelpMessage, "")
}
