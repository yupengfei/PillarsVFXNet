package loginAction

import (
	"PillarsPhenomVFXWeb/session"
	us "PillarsPhenomVFXWeb/storage/userStorage"
	u "PillarsPhenomVFXWeb/utility"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	olen := len(r.Form["Password"]) + len(r.Form["Email"])
	if olen != 2 {
		u.OutputJson(w, 1, "Error parameter format", nil)
		return
	}

	if len(r.Form["Password"][0]) == 0 {
		u.OutputJson(w, 12, "Error parameter Password", nil)
		return
	}

	if len(r.Form["Email"][0]) == 0 {
		u.OutputJson(w, 13, "Error parameter Email", nil)
		return
	}

	e := r.Form["Email"][0]
	p := r.Form["Password"][0]
	user, err := us.CheckEmailAndPassword(&(e), &(p))
	if err != nil {
		u.OutputJson(w, 14, "Email or Password wrong!", nil)
		return
	}

	if user.UserCode == "" {
		u.OutputJson(w, 15, "Login failed!", nil)
		return
	}

	// 登陆成功，用户信息放入session
	userSession := session.GlobalSessions.SessionStart(w, r)
	userSession.Set("userCode", user.UserCode)
	userSession.Set("errorTimes", 0)
	userSession.Set("loginTime", time.Now().Unix())
	userSession.Set("lastAction", time.Now().Unix())

	// 根据用户权限类型，跳转不同页面
	if user.UserAuthority == "admin" {
		u.OutputJson(w, 0, "user_list", nil)
	} else if user.UserAuthority == "制片" {
		u.OutputJson(w, 0, "project_list", nil)
	} else if user.UserAuthority == "制片助理" {
		u.OutputJson(w, 0, "project_list", nil)
	} else if user.UserAuthority == "分包商" {
		u.OutputJson(w, 0, "vendor.html", nil)
	}

}
