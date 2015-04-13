package postAction

import (
	s "PillarsPhenomVFXWeb/session"
	"PillarsPhenomVFXWeb/storage/postStorage"
	u "PillarsPhenomVFXWeb/utility"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func QueryShotVersion(w http.ResponseWriter, r *http.Request) {
	flag, _ := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		u.OutputJsonLog(w, 404, "session error!", nil, "")
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.QueryShotVersion: ioutil.ReadAll(r.Body) failed!")
		return
	}
	var version u.ShotVersion
	json.Unmarshal(data, &version)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.QueryShotVersion: json.Unmarshal(data, &version) failed!")
		return
	}
	if len(version.ShotCode) == 0 {
		u.OutputJsonLog(w, 13, "Parameters Checked failed!", nil, "postAction.QueryShotVersion: Parameters Checked failed!")
		return
	}
	result, err := postStorage.QueryShotVersion(&version.ShotCode)
	if err != nil || result == nil {
		u.OutputJsonLog(w, 14, "Query failed!", nil, "postAction.QueryShotVersion: postStorage.QueryShotVersion(&ShotCode) failed!")
		return
	}

	u.OutputJsonLog(w, 0, "Query success.", result, "")
}

// 成品下载
func DownloadShotProduct(w http.ResponseWriter, r *http.Request) {
	flag, _ := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	r.ParseForm()
	if len(r.Form["VersionCode"]) == 0 {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}
	var versionCode = r.Form["VersionCode"][0]
	if len(versionCode) == 0 {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}
	//查询文件的路径
	sv, err := postStorage.QueryShotProduct(&versionCode)
	if err != nil || sv == nil {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}
	if checkFileIsExist(sv.ProductPath) {
		w.Header().Set("Content-Disposition", "attachment; filename="+sv.ProductName)
		w.Header().Set("Content-Type", "application/"+sv.ProductType)
		file, _ := os.Open(sv.ProductPath)
		defer file.Close()
		io.Copy(w, file)
		return
	}

	http.Redirect(w, r, "/404.html", http.StatusFound)
	return
}
