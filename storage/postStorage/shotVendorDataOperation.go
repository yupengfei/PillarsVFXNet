package postStorage

import (
	"PillarsPhenomVFXWeb/mysqlUtility"
	"PillarsPhenomVFXWeb/pillarsLog"
	"PillarsPhenomVFXWeb/utility"
)

func InsertShotVendorData(d *utility.ShotVendorData) error {
	stmt, err := mysqlUtility.DBConn.Prepare("INSERT INTO shot_vendor_data(data_code, vendor_code, vendor_user, shot_code, project_code, user_code, status, insert_datetime, update_datetime) SELECT ?, ?, ?, ?, ?, ?, ?, NOW(), NOW() FROM DUAL WHERE NOT EXISTS(SELECT shot_code FROM shot_vendor_data WHERE status = 0 AND project_code = ? AND vendor_code = ? AND shot_code = ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(d.DataCode, d.VendorCode, d.VendorUser, d.ShotCode, d.ProjectCode, d.UserCode, d.Status, d.ProjectCode, d.VendorCode, d.ShotCode)
	if err != nil {
		return err
	}

	return nil
}

func DeleteShotVendorData(userCode string, projectCode string, vendorCode string, shotCodes string) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("UPDATE shot_vendor_data SET user_code = ?, status = 1, update_datetime = NOW() WHERE status = 0 AND project_code = ? AND vendor_code = ? AND shot_code in (" + shotCodes + ")")
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userCode, projectCode, vendorCode)
	if err != nil {
		return false, err
	}

	return true, err
}

func QueryShotVendorShots(projectCode string, vendorCode string) (*[]shotOut, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT a.shot_code, a.shot_name, a.shot_status, a.picture, a.shot_flag, IF(b.library_path LIKE '', 'N', 'Y') AS source_path, IF(b.dpx_path LIKE '', 'N', 'Y') AS dpx_path, IF(b.jpg_path LIKE '', 'N', 'Y') AS jpg_path, IF(b.mov_path LIKE '', 'N', 'Y') AS mov_path FROM shot a LEFT JOIN library b ON a.library_code = b.library_code AND a.status = b.status WHERE a.status = 0 AND a.shot_code IN (SELECT shot_code from shot_vendor_data where status = 0 AND project_code = ? AND vendor_code = ?)")
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Query(projectCode, vendorCode)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer result.Close()
	var shots []shotOut
	for result.Next() {
		var so shotOut
		err = result.Scan(&(so.ShotCode), &so.ShotName, &so.ShotStatus, &so.Picture, &so.ShotFlag, &so.SourcePath, &so.DpxPath, &so.JpgPath, &so.MovPath)
		if err != nil {
			return nil, err
		}
		shots = append(shots, so)
	}
	return &shots, err
}

//外包商项目的镜头列表
func QueryVendorProjectShots(vendorCode string, userCode string) (*[]utility.Shot, error) {
	var code = vendorCode
	var and = "vendor_code"
	if vendorCode == "all" {
		code = userCode
		and = "vendor_user"
	}
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT project_code, shot_code, shot_name, picture, width, height, shot_fps FROM shot WHERE status = 0 AND shot_code IN (SELECT shot_code FROM shot_vendor_data WHERE status = 0 AND " + and + " = ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Query(code)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	var shots []utility.Shot
	for result.Next() {
		var s utility.Shot
		err = result.Scan(&(s.ProjectCode), &(s.ShotCode), &s.ShotName, &s.Picture, &s.Width, &s.Height, &s.ShotFps)
		if err != nil {
			return nil, err
		}
		shots = append(shots, s)
	}
	return &shots, err
}

// 指定外包商后的同步更新
func SpecifyShotVendorDataUser(sv *utility.ShotVendor) error {
	stmt, err := mysqlUtility.DBConn.Prepare("UPDATE shot_vendor_data SET vendor_user = ?, user_code = ?, update_datetime = NOW() WHERE status = 0 AND vendor_code = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(sv.VendorUser, sv.UserCode, sv.VendorCode)
	if err != nil {
		return err
	}

	return nil
}
