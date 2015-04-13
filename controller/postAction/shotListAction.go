package postAction

import (
	s "PillarsPhenomVFXWeb/session"
	ps "PillarsPhenomVFXWeb/storage/postStorage"
	u "PillarsPhenomVFXWeb/utility"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetShotList(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		u.OutputJson(w, 404, "session error!", nil)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.GetShotList: ioutil.ReadAll(r.Body) failed!")
		return
	}

	var i interim
	err = json.Unmarshal(data, &i)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.GetShotList: json.Unmarshal(data, &interim) failed!")
		return
	}
	if len(i.ProjectCode) == 0 {
		u.OutputJson(w, 13, "Error parameter ShotCode", nil)
		return
	}

	shotList, err := ps.QueryFolders(&i.ProjectCode, &userCode)
	if err != nil {
		u.OutputJson(w, 14, "Query Filetypes failed!", nil)
		return
	}

	u.OutputJson(w, 0, "Query Filetypes succeed!", shotList)
}

func QueryFolderShots(w http.ResponseWriter, r *http.Request) {
	if !s.CheckAuthority(w, r, "制片") {
		u.OutputJson(w, 404, "session error!", nil)
		return
	}
	r.ParseForm()
	olen := len(r.Form["ProjectCode"]) + len(r.Form["FolderId"])
	if olen != 2 {
		u.OutputJson(w, 1, "Error parameter format", nil)
		return
	}
	if len(r.Form["ProjectCode"][0]) == 0 {
		u.OutputJson(w, 12, "Error parameter ProjectCode", nil)
		return
	}
	if len(r.Form["FolderId"][0]) == 0 {
		u.OutputJson(w, 13, "Error parameter FolderId", nil)
		return
	}

	materials, err := ps.FindFolderShots(r.Form["ProjectCode"][0], r.Form["FolderId"][0])
	if err != nil {
		u.OutputJson(w, 14, "Find Shots failed!", nil)
		return
	}

	u.OutputJson(w, 0, "Find Shots succeed!", materials)
}

func chectString(w http.ResponseWriter, r *http.Request, num int, args []string) bool {
	for _, str := range args {
		if len(r.Form[str][0]) == 0 {
			u.OutputJson(w, num, "Error parameter "+str, nil)
			return false
		}
		num += 1
	}
	return true
}

func AddFolder(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		u.OutputJson(w, 404, "session error!", nil)
		return
	}

	r.ParseForm()
	olen := len(r.Form["ProjectCode"]) + len(r.Form["FolderName"]) + len(r.Form["FatherCode"]) + len(r.Form["FolderDetail"])
	if olen != 4 {
		u.OutputJson(w, 1, "Error parameter format", nil)
		return
	}
	var args = []string{"ProjectCode", "FolderName", "FatherCode", "FolderDetail"}
	if !chectString(w, r, 12, args) {
		return
	}

	mf := u.ShotFolder{
		FolderName:   r.Form["FolderName"][0],
		FatherCode:   r.Form["FatherCode"][0],
		LeafFlag:     "0",
		FolderDetail: r.Form["FolderDetail"][0],
		UserCode:     userCode,
		ProjectCode:  r.Form["ProjectCode"][0],
		Status:       0,
	}
	result, err := ps.InsertShotFolder(&mf)
	if err != nil {
		u.OutputJson(w, 16, "Insert into material_folder failed!", nil)
		return
	}
	u.OutputJson(w, 0, "Add material_folder succeed!", result)
}

func DeleteFolder(w http.ResponseWriter, r *http.Request) {
	if !s.CheckAuthority(w, r, "制片") {
		u.OutputJson(w, 404, "session error!", nil)
		return
	}

	r.ParseForm()
	olen := len(r.Form["ProjectCode"]) + len(r.Form["FolderCode"])
	if olen != 2 {
		u.OutputJson(w, 1, "Error parameter format", nil)
		return
	}

	if len(r.Form["ProjectCode"][0]) == 0 {
		u.OutputJson(w, 12, "Error parameter ProjectCode", nil)
		return
	}

	if len(r.Form["FolderCode"][0]) == 0 {
		u.OutputJson(w, 13, "Error parameter FolderCode", nil)
		return
	}

	result, _ := ps.DeleteShotFolder(r.Form["FolderCode"][0], r.Form["ProjectCode"][0])
	if result == false {
		u.OutputJson(w, 14, "Delete material_folder failed!", nil)
		return
	}

	u.OutputJson(w, 0, "Delete material_folder succeed!", nil)
}

func QueryFolder(w http.ResponseWriter, r *http.Request) {
	if !s.CheckAuthority(w, r, "制片") {
		u.OutputJson(w, 404, "session error!", nil)
		return
	}

	r.ParseForm()
	olen := len(r.Form["FolderCode"])
	if olen != 1 {
		u.OutputJson(w, 1, "Error parameter format", nil)
		return
	}

	if len(r.Form["FolderCode"][0]) == 0 {
		u.OutputJson(w, 12, "Error parameter FolderCode", nil)
		return
	}

	result, err := ps.QueryFolderById(r.Form["FolderCode"][0])
	if err != nil {
		u.OutputJson(w, 13, "Query material_folder failed!", nil)
		return
	}

	u.OutputJson(w, 0, "Query material_folder succeed!", result)
}

func UpdateFolder(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		u.OutputJson(w, 404, "session error!", nil)
		return
	}

	r.ParseForm()
	olen := len(r.Form["FolderCode"]) + len(r.Form["FolderName"]) + len(r.Form["FolderDetail"])
	if olen != 3 {
		u.OutputJson(w, 1, "Error parameter format", nil)
		return
	}
	var args = []string{"FolderCode", "FolderName", "FolderDetail"}
	if !chectString(w, r, 12, args) {
		return
	}

	mf := u.ShotFolder{
		FolderCode:   r.Form["FolderCode"][0],
		FolderName:   r.Form["FolderName"][0],
		FolderDetail: r.Form["FolderDetail"][0],
		UserCode:     userCode,
	}
	result, _ := ps.UpdateShotFolder(&mf)
	if result == false {
		u.OutputJson(w, 15, "Update material_folder failed!", nil)
		return
	}

	u.OutputJson(w, 0, "Update material_folder succeed!", mf)
}

func AddFolderFiles(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		u.OutputJson(w, 404, "session error!", nil)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.AddFolderFiles: ioutil.ReadAll(r.Body) failed!")
		return
	}

	folderFiles := addShots{}
	err = json.Unmarshal(data, &folderFiles)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.AddFolderFiles: json.Unmarshal(data, &interim) failed!")
		return
	}
	if len(folderFiles.ProjectCode) == 0 || len(folderFiles.FolderCode) == 0 || len(folderFiles.ShotCodes) == 0 {
		u.OutputJsonLog(w, 13, "Parameters Checked failed!", nil, "postAction.AddNote: Parameters Checked failed!")
		return
	}
	temp := "insert"
	for _, value := range folderFiles.ShotCodes {
		mfd := u.ShotFolderData{
			DataCode:    *u.GenerateCode(&temp),
			FolderCode:  folderFiles.FolderCode,
			ShotCode:    value,
			UserCode:    userCode,
			ProjectCode: folderFiles.ProjectCode,
			Status:      0,
		}
		result, _ := ps.InsertShotFolderData(&mfd)
		if result == false {
			u.OutputJson(w, 14, "Insert into material_folder_data failed!", nil)
			return
		}
	}

	u.OutputJson(w, 0, "Insert into material_folder_data succeed!", nil)
}

func DeleteFolderFiles(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		u.OutputJson(w, 404, "session error!", nil)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.DeleteFolderFiles: ioutil.ReadAll(r.Body) failed!")
		return
	}

	af := addShots{}
	err = json.Unmarshal(data, &af)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.DeleteFolderFiles: json.Unmarshal(data, &interim) failed!")
		return
	}

	if len(af.ProjectCode) == 0 || len(af.FolderCode) == 0 || len(af.ShotCodes) == 0 {
		u.OutputJsonLog(w, 13, "Parameters Checked failed!", nil, "postAction.AddNote: Parameters Checked failed!")
		return
	}
	var temp string
	for _, value := range af.ShotCodes {
		temp += "'" + value + "', "
	}
	temp = temp[0 : len(temp)-2]
	result, _ := ps.DeleteShotFolderData(userCode, af.ProjectCode, af.FolderCode, temp)
	if result == false {
		u.OutputJson(w, 14, "Insert into material_folder_data failed!", nil)
		return
	}
	u.OutputJson(w, 0, "Insert into material_folder_data succeed!", nil)
}

func CountFolderFiles(w http.ResponseWriter, r *http.Request) {
	flag, _ := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		u.OutputJson(w, 404, "session error!", nil)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.CountFolderFiles: ioutil.ReadAll(r.Body) failed!")
		return
	}

	af := addShots{}
	err = json.Unmarshal(data, &af)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.CountFolderFiles: json.Unmarshal(data, &interim) failed!")
		return
	}

	if len(af.ProjectCode) == 0 || len(af.FolderCode) == 0 {
		u.OutputJsonLog(w, 13, "Parameters Checked failed!", nil, "postAction.CountFolderFiles: Parameters Checked failed!")
		return
	}

	result, err := ps.QueryFolderShotsCount(&af.ProjectCode, &af.FolderCode)
	if err != nil {
		u.OutputJson(w, 14, "Query material_folder_data count failed!", nil)
		return
	}
	u.OutputJson(w, 0, "Query material_folder_data count succeed!", result)
}
