package editoralStorage

import (
	"PillarsPhenomVFXWeb/mysqlUtility"
	"PillarsPhenomVFXWeb/pillarsLog"
	"PillarsPhenomVFXWeb/utility"
	"strconv"
)

func QueryFolders(projectCode *string, userCode *string) (*[]utility.MaterialFolder, error) {
	var num int
	count := mysqlUtility.DBConn.QueryRow("SELECT COUNT(1) FROM material_folder WHERE status = 0 AND project_code = ?", projectCode)
	count.Scan(&(num))
	if num == 0 {
		//没有数据,插入一条默认数据
		mf := utility.MaterialFolder{
			FolderName:   "自定义素材组",
			FatherCode:   "-1",
			LeafFlag:     "0",
			FolderDetail: "root",
			UserCode:     *userCode,
			ProjectCode:  *projectCode}
		_, err := InsertMaterialFolder(&mf)
		if err != nil {
			return nil, err
		}
	}

	stmt, err := mysqlUtility.DBConn.Prepare(`SELECT folder_id, folder_name, father_code, leaf_flag, folder_detail FROM material_folder WHERE status = 0 AND project_code = ?`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Query(projectCode)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer result.Close()
	var folders []utility.MaterialFolder
	for result.Next() {
		var mf utility.MaterialFolder
		err = result.Scan(&(mf.FolderCode), &(mf.FolderName), &(mf.FatherCode), &(mf.LeafFlag), &(mf.FolderDetail))
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
		folders = append(folders, mf)
	}
	return &folders, err
}

func InsertMaterialFolder(g *utility.MaterialFolder) (*utility.MaterialFolder, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`INSERT INTO material_folder (folder_name, father_code, leaf_flag, folder_Detail, user_code, project_code, status, insert_datetime, update_datetime) VALUES(?, ?, ?, ?, ?, ?, ?, NOW(), NOW())`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer stmt.Close()
	rs, err := stmt.Exec(g.FolderName, g.FatherCode, g.LeafFlag, g.FolderDetail, g.UserCode, g.ProjectCode, g.Status)
	if err != nil {
		return nil, err
	}
	id, err := rs.LastInsertId()
	g.FolderCode = strconv.FormatInt(id, 10)
	return g, err
}

func DeleteMaterialFolder(folderCode string, projectCode string) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`UPDATE material_folder SET status = 1 WHERE status = 0 AND (folder_id = ? OR father_code = ?)`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(folderCode, folderCode)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	// 删除素材组成功后，继续删除数据表的数据
	stmt, err = mysqlUtility.DBConn.Prepare(`UPDATE material_folder_data SET status = 1 WHERE status = 0 AND folder_id = ? AND project_code = ?`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(folderCode, projectCode)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}

	return true, err
}

func UpdateMaterialFolder(g *utility.MaterialFolder) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`UPDATE material_folder SET folder_name = ?, folder_detail = ?, user_code = ?, update_datetime = now() WHERE status = 0 AND folder_id = ?`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(g.FolderName, g.FolderDetail, g.UserCode, g.FolderCode)
	if err != nil {
		return false, err
	}

	return true, err
}

func QueryFolderById(folderId string) (*utility.MaterialFolder, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`SELECT folder_name, folder_Detail FROM material_folder WHERE status = 0 AND folder_id = ?`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Query(folderId)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer result.Close()

	var m utility.MaterialFolder
	if result.Next() {
		err = result.Scan(&(m.FolderName), &(m.FolderDetail))
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
	}
	return &m, err
}

func QueryFolderMaterialsCount(projectCode *string, id *string) (interface{}, error) {
	type result struct {
		IsHaveLeaf     bool
		IsHaveMaterial bool
		FatherCode     string
	}
	var rs result
	var num int
	// IsHaveLeaf
	count := mysqlUtility.DBConn.QueryRow("SELECT COUNT(1) FROM material_folder WHERE status = 0 AND project_code = ? AND father_code = ?", projectCode, id)
	count.Scan(&(num))
	if num == 0 {
		rs.IsHaveLeaf = false
	} else {
		rs.IsHaveLeaf = true
	}

	// IsHaveMaterial
	count = mysqlUtility.DBConn.QueryRow("SELECT COUNT(1) FROM material_folder_data WHERE status = 0 AND project_code = ? AND folder_id = ?", projectCode, id)
	err := count.Scan(&(num))
	if num == 0 {
		rs.IsHaveMaterial = false
	} else {
		rs.IsHaveMaterial = true
	}

	// FatherCode
	count = mysqlUtility.DBConn.QueryRow("SELECT father_code FROM material_folder WHERE status = 0 AND project_code = ? AND folder_id = ?", projectCode, id)
	err = count.Scan(&(rs.FatherCode))

	return rs, err
}
