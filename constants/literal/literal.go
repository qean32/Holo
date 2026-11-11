package literal

var Flags = flagType{
	NOLOG: "-nl",
}

var Extension = extensionType{
	Script: ".bat",
	Txt:    ".txt",
}

var Path = pathType{
	PathLog: "log" + Extension.Txt,
}

var Titles = titleType{
	List:        "Меню",
	ListScript:  "Список команд",
	EnterScript: "Добавить команду",
	Access:      "Подтвердите действие (yea/no)",
}
