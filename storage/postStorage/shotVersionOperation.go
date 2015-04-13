package postStorage

import (
	"PillarsPhenomVFXWeb/mysqlUtility"
	"PillarsPhenomVFXWeb/utility"
)

func AddShotDemo(sv *utility.ShotVersion) error {
	stmt, err := mysqlUtility.DBConn.Prepare("INSERT INTO shot_version(version_code, shot_code, vendor_user, version_num, picture, demo_name, demo_type, demo_path, demo_detail, product_name, product_type, product_path, product_detail, status, insert_datetime, update_datetime) VALUE(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(sv.VersionCode, sv.ShotCode, sv.VendorUser, sv.VersionNum, sv.Picture, sv.DemoName, sv.DemoType, sv.DemoPath, sv.DemoDetail, sv.ProductName, sv.ProductType, sv.ProductPath, sv.ProductDetail, sv.Status)
	if err != nil {
		return err
	}
	return nil
}

func AddShotProduct(sv *utility.ShotVersion) error {
	stmt, err := mysqlUtility.DBConn.Prepare("UPDATE shot_version a JOIN (SELECT version_id, MAX(version_num) FROM shot_version WHERE status = 0 AND shot_code = ? AND vendor_user = ?) b ON a.version_id = b.version_id SET product_name = ?, product_type = ?, product_path = ?, product_detail = ?, update_datetime = NOW()")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(sv.ShotCode, sv.VendorUser, sv.ProductName, sv.ProductType, sv.ProductPath, sv.ProductDetail)
	if err != nil {
		return err
	}
	return nil
}

func GetShotVersionNum(sv *utility.ShotVersion) (*int, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT MAX(version_num) FROM shot_version WHERE status = 0 AND shot_code = ? AND vendor_user = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var n int
	result := stmt.QueryRow(sv.ShotCode, sv.VendorUser)
	result.Scan(&n)
	return &n, nil
}

func QueryShotVersion(shotCode *string) (*[]utility.ShotVersion, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT version_code, version_num, picture, demo_detail, IF(product_path LIKE '', 'N', 'Y') AS product_path FROM shot_version WHERE status = 0 AND shot_code = ? ORDER BY version_num DESC")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Query(shotCode)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	var versions []utility.ShotVersion
	for result.Next() {
		var version utility.ShotVersion
		err = result.Scan(&version.VersionCode, &version.VersionNum, &version.Picture, &version.DemoDetail, &version.ProductPath)
		if err != nil {
			return nil, err
		}
		versions = append(versions, version)
	}
	return &versions, nil
}

func QueryShotProduct(versionCode *string) (*utility.ShotVersion, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT product_name, product_type, product_path FROM shot_version WHERE status = 0 AND version_code = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var sv utility.ShotVersion
	result := stmt.QueryRow(versionCode)
	err = result.Scan(&sv.ProductName, &sv.ProductType, &sv.ProductPath)
	if err != nil {
		return nil, err
	}
	return &sv, nil
}
