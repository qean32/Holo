package literal

var Flags = flagType{
	NOLOG: "-nl",
}

var Extension = extensionType{
	Script: ".bat",
	Txt:    ".txt",
}

var Path = pathType{
	PathLog: "log.txt",
}

var Titles = titleType{
	Menu:        "Меню",
	ListScript:  "Список сценариев",
	EnterScript: "Добавить сценарий",
	Access:      "Подтвердите действие (yea/no)",
}
