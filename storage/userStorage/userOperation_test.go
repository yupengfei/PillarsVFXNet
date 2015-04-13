package userStorage

import (
	"PillarsPhenomVFXWeb/utility"
	"fmt"
	"testing"
)

func Test_InsertIntoUser(t *testing.T) {
	temp := "testing"
	userCode := utility.GenerateCode(&temp)
	password := utility.GenerateCode(&temp)
	displayName := utility.GenerateCode(&temp)
	picture := utility.GenerateCode(&temp)
	email := "yupengfei@foxmail.com"
	phone := "18053182006"

	user := utility.User{
		UserCode:    *userCode,
		Password:    *password,
		DisplayName: *displayName,
		Picture:     *picture,
		Email:       email,
		Phone:       phone,
		Status:      0,
	}

	result, err := InsertIntoUser(&user)
	if result == false {
		fmt.Println(err.Error())
		t.Error("Inert into user failed")
	} else {
		fmt.Println(*utility.ObjectToJsonString(result))
	}
}

func Test_DeleteUserByEmail(t *testing.T) {
	email := "yupengfei@foxmail.com"
	result, err := DeleteUserByEmail(&email)
	if err != nil {
		t.Error("delete user failed")
	} else {
		fmt.Println(*utility.ObjectToJsonString(result))
	}
}

func Test_QueryUserByEmail(t *testing.T) {
	email := "yupengfei@foxmail.com"
	_, err := QueryUserByEmail(&email)
	if err != nil {
		t.Error("query user by emial failed")
	}
}
