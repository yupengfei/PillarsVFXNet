package postAction

import (
	s "PillarsPhenomVFXWeb/session"
	ps "PillarsPhenomVFXWeb/storage/postStorage"
	u "PillarsPhenomVFXWeb/utility"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func AddShotVendorShots(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		u.OutputJson(w, 404, "session error!", nil)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.AddShotVendorshots: ioutil.ReadAll(r.Body) failed!")
		return
	}

	shots := addShots{}
	err = json.Unmarshal(data, &shots)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.AddShotVendorshots: json.Unmarshal(data, &shots) failed!")
		return
	}
	if len(shots.ProjectCode) == 0 || len(shots.VendorCode) == 0 || len(shots.ShotCodes) == 0 {
		u.OutputJsonLog(w, 13, "Parameters Checked failed!", nil, "postAction.AddShotVendorshots: Parameters Checked failed!")
		return
	}
	for _, value := range shots.ShotCodes {
		svd := u.ShotVendorData{
			DataCode:    *u.GenerateCode(&userCode),
			VendorCode:  shots.VendorCode,
			ShotCode:    value,
			UserCode:    userCode,
			ProjectCode: shots.ProjectCode,
			Status:      0,
		}
		err := ps.InsertShotVendorData(&svd)
		if err != nil {
			u.OutputJsonLog(w, 14, "Insert into shot_vendor_data failed!", nil, "postAction.AddShotVendorshots: ps.InsertShotVendorData(&svd) failed!")
			return
		}
	}

	u.OutputJsonLog(w, 0, "Add succeed.", nil, "")
}

func DeleteShotVendorShots(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		u.OutputJson(w, 404, "session error!", nil)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.DeleteShotVendorShots: ioutil.ReadAll(r.Body) failed!")
		return
	}

	shots := addShots{}
	err = json.Unmarshal(data, &shots)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.DeleteShotVendorShots: json.Unmarshal(data, &shots) failed!")
		return
	}

	if len(shots.ProjectCode) == 0 || len(shots.VendorCode) == 0 || len(shots.ShotCodes) == 0 {
		u.OutputJsonLog(w, 13, "Parameters Checked failed!", nil, "postAction.DeleteShotVendorShots: Parameters Checked failed!")
		return
	}
	var temp string
	for _, value := range shots.ShotCodes {
		temp += "'" + value + "', "
	}
	temp = temp[0 : len(temp)-2]
	result, _ := ps.DeleteShotVendorData(userCode, shots.ProjectCode, shots.VendorCode, temp)
	if result == false {
		u.OutputJsonLog(w, 14, "Delete data failed!", nil, "postAction.DeleteShotVendorShots: ps.DeleteShotVendorData(userCode, af.ProjectCode, af.VendorCode, temp) failed!")
		return
	}

	u.OutputJsonLog(w, 0, "Delete succeed!", nil, "")
}

func QueryShotVendorShots(w http.ResponseWriter, r *http.Request) {
	if !s.CheckAuthority(w, r, "制片") {
		u.OutputJsonLog(w, 404, "session error!", nil, "")
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.QueryShotVendorShots: ioutil.ReadAll(r.Body) failed!")
		return
	}

	shots := addShots{}
	err = json.Unmarshal(data, &shots)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.QueryShotVendorShots: json.Unmarshal(data, &shots) failed!")
		return
	}
	if len(shots.ProjectCode) == 0 || len(shots.VendorCode) == 0 {
		u.OutputJsonLog(w, 13, "Parameters Checked failed!", nil, "postAction.QueryShotVendorShots: Parameters Checked failed!")
		return
	}
	result, err := ps.QueryShotVendorShots(shots.ProjectCode, shots.VendorCode)
	if err != nil || result == nil {
		u.OutputJsonLog(w, 14, "Query failed!", nil, "postAction.QueryShotVendorShots: postStorage.QueryShotVendorShots(ProjectCode, VendorCode) failed!")
		return
	}

	u.OutputJsonLog(w, 0, "Query success.", result, "")
}
