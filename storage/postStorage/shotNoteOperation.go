package postStorage

import (
	"PillarsPhenomVFXWeb/mysqlUtility"
	"PillarsPhenomVFXWeb/utility"
)

func AddNote(n *utility.ShotNote) error {
	stmt, err := mysqlUtility.DBConn.Prepare("INSERT INTO shot_note (note_code, shot_code, project_code, picture, note_detail, note_type, note_verson, user_code, status, insert_datetime, update_datetime) VALUE(?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(n.NoteCode, n.ShotCode, n.ProjectCode, n.Picture, n.NoteDetail, n.NoteType, n.NoteVerson, n.UserCode, n.Status)
	if err != nil {
		return err
	}

	// 外包商保存note不含有ProjectCode,更新操作
	stmt, err = mysqlUtility.DBConn.Prepare("UPDATE shot_note SET project_code = (SELECT a.code FROM (SELECT MAX(project_code) code FROM shot_note WHERE status = 0 AND shot_code = ?) a) WHERE status = 0 AND project_code = 'aaa' AND shot_code = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(n.ShotCode, n.ShotCode)
	if err != nil {
		return err
	}

	return nil
}

func QueryNotes(shotCode *string) (*[]utility.ShotNote, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT note_code, picture, note_detail, note_type, note_verson, user_code FROM shot_note WHERE status = 0 AND shot_code = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Query(shotCode)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	var notes []utility.ShotNote
	for result.Next() {
		var note utility.ShotNote
		err = result.Scan(&note.NoteCode, &note.Picture, &note.NoteDetail, &note.NoteType, &note.NoteVerson, &note.UserCode)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}
	return &notes, nil
}
