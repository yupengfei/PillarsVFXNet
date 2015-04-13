package postStorage

import (
	"PillarsPhenomVFXWeb/mysqlUtility"
	"PillarsPhenomVFXWeb/pillarsLog"
	"PillarsPhenomVFXWeb/utility"
)

func InsertShotFolderData(d *utility.ShotFolderData) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`INSERT INTO shot_folder_data(data_code, folder_id, shot_code, user_code, project_code, status, insert_datetime, update_datetime) SELECT ?, ?, ?, ?, ?, ?, NOW(), NOW() FROM DUAL WHERE NOT EXISTS(SELECT shot_code FROM shot_folder_data WHERE status = 0 AND project_code = ? AND folder_id = ? AND shot_code = ?)`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(d.DataCode, d.FolderCode, d.ShotCode, d.UserCode, d.ProjectCode, d.Status, d.ProjectCode, d.FolderCode, d.ShotCode)
	if err != nil {
		return false, err
	}

	return true, err
}

func DeleteShotFolderData(userCode string, projectCode string, id string, codes string) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("UPDATE shot_folder_data SET status = 1, user_code = ? WHERE status = 0 AND project_code = ? AND folder_id = ? AND shot_code in (" + codes + ")")
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userCode, projectCode, id)
	if err != nil {
		return false, err
	}

	return true, err
}
