package postAction

import (
	s "PillarsPhenomVFXWeb/session"
	"PillarsPhenomVFXWeb/storage/postStorage"
	u "PillarsPhenomVFXWeb/utility"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

func AddShotVendor(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		u.OutputJsonLog(w, 404, "session error!", nil, "")
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.AddShotVendor: ioutil.ReadAll(r.Body) failed!")
		return
	}
	var vendor u.ShotVendor
	json.Unmarshal(data, &vendor)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.AddShotVendor: json.Unmarshal(data, &vendor) failed!")
		return
	}
	if len(vendor.ProjectCode) == 0 || len(vendor.VendorName) == 0 || len(vendor.VendorDetail) == 0 {
		u.OutputJsonLog(w, 13, "Parameters Checked failed!", nil, "postAction.AddShotVendor: Parameters Checked failed!")
		return
	}
	vendor.VendorCode = *u.GenerateCode(&userCode)
	vendor.UserCode = userCode
	err = postStorage.AddShotVendor(&vendor)
	if err != nil {
		u.OutputJsonLog(w, 14, err.Error(), nil, "postAction.AddShotVendor: postStorage.AddShotVendor(&vendor) failed!")
		return
	}

	u.OutputJsonLog(w, 0, "Modify success.", vendor, "")
}

// TODO 确定级联删除内容
func DeleteShotVendor(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		u.OutputJsonLog(w, 404, "session error!", nil, "")
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.DeleteShotVendor: ioutil.ReadAll(r.Body) failed!")
		return
	}
	var vendor u.ShotVendor
	json.Unmarshal(data, &vendor)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.DeleteShotVendor: json.Unmarshal(data, &vendor) failed!")
		return
	}
	if len(vendor.VendorCode) == 0 {
		u.OutputJsonLog(w, 13, "Parameters Checked failed!", nil, "postAction.DeleteShotVendor: Parameters Checked failed!")
		return
	}
	vendor.UserCode = userCode
	err = postStorage.DeleteShotVendor(&vendor)
	if err != nil {
		u.OutputJsonLog(w, 14, err.Error(), nil, "postAction.DeleteShotVendor: postStorage.DeleteShotVendor(&vendor) failed!")
		return
	}

	u.OutputJsonLog(w, 0, "Modify success.", nil, "")
}

func SpecifyShotVendorUser(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		u.OutputJsonLog(w, 404, "session error!", nil, "")
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.SpecifyShotVendorUser: ioutil.ReadAll(r.Body) failed!")
		return
	}
	var vendor u.ShotVendor
	json.Unmarshal(data, &vendor)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.SpecifyShotVendorUser: json.Unmarshal(data, &vendor) failed!")
		return
	}
	if len(vendor.VendorCode) == 0 || len(vendor.VendorUser) == 0 {
		u.OutputJsonLog(w, 13, "Parameters Checked failed!", nil, "postAction.SpecifyShotVendorUser: Parameters Checked failed!")
		return
	}
	vendor.UserCode = userCode
	err = postStorage.SpecifyShotVendorUser(&vendor)
	if err != nil {
		u.OutputJsonLog(w, 14, err.Error(), nil, "postAction.SpecifyShotVendorUser: postStorage.SpecifyShotVendorUser(&vendor) failed!")
		return
	}
	// 指定外包商后更新data表
	err = postStorage.SpecifyShotVendorDataUser(&vendor)
	if err != nil {
		u.OutputJsonLog(w, 15, err.Error(), nil, "postAction.SpecifyShotVendorUser: postStorage.SpecifyShotVendorUser(&vendor) failed!")
		return
	}

	u.OutputJsonLog(w, 0, "Modify success.", nil, "")
}

func ModifyVendorDetail(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		u.OutputJsonLog(w, 404, "session error!", nil, "")
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.ModifyVendorDetail: ioutil.ReadAll(r.Body) failed!")
		return
	}
	var vendor u.ShotVendor
	json.Unmarshal(data, &vendor)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.ModifyVendorDetail: json.Unmarshal(data, &vendor) failed!")
		return
	}
	if len(vendor.VendorCode) == 0 || len(vendor.VendorDetail) == 0 {
		u.OutputJsonLog(w, 13, "Parameters Checked failed!", nil, "postAction.ModifyVendorDetail: Parameters Checked failed!")
		return
	}
	vendor.UserCode = userCode
	err = postStorage.ModifyVendorDetail(&vendor)
	if err != nil {
		u.OutputJsonLog(w, 14, err.Error(), nil, "postAction.ModifyVendorDetail: postStorage.ModifyVendorDetail(&vendor) failed!")
		return
	}

	u.OutputJsonLog(w, 0, "Modify success.", nil, "")
}

func ModifyShotVendorAuth(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		u.OutputJsonLog(w, 404, "session error!", nil, "")
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.ModifyShotVendorAuth: ioutil.ReadAll(r.Body) failed!")
		return
	}
	var vendor u.ShotVendor
	json.Unmarshal(data, &vendor)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.ModifyShotVendorAuth: json.Unmarshal(data, &vendor) failed!")
		return
	}
	if len(vendor.VendorCode) == 0 {
		u.OutputJsonLog(w, 13, "Parameters Checked failed!", nil, "postAction.ModifyShotVendorAuth: Parameters Checked failed!")
		return
	}
	vendor.UserCode = userCode
	err = postStorage.ModifyShotVendorAuth(&vendor)
	if err != nil {
		u.OutputJsonLog(w, 14, err.Error(), nil, "postAction.ModifyShotVendorAuth: postStorage.ModifyShotVendorAuth(&vendor) failed!")
		return
	}

	u.OutputJsonLog(w, 0, "Modify success.", nil, "")
}

func QueryShotVendor(w http.ResponseWriter, r *http.Request) {
	flag, _ := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		u.OutputJsonLog(w, 404, "session error!", nil, "")
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.QueryShotVendor: ioutil.ReadAll(r.Body) failed!")
		return
	}
	var vendor u.ShotVendor
	json.Unmarshal(data, &vendor)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.QueryShotVendor: json.Unmarshal(data, &vendor) failed!")
		return
	}
	if len(vendor.VendorCode) == 0 {
		u.OutputJsonLog(w, 13, "Parameters Checked failed!", nil, "postAction.QueryShotVendor: Parameters Checked failed!")
		return
	}
	result, err := postStorage.QueryShotVendor(&vendor.VendorCode)
	if err != nil || result == nil {
		u.OutputJsonLog(w, 14, "Query failed!", nil, "postAction.QueryShotVendor: postStorage.QueryShotVendor(&vendorCode) failed!")
		return
	}
	//用户权限前端需要拼接成字符串
	result.UserCode = strconv.Itoa(result.OpenDetail) + "," + strconv.Itoa(result.OpenDemo) + "," + strconv.Itoa(result.DownMaterial) + "," + strconv.Itoa(result.UpDemo) + "," + strconv.Itoa(result.UpProduct)

	u.OutputJsonLog(w, 0, "Query success.", result, "")
}

func QueryShotVendorList(w http.ResponseWriter, r *http.Request) {
	flag, _ := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		u.OutputJsonLog(w, 404, "session error!", nil, "")
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.QueryShotVendorList: ioutil.ReadAll(r.Body) failed!")
		return
	}
	var vendor u.ShotVendor
	json.Unmarshal(data, &vendor)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.QueryShotVendorList: json.Unmarshal(data, &vendor) failed!")
		return
	}
	if len(vendor.ProjectCode) == 0 {
		u.OutputJsonLog(w, 13, "Parameters Checked failed!", nil, "postAction.QueryShotVendorList: Parameters Checked failed!")
		return
	}
	result, err := postStorage.QueryShotVendorList(&vendor.ProjectCode)
	if err != nil || result == nil {
		u.OutputJsonLog(w, 14, "Query failed!", nil, "postAction.QueryShotVendorList: postStorage.QueryShotVendorList(&vendor) failed!")
		return
	}

	u.OutputJsonLog(w, 0, "Query success.", result, "")
}
