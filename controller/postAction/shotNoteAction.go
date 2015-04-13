package postAction

import (
	s "PillarsPhenomVFXWeb/session"
	"PillarsPhenomVFXWeb/storage/postStorage"
	u "PillarsPhenomVFXWeb/utility"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func AddNote(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "") // 不需要权限
	if !flag {
		u.OutputJsonLog(w, 404, "session error!", nil, "")
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.AddNote: ioutil.ReadAll(r.Body) failed!")
		return
	}
	var note u.ShotNote
	err = json.Unmarshal(data, &note)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.AddNote: json.Unmarshal(data, &note) failed!")
		return
	}
	if len(note.ShotCode) == 0 || len(note.ProjectCode) == 0 || (len(note.NoteDetail) == 0 && len(note.Picture) == 0) {
		u.OutputJsonLog(w, 13, "Parameters Checked failed!", nil, "postAction.AddNote: Parameters Checked failed!")
		return
	}
	note.NoteCode = *u.GenerateCode(&userCode)
	note.UserCode = userCode

	err = postStorage.AddNote(&note)
	if err != nil {
		u.OutputJsonLog(w, 14, err.Error(), nil, "postAction.AddNote: postStorage.AddNote(&note) failed!")
		return
	}

	u.OutputJson(w, 0, "Add success.", nil)
}

func QueryNotes(w http.ResponseWriter, r *http.Request) {
	flag, _ := s.GetAuthorityCode(w, r, "") // 不需要权限
	if !flag {
		u.OutputJsonLog(w, 404, "session error!", nil, "")
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "postAction.QueryNotes: ioutil.ReadAll(r.Body) failed!")
		return
	}
	var i interim
	err = json.Unmarshal(data, &i)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "postAction.QueryNotes: json.Unmarshal(data, &ShotCode) failed!")
		return
	}
	if len(i.ShotCode) == 0 {
		u.OutputJsonLog(w, 13, "Parameters Checked failed!", nil, "postAction.AddNote: Parameters Checked failed!")
		return
	}

	result, err := postStorage.QueryNotes(&i.ShotCode)
	if result == nil || err != nil {
		u.OutputJsonLog(w, 14, "Query Notes failed!", nil, "postAction.QueryNotes: postStorage.QueryNotes(&ShotCode) failed!")
		return
	}

	u.OutputJson(w, 0, "Query success.", result)
}
