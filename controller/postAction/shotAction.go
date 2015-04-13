package postAction

import (
	s "PillarsPhenomVFXWeb/session"
	"PillarsPhenomVFXWeb/storage/postStorage"
	u "PillarsPhenomVFXWeb/utility"
	"encoding/json"
	"os"

	"io"
	"io/ioutil"
	"net/http"
)

func LoadEdlFile(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		u.OutputJson(w, 1, "parse upload error!", nil)
		return
	}
	formData := r.MultipartForm
	files := formData.File["files"]
	if len(files) > 0 {
		file, err := files[0].Open()
		defer file.Close()
		if err != nil {
			u.OutputJson(w, 12, "open edl file error!", nil)
			return
		}
		out, err := os.Create("./upload/" + files[0].Filename)
		defer out.Close()
		if err != nil {
			u.OutputJson(w, 13, "create edl file failed!", nil)
			return
		}
		_, err = io.Copy(out, file)
		if err != nil {
			u.OutputJson(w, 14, "io copy edl file failed!", nil)
		}
		// 解析edl文件得到镜头的信息
		edlShots, err := u.ReadEdl(out.Name())
		if err != nil && err.Error() != "EOF" {
			msg := "解析文件错误:" + err.Error()
			u.OutputJson(w, 15, msg, nil)
			return
		}
		if len(edlShots) == 0 {
			u.OutputJson(w, 16, "edl not find short!", nil)
			return
		}
		projectCode := formData.Value["ProjectCode"][0]
		// 查询镜头关联素材的详细信息
		shots, err := postStorage.EdlShotsToShots(files[0].Filename, projectCode, edlShots)
		if shots == nil || err != nil {
			u.OutputJson(w, 17, "edl not find material!", nil)
			return
		}
		// 保存镜头信息
		err = postStorage.InsertMultipleShot(userCode, projectCode, shots)
		if err != nil {
			u.OutputJson(w, 18, err.Error(), nil)
			return
		}

		u.OutputJson(w, 0, "upload success!", shots)
		return
	}

	//请求没有文件,返回错误信息
	u.OutputJson(w, 204, "not find upload file!", nil)
}

func QueryShots(w http.ResponseWriter, r *http.Request) {
	flag, _ := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		u.OutputJsonLog(w, 404, "Session failed!", nil, "")
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.QueryShots: ioutil.ReadAll(r.Body) failed!")
		return
	}
	var i interim
	err = json.Unmarshal(data, &i)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.QueryShots: json.Unmarshal(data, &ProjectCode) failed!")
		return
	}
	if len(i.ProjectCode) == 0 {
		u.OutputJsonLog(w, 13, "Parameter ProjectCode failed!", nil, "postAction.QueryShots: Parameter ProjectCode failed!")
		return
	}
	shots, err := postStorage.QueryShots(&i.ProjectCode)
	if shots == nil || err != nil {
		u.OutputJsonLog(w, 14, "Query shot list failed!", nil, "postAction.QueryShots: QueryShots(&i.ProjectCode) failed!")
		return
	}

	u.OutputJson(w, 0, "Query shot list success.", shots)
}

func QueryShotByShotCode(w http.ResponseWriter, r *http.Request) {
	flag, _ := s.GetAuthorityCode(w, r, "") // 不需要权限
	if !flag {
		u.OutputJsonLog(w, 404, "Session failed!", nil, "")
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.QueryShotByShotCode: ioutil.ReadAll(r.Body) failed!")
		return
	}
	var i interim
	err = json.Unmarshal(data, &i)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.QueryShotByShotCode: json.Unmarshal(data, &ShotCode) failed!")
		return
	}
	if len(i.ShotCode) == 0 {
		u.OutputJsonLog(w, 13, "Parameter ShotCode failed!", nil, "postAction.QueryShotByShotCode: Parameter ShotCode failed!")
		return
	}
	shot, err := postStorage.QueryShotByShotCode(&i.ShotCode)
	if shot == nil || err != nil {
		u.OutputJsonLog(w, 14, "Query shot failed!", nil, "postAction.QueryShotByShotCode: QueryShotByShotCode(&i.ShotCode) failed!")
		return
	}
	u.OutputJson(w, 0, "Query shot success.", shot)
}

func UpdateShot(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.UpdateShot: ioutil.ReadAll(r.Body) failed!")
		return
	}
	var shot u.Shot
	err = json.Unmarshal(data, &shot)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.UpdateShot: json.Unmarshal(data, &shot) failed!")
		return
	}
	// TODO 检查传入字段的有效性(需求未能确定非空字段)
	if len(shot.ShotCode) == 0 || len(shot.ShotName) == 0 {
		u.OutputJsonLog(w, 13, "Parameter Checked failed!", nil, "postAction.QueryShotByShotCode: Parameter Checked failed!")
		return
	}
	shot.UserCode = userCode

	err = postStorage.UpdateShot(&shot)
	if err != nil {
		u.OutputJsonLog(w, 14, err.Error(), nil, "postAction.UpdateShot: postStorage.UpdateShot(&shot) failed!")
		return
	}

	u.OutputJson(w, 0, "Update success.", shot)
}

func AddShot(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.AddShot: ioutil.ReadAll(r.Body) failed!")
		return
	}
	var shot u.Shot
	err = json.Unmarshal(data, &shot)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.AddShot: json.Unmarshal(data, &shot) failed!")
		return
	}
	if len(shot.ProjectCode) == 0 || len(shot.ShotName) == 0 || len(shot.ShotDetail) == 0 {
		u.OutputJsonLog(w, 13, "Parameter Checked failed!", nil, "postAction.AddShot: Parameter Checked failed!")
		return
	}
	shot.ShotCode = *u.GenerateCode(&userCode)
	shot.ShotFlag = "1" // 手动插入镜头的标识
	shot.UserCode = userCode

	err = postStorage.AddSingleShot(&shot)
	if err != nil {
		u.OutputJsonLog(w, 14, err.Error(), nil, "postAction.AddShot: postStorage.AddSingleShot(&shot) failed!")
		return
	}

	u.OutputJson(w, 0, "Add success.", shot)
}

func ModifyShotName(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.ModifyShotName: ioutil.ReadAll(r.Body) failed!")
		return
	}
	var shot u.Shot
	err = json.Unmarshal(data, &shot)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.ModifyShotName: json.Unmarshal(data, &shot) failed!")
		return
	}
	if len(shot.ShotCode) == 0 || len(shot.ShotName) == 0 {
		u.OutputJson(w, 13, "Parameter ShotName OR ShotCode failed!", nil)
		return
	}
	shot.UserCode = userCode

	err = postStorage.ModifyShotName(&shot)
	if err != nil {
		u.OutputJsonLog(w, 13, err.Error(), nil, "postAction.ModifyShotName: postStorage.ModifyShotName(&shot) failed!")
		return
	}

	u.OutputJson(w, 0, "ModifyShotName success.", shot)
}

func DeleteShot(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.DeleteShot: ioutil.ReadAll(r.Body) failed!")
		return
	}
	var shot u.Shot
	err = json.Unmarshal(data, &shot)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.DeleteShot: json.Unmarshal(data, &shot) failed!")
		return
	}
	if len(shot.ShotCode) == 0 {
		u.OutputJson(w, 13, "Parameter Checked failed!", nil)
		return
	}
	shot.UserCode = userCode

	err = postStorage.DeleteSingleShot(&shot)
	if err != nil {
		u.OutputJsonLog(w, 14, err.Error(), nil, "postAction.DeleteShot: postStorage.DeleteSingleShot(&shot) failed!")
		return
	}

	u.OutputJson(w, 0, "Delete success.", shot)
}
