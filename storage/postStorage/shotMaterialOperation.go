package postStorage

import (
	"PillarsPhenomVFXWeb/mysqlUtility"
	"PillarsPhenomVFXWeb/utility"
)

func AddShotMaterial(sm *utility.ShotMaterial) error {
	stmt, err := mysqlUtility.DBConn.Prepare("INSERT INTO shot_material(material_code, shot_code, project_code, picture, material_name, material_type, material_detail, material_path, user_code, status, insert_datetime, update_datetime) VALUE(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(sm.MaterialCode, sm.ShotCode, sm.ProjectCode, sm.Picture, sm.MaterialName, sm.MaterialType, sm.MaterialDetail, sm.MaterialPath, sm.UserCode, sm.Status)
	if err != nil {
		return err
	}
	return nil
}

func DeleteShotMaterial(sm *utility.ShotMaterial) error {
	stmt, err := mysqlUtility.DBConn.Prepare("UPDATE shot_material SET status = 1, user_code = ?, update_datetime = NOW() WHERE material_code = ?")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(sm.UserCode, sm.MaterialCode)
	if err != nil {
		return err
	}
	return nil
}

func UpdateShotMaterial(sm *utility.ShotMaterial) error {
	stmt, err := mysqlUtility.DBConn.Prepare("UPDATE shot_material SET material_detail = ?, user_code = ?, update_datetime = NOW() WHERE material_code = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(sm.MaterialDetail, sm.UserCode, sm.MaterialCode)
	if err != nil {
		return err
	}
	return nil
}

func QueryShotMaterials(shotCode *string) (*[]utility.ShotMaterial, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT material_code, picture, material_name, material_type, material_detail FROM shot_material WHERE status = 0 AND shot_code = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Query(shotCode)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	var shotMaterials []utility.ShotMaterial
	for result.Next() {
		var sm utility.ShotMaterial
		err = result.Scan(&sm.MaterialCode, &sm.Picture, &sm.MaterialName, &sm.MaterialType, &sm.MaterialDetail)
		if err != nil {
			return nil, err
		}
		shotMaterials = append(shotMaterials, sm)
	}
	return &shotMaterials, nil
}

func QueryShotMaterial(materialCode *string) (*utility.ShotMaterial, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT picture, material_name, material_type, material_detail, material_path FROM shot_material WHERE status = 0 AND material_code = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var sm utility.ShotMaterial
	result := stmt.QueryRow(materialCode)
	err = result.Scan(&sm.Picture, &sm.MaterialName, &sm.MaterialType, &sm.MaterialDetail, &sm.MaterialPath)
	if err != nil {
		return nil, err
	}
	return &sm, nil
}
