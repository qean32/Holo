package literal

var Flags = flagType{
	NOLOG: "-nl",
}

var Extension = extensionType{
	Scenario: ".bat",
}

var Path = pathType{
	PathLog: "log.txt",
}

var Titles = titleType{
	Menu:          "Меню",
	ListScenario:  "Список сценариев",
	EnterScenario: "Добавить сценарий",
	Access:        "Подтвердите действие (yea/no)",
}
