package models

import "testing"

func NewUser(username string, age int64, gender bool, nationality string) *CreateUserCMD {
	return &CreateUserCMD{
		Username:    username,
		Age:         age,
		Gender:      gender,
		Nationality: nationality,
	}
}

func Test_ValidateCreateUserWithCorrectParams(t *testing.T) {
	u := NewUser("Ramiro", 21, true, "Argentina")

	err := u.validate()

	if err != nil {
		t.Error("Failed the test at the user creation")
		t.Fail()
	} else {
		t.Log("Success")
	}
}

func Test_ValidateCreateUserWithWrongUsername(t *testing.T) {
	u := NewUser("", 21, true, "Argentina")

	err := u.validate()

	if err == nil {
		t.Error("Fail Test. User was created but with an invalid parameter.")
		t.Fail()
	} else {
		t.Log("Success")
	}
}

func Test_ValidateCreateUserWithWrongAge(t *testing.T) {
	u := NewUser("Ramiro", 0, true, "Argentina")

	err := u.validate()

	if err == nil {
		t.Error("Fail Test. User can not have a negative age.")
		t.Fail()
	} else {
		t.Log("Success")
	}
}
