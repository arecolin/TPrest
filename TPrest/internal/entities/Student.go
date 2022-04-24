package entities

import "strconv"

type Student struct {
	Id           string
	FirstName    string
	LastName     string
	Age          int
	LanguageCode string
}

type AllStudents []Student

var Students = AllStudents{
	{
		Id:           "1",
		FirstName:    "Jean",
		LastName:     "Jacques",
		Age:          20,
		LanguageCode: "code1",
	},
}

func NewStudent() Student {
	return Student{"0", "", "", 0, ""}
}

func NewStudentParam(id string, fn string, ln string, age int, lc string) Student {
	return Student{id, fn, ln, age, lc}
}

func StringStudent(student Student) string {
	return "id : " + student.Id + " FirstName : " + student.FirstName + " LastName : " + student.LastName + " Age : " + strconv.Itoa(student.Age) + " LanguageCode : " + student.LanguageCode
}
