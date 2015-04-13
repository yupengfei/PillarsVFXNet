package userAction

import (
	s "PillarsPhenomVFXWeb/session"
	us "PillarsPhenomVFXWeb/storage/userStorage"
	u "PillarsPhenomVFXWeb/utility"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	if !s.CheckAuthority(w, r, "") {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	r.ParseForm()
	olen := len(r.Form["Email"]) + len(r.Form["UserName"]) + len(r.Form["Phone"]) + len(r.Form["UserAuthority"]) + len(r.Form["FilePath"])
	if olen != 5 {
		u.OutputJson(w, 1, "Error parameter format", nil)
		return
	}

	if len(r.Form["Email"][0]) == 0 {
		u.OutputJson(w, 12, "Error parameter Email", nil)
		return
	}

	if len(r.Form["UserName"][0]) == 0 {
		u.OutputJson(w, 13, "Error parameter UserName", nil)
		return
	}

	if len(r.Form["Phone"][0]) == 0 {
		u.OutputJson(w, 14, "Error parameter Phone", nil)
		return
	}

	if len(r.Form["UserAuthority"][0]) == 0 {
		u.OutputJson(w, 15, "Error parameter UserAuthority", nil)
		return
	}

	if len(r.Form["FilePath"][0]) == 0 {
		u.OutputJson(w, 16, "Error parameter FilePath", nil)
		return
	}

	temp := "insert"
	userCode := u.GenerateCode(&temp)
	picture := u.GenerateCode(&temp)
	user := u.User{
		UserCode:      *userCode,
		Password:      "E10ADC3949BA59ABBE56E057F20F883E", // 默认为md5(123456, 32)
		DisplayName:   r.Form["UserName"][0],
		Picture:       *picture,
		Email:         r.Form["Email"][0],
		Phone:         r.Form["Phone"][0],
		UserAuthority: r.Form["UserAuthority"][0],
		FilePath:      r.Form["FilePath"][0],
		Status:        0,
	}
	result, _ := us.InsertIntoUser(&user)
	if result == false {
		u.OutputJson(w, 17, "Inert into user failed!", nil)
		return
	}

	UserList(w, r)
}

func UserList(w http.ResponseWriter, r *http.Request) {
	if !s.CheckAuthority(w, r, "") {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	list, err := template.ParseFiles("pages/userlist.gtpl")
	if err != nil {
		panic(err.Error())
	}
	userList, err := us.QueryUserList()
	if err != nil {
		panic(err.Error())
	}
	list.Execute(w, userList)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if !s.CheckAuthority(w, r, "") {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	r.ParseForm()
	olen := len(r.Form["Email"])
	if olen != 1 {
		u.OutputJson(w, 1, "Error parameter format", nil)
		return
	}

	if len(r.Form["Email"][0]) == 0 {
		u.OutputJson(w, 12, "Error parameter Email", nil)
		return
	}

	email := r.Form["Email"][0]
	result, _ := us.DeleteUserByEmail(&email)
	if result == false {
		u.OutputJson(w, 13, "Delete user failed!", nil)
		return
	}

	u.OutputJson(w, 0, "Delete user succeed!", nil)
}

func QueryUser(w http.ResponseWriter, r *http.Request) {
	if !s.CheckAuthority(w, r, "") {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	r.ParseForm()
	olen := len(r.Form["Email"])
	if olen != 1 {
		u.OutputJson(w, 1, "Error parameter format", nil)
		return
	}

	if len(r.Form["Email"][0]) == 0 {
		u.OutputJson(w, 12, "Error parameter Email", nil)
		return
	}

	email := r.Form["Email"][0]
	user, err := us.QueryUserByEmail(&email)
	if err != nil {
		u.OutputJson(w, 13, "Query user failed!", nil)
		return
	}

	result, _ := json.Marshal(user)
	fmt.Fprintf(w, string(result))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if !s.CheckAuthority(w, r, "") {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	r.ParseForm()
	olen := len(r.Form["Email"]) + len(r.Form["UserName"]) + len(r.Form["Phone"]) + len(r.Form["UserAuthority"]) + len(r.Form["FilePath"])
	if olen != 5 {
		u.OutputJson(w, 1, "Error parameter format", nil)
		return
	}

	if len(r.Form["Email"][0]) == 0 {
		u.OutputJson(w, 12, "Error parameter Email", nil)
		return
	}

	if len(r.Form["UserName"][0]) == 0 {
		u.OutputJson(w, 13, "Error parameter UserName", nil)
		return
	}

	if len(r.Form["Phone"][0]) == 0 {
		u.OutputJson(w, 14, "Error parameter Phone", nil)
		return
	}

	if len(r.Form["UserAuthority"][0]) == 0 {
		u.OutputJson(w, 15, "Error parameter UserAuthority", nil)
		return
	}

	if len(r.Form["FilePath"][0]) == 0 {
		u.OutputJson(w, 16, "Error parameter FilePath", nil)
		return
	}

	user := u.User{
		DisplayName:   r.Form["UserName"][0],
		Email:         r.Form["Email"][0],
		Phone:         r.Form["Phone"][0],
		UserAuthority: r.Form["UserAuthority"][0],
		FilePath:      r.Form["FilePath"][0],
	}
	result, _ := us.UpdateUserByEmail(&user)
	if result == false {
		u.OutputJson(w, 17, "Update user failed!", nil)
		return
	}

	UserList(w, r)
}

//查询外包商列表
func GetVendorList(w http.ResponseWriter, r *http.Request) {
	flag, _ := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		u.OutputJsonLog(w, 404, "session error!", nil, "")
		return
	}

	result, err := us.QueryVendorUserList()
	if err != nil || result == nil {
		u.OutputJsonLog(w, 1, "Query failed!", nil, "userAction.GetVendorList: userStorage.GetVendorList() failed!")
		return
	}

	u.OutputJsonLog(w, 0, "Query success.", result, "")
}
