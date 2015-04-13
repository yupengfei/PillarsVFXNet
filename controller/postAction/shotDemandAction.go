package postAction

import (
	s "PillarsPhenomVFXWeb/session"
	"PillarsPhenomVFXWeb/storage/postStorage"
	u "PillarsPhenomVFXWeb/utility"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func AddDemand(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.AddDemand: ioutil.ReadAll(r.Body) failed!")
		return
	}
	var demand u.ShotDemand
	err = json.Unmarshal(data, &demand)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.AddDemand: json.Unmarshal(data, &demand) failed!")
		return
	}
	if len(demand.ProjectCode) == 0 || len(demand.ShotCode) == 0 || (len(demand.DemandDetail) == 0 && len(demand.Picture) == 0) {
		u.OutputJsonLog(w, 13, "Parameters Checked failed!", nil, "postAction.AddDemand: Parameters Checked failed!")
		return
	}
	demand.DemandCode = *u.GenerateCode(&userCode)
	demand.UserCode = userCode

	err = postStorage.AddDemand(&demand)
	if err != nil {
		u.OutputJsonLog(w, 14, err.Error(), nil, "postAction.AddDemand: postStorage.AddDemand(&demand) failed!")
		return
	}
	demand.Picture = "" //图片不需要传回前台,减少返回的数据量
	u.OutputJson(w, 0, "Add success.", demand)
}

func DeleteDemand(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.DeleteDemand: ioutil.ReadAll(r.Body) failed!")
		return
	}
	var demand u.ShotDemand
	err = json.Unmarshal(data, &demand)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.DeleteDemand: json.Unmarshal(data, &demand) failed!")
		return
	}
	if len(demand.DemandCode) == 0 {
		u.OutputJsonLog(w, 13, "Parameters Checked failed!", nil, "postAction.DeleteDemand: Parameters Checked failed!")
		return
	}
	demand.UserCode = userCode

	err = postStorage.DeleteDemand(&demand)
	if err != nil {
		u.OutputJsonLog(w, 14, err.Error(), nil, "postAction.DeleteDemand: postStorage.DeleteDemand(&demand) failed!")
		return
	}

	u.OutputJson(w, 0, "Delete success.", nil)
}

func UpdateDemand(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.UpdateDemand: ioutil.ReadAll(r.Body) failed!")
		return
	}
	var demand u.ShotDemand
	err = json.Unmarshal(data, &demand)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.UpdateDemand: json.Unmarshal(data, &demand) failed!")
		return
	}
	if len(demand.DemandCode) == 0 || (len(demand.DemandDetail) == 0 && len(demand.Picture) == 0) {
		u.OutputJsonLog(w, 13, "Parameter Checked failed!", nil, "postAction.UpdateDemand: Parameter Checked failed!")
		return
	}
	demand.UserCode = userCode

	err = postStorage.UpdateDemand(&demand)
	if err != nil {
		u.OutputJsonLog(w, 14, err.Error(), nil, "postAction.UpdateDemand: postStorage.UpdateDemand(&demand) failed!")
		return
	}

	u.OutputJson(w, 0, "Update success.", nil)
}

func QueryDemands(w http.ResponseWriter, r *http.Request) {
	flag, _ := s.GetAuthorityCode(w, r, "") // 不需要权限
	if !flag {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.QueryDemands: ioutil.ReadAll(r.Body) failed!")
		return
	}
	var i interim
	err = json.Unmarshal(data, &i)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.QueryDemands: json.Unmarshal(data, &ShotCode) failed!")
		return
	}
	if len(i.ShotCode) == 0 {
		u.OutputJsonLog(w, 13, "Parameter ShotCode failed!", nil, "postAction.QueryDemands: Parameter ShotCode failed!")
		return
	}

	result, err := postStorage.QueryDemands(&i.ShotCode)
	if result == nil || err != nil {
		u.OutputJsonLog(w, 14, "Query ShotDemands failed!", nil, "postAction.QueryDemands: postStorage.QueryDemands(&ShotCode) failed!")
		return
	}

	u.OutputJson(w, 0, "Query success.", result)
}
