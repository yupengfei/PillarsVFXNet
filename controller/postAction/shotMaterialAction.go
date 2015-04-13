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

// 判断文件是否存在: 存在返回true,不存在返回false
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
func AddShotMaterial(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		u.OutputJson(w, 404, "session error!", nil)
		return
	}

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		u.OutputJsonLog(w, 12, "parse upload error!", nil, "postAction.AddShotMaterial: r.ParseMultipartForm(32 << 20) failed!")
		return
	}
	formData := r.MultipartForm
	var sm u.ShotMaterial

	sm.ShotCode = formData.Value["ShotCode"][0]
	sm.ProjectCode = formData.Value["ProjectCode"][0]
	sm.MaterialType = formData.Value["MaterialType"][0]
	sm.MaterialDetail = formData.Value["MaterialDetail"][0]
	if len(sm.ShotCode) == 0 || len(sm.ProjectCode) == 0 || len(sm.MaterialType) == 0 || len(sm.MaterialDetail) == 0 {
		u.OutputJsonLog(w, 13, "Parameter Checked failed!", nil, "postAction.AddShotMaterial: Parameter Checked failed!")
		return
	}
	files := formData.File["files"]
	if len(files) > 0 {
		sm.MaterialName = files[0].Filename
		file, err := files[0].Open()
		defer file.Close()
		if err != nil {
			u.OutputJsonLog(w, 14, "Open upload file failed!", nil, "postAction.AddShotMaterial: Open upload file failed!")
			return
		}
		var path = "/home/pillars/Upload/material/" + sm.ProjectCode
		err = os.MkdirAll(path, 0777)
		if err != nil {
			u.OutputJsonLog(w, 15, "Create file path failed!", nil, "postAction.AddShotMaterial: Create file path failed!")
			return
		}
		createFile := path + "/" + sm.MaterialName
		if checkFileIsExist(createFile) { //如果文件存在
			u.OutputJsonLog(w, 202, "File Exist!", nil, "postAction.AddShotMaterial: File Exist!")
			return
		}
		out, err := os.OpenFile(createFile, os.O_CREATE|os.O_RDWR, 0777)
		if err != nil {
			u.OutputJsonLog(w, 16, "Create file failed!", nil, "postAction.AddShotMaterial: Create file failed!")
			return
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			u.OutputJsonLog(w, 17, "Copy file failed!", nil, "postAction.AddShotMaterial: Copy file failed!")
			return
		}

		// TODO 文件上传,保存成功,是否需要调用C++对素材抓图及其他信息
		sm.MaterialCode = *u.GenerateCode(&userCode)
		sm.MaterialPath = out.Name()
		sm.UserCode = userCode
		err = postStorage.AddShotMaterial(&sm)
		if err != nil {
			u.OutputJsonLog(w, 18, err.Error(), nil, "postAction.AddShotMaterial: postStorage.AddShotMaterial(&ShotMaterial) failed!")
			return
		}

		u.OutputJsonLog(w, 0, "Upload success.", nil, "")
		return
	}

	//请求没有文件,返回错误信息
	u.OutputJson(w, 204, "not find upload file!", nil)
}

func DeleteShotMaterial(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		u.OutputJson(w, 404, "session error!", nil)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.DeleteShotMaterial: ioutil.ReadAll(r.Body) failed!")
		return
	}
	var shotMaterial u.ShotMaterial
	err = json.Unmarshal(data, &shotMaterial)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.DeleteShotMaterial: json.Unmarshal(data, &shotMaterial) failed!")
		return
	}
	if len(shotMaterial.MaterialCode) == 0 {
		u.OutputJsonLog(w, 13, "Parameters Checked failed!", nil, "postAction.DeleteShotMaterial: Parameters Checked failed!")
		return
	}
	//查询文件的路径,删除服务器的文件
	sm, err := postStorage.QueryShotMaterial(&shotMaterial.MaterialCode)
	if err != nil || sm == nil {
		u.OutputJsonLog(w, 14, "File not found!", nil, "postAction.DeleteShotMaterial: postStorage.QueryShotMaterial(&MaterialCode) failed!")
		return
	}
	if checkFileIsExist(sm.MaterialPath) {
		err = os.Remove(sm.MaterialPath)
		if err != nil {
			u.OutputJsonLog(w, 15, "File delete failed!", nil, "postAction.DeleteShotMaterial: File delete failed!")
			return
		}
	}

	shotMaterial.UserCode = userCode
	err = postStorage.DeleteShotMaterial(&shotMaterial)
	if err != nil {
		u.OutputJsonLog(w, 16, err.Error(), nil, "postAction.DeleteShotMaterial: postStorage.DeleteShotMaterial(&shotMaterial) failed!")
		return
	}

	u.OutputJson(w, 0, "Delete success.", nil)
}

func UpdateShotMaterial(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.UpdateShotMaterial: ioutil.ReadAll(r.Body) failed!")
		return
	}
	var shotMaterial u.ShotMaterial
	err = json.Unmarshal(data, &shotMaterial)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.UpdateShotMaterial: json.Unmarshal(data, &shotMaterial) failed!")
		return
	}
	if len(shotMaterial.MaterialCode) == 0 || len(shotMaterial.MaterialDetail) == 0 {
		u.OutputJsonLog(w, 13, "Parameter Checked failed!", nil, "postAction.UpdateShotMaterial: Parameter Checked failed!")
		return
	}
	shotMaterial.UserCode = userCode

	err = postStorage.UpdateShotMaterial(&shotMaterial)
	if err != nil {
		u.OutputJsonLog(w, 14, err.Error(), nil, "postAction.UpdateShotMaterial: postStorage.UpdateShotMaterial(&shotMaterial) failed!")
		return
	}

	u.OutputJson(w, 0, "Update success.", nil)
}

func QueryShotMaterials(w http.ResponseWriter, r *http.Request) {
	flag, _ := s.GetAuthorityCode(w, r, "") // 不需要权限
	if !flag {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.QueryShotMaterials: ioutil.ReadAll(r.Body) failed!")
		return
	}
	var i interim
	err = json.Unmarshal(data, &i)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.QueryShotMaterials: json.Unmarshal(data, &ShotCode) failed!")
		return
	}
	if len(i.ShotCode) == 0 {
		u.OutputJsonLog(w, 13, "Parameter ShotCode failed!", nil, "postAction.QueryShotMaterials: Parameter ShotCode failed!")
		return
	}

	result, err := postStorage.QueryShotMaterials(&i.ShotCode)
	if result == nil || err != nil {
		u.OutputJsonLog(w, 14, "Query QueryShotMaterials failed!", nil, "postAction.QueryDemands: postStorage.QueryShotMaterials(&ShotCode) failed!")
		return
	}

	u.OutputJson(w, 0, "Query success.", result)
}

func DownloadShotMaterials(w http.ResponseWriter, r *http.Request) {
	flag, _ := s.GetAuthorityCode(w, r, "") // 不需要权限
	if !flag {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	r.ParseForm()
	olen := len(r.Form["MaterialCode"])
	if olen != 1 {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}
	var materialCode = r.Form["MaterialCode"][0]
	if len(materialCode) == 0 {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}
	//查询文件的路径
	sm, err := postStorage.QueryShotMaterial(&materialCode)
	if err != nil || sm == nil {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}
	if checkFileIsExist(sm.MaterialPath) {
		w.Header().Set("Content-Disposition", "attachment; filename="+sm.MaterialName)
		w.Header().Set("Content-Type", "application/"+sm.MaterialType)
		file, _ := os.Open(sm.MaterialPath)
		defer file.Close()
		io.Copy(w, file)
		return
	}

	http.Redirect(w, r, "/404.html", http.StatusFound)
	return
}
