package entities

type Language struct {
	Code string
	Name string
}

type AllLanguages []Language

var Languages = AllLanguages{
	{
		Code:    "code1",
		Name:    "Language",
	},
}

func NewLanguage() Language {
	return Language{"", ""}
}

func NewLanguageParam(code string, name string) Language {
	return Language{code, name}
}

func StringLanguage(language Language) string {
	return " Code : " + language.Code + " Name : " + language.Name
}
