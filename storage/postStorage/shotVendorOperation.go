package postStorage

import (
	"PillarsPhenomVFXWeb/mysqlUtility"
	"PillarsPhenomVFXWeb/utility"
)

func AddShotVendor(sv *utility.ShotVendor) error {
	stmt, err := mysqlUtility.DBConn.Prepare("INSERT INTO shot_vendor(vendor_code, project_code, vendor_user, vendor_name, vendor_detail, user_code, status, insert_datetime, update_datetime) VALUE(?, ?, ?, ?, ?, ?, ?, NOW(), NOW())")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(sv.VendorCode, sv.ProjectCode, sv.VendorUser, sv.VendorName, sv.VendorDetail, sv.UserCode, sv.Status)
	if err != nil {
		return err
	}
	return nil
}

func DeleteShotVendor(sv *utility.ShotVendor) error {
	stmt, err := mysqlUtility.DBConn.Prepare("UPDATE shot_vendor SET status = 1, user_code = ?, update_datetime = NOW() WHERE status = 0 AND vendor_code = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(sv.UserCode, sv.VendorCode)
	if err != nil {
		return err
	}

	// TODO 关联删除添加的镜头
	return nil
}

// 指定外包商(外包商的user_code)
func SpecifyShotVendorUser(sv *utility.ShotVendor) error {
	stmt, err := mysqlUtility.DBConn.Prepare("UPDATE shot_vendor SET vendor_user = ?, user_code = ?, update_datetime = NOW() WHERE status = 0 AND vendor_code = ?")
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

// 更新外包列表描述
func ModifyVendorDetail(sv *utility.ShotVendor) error {
	stmt, err := mysqlUtility.DBConn.Prepare("UPDATE shot_vendor SET vendor_detail = ?, user_code = ?, update_datetime = NOW() WHERE status = 0 AND vendor_code = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(sv.VendorDetail, sv.UserCode, sv.VendorCode)
	if err != nil {
		return err
	}
	return nil
}

// 单条外包列表的权限
func ModifyShotVendorAuth(sv *utility.ShotVendor) error {
	stmt, err := mysqlUtility.DBConn.Prepare("UPDATE shot_vendor SET open_detail = ?, open_demo = ?, down_material = ?, up_demo = ?, up_product = ?, user_code = ?, update_datetime = NOW() WHERE status = 0 AND vendor_code = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(sv.OpenDetail, sv.OpenDemo, sv.DownMaterial, sv.UpDemo, sv.UpProduct, sv.UserCode, sv.VendorCode)
	if err != nil {
		return err
	}
	return nil
}

// 查询外包商列表
func QueryShotVendorList(projectCode *string) (*[]utility.ShotVendor, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT vendor_code, vendor_name, vendor_user, vendor_detail FROM shot_vendor WHERE status = 0 AND project_code = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Query(projectCode)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	var vendors []utility.ShotVendor
	for result.Next() {
		var v utility.ShotVendor
		err = result.Scan(&v.VendorCode, &v.VendorName, &v.VendorUser, &v.VendorDetail)
		if err != nil {
			return nil, err
		}
		vendors = append(vendors, v)
	}
	return &vendors, nil
}

// 查询单条外包商列表信息
func QueryShotVendor(vendorCode *string) (*utility.ShotVendor, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT vendor_code, vendor_name, vendor_user, vendor_detail, open_detail, open_demo, down_material, up_demo, up_product FROM shot_vendor WHERE status = 0 AND vendor_code = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var v utility.ShotVendor
	result := stmt.QueryRow(vendorCode)
	err = result.Scan(&v.VendorCode, &v.VendorName, &v.VendorUser, &v.VendorDetail, &v.OpenDetail, &v.OpenDemo, &v.DownMaterial, &v.UpDemo, &v.UpProduct)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

//外包商项目列表查询
func GetVendorProject(vendorUser *string) (*[]utility.ShotVendor, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT vendor_code, vendor_name, vendor_detail FROM shot_vendor WHERE status = 0 AND vendor_user = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Query(vendorUser)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	var vendors []utility.ShotVendor
	for result.Next() {
		var v utility.ShotVendor
		err = result.Scan(&v.VendorCode, &v.VendorName, &v.VendorDetail)
		if err != nil {
			return nil, err
		}
		vendors = append(vendors, v)
	}
	return &vendors, nil
}
