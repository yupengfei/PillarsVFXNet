package userStorage

import (
	"PillarsPhenomVFXWeb/mysqlUtility"
	"PillarsPhenomVFXWeb/pillarsLog"
	"PillarsPhenomVFXWeb/utility"
)

func InsertIntoUser(user *utility.User) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`INSERT INTO user
		(user_code, password, display_name, picture, email, phone, user_authority, file_path, status)
		VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.UserCode, user.Password,
		user.DisplayName, user.Picture, user.Email,
		user.Phone, user.UserAuthority, user.FilePath, user.Status)
	if err != nil {
		return false, err
	} else {
		return true, err
	}
}

func DeleteUserByEmail(email *string) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`UPDATE user SET status = 1 WHERE email = ?`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(email)
	if err != nil {
		return false, err
	} else {
		return true, err
	}
}

func QueryUserList() (*[]utility.User, error) {
	result, err := mysqlUtility.DBConn.Query(`SELECT user_code, display_name,
		picture, email, phone, user_authority, file_path, status, insert_datetime, update_datetime
		FROM user WHERE status = 0`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer result.Close()
	var userList []utility.User
	var i int = 0
	for result.Next() {
		var user utility.User
		err = result.Scan(&(user.UserCode), &(user.DisplayName), &(user.Picture),
			&(user.Email), &(user.Phone), &(user.UserAuthority), &(user.FilePath), &(user.Status),
			&(user.InsertDatetime), &(user.UpdateDatetime))
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
		user.Status = i%2 + 1
		userList = append(userList, user)
		i++
	}
	return &userList, err
}

func CheckEmailAndPassword(email *string, password *string) (*utility.User, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`SELECT user_code, display_name,
		user_authority FROM user WHERE email = ? AND password = ? AND status = 0`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Query(*email, *password)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer result.Close()
	var user utility.User
	if result.Next() {
		err = result.Scan(&(user.UserCode), &(user.DisplayName), &(user.UserAuthority))
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
	}
	return &user, err
}

func GetUserAuthority(code *string) (*string, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`SELECT user_authority FROM user
		WHERE user_code = ? AND status = 0`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Query(code)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer result.Close()
	var authority string

	if result.Next() {
		err := result.Scan(&authority)
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
	}
	return &authority, err
}

func DeleteUserByUserCode(userCode *string) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`UPDATE user SET status = 1
		WHERE user_code = ?`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userCode)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	} else {
		return true, err
	}
}

func QueryUserByEmail(email *string) (*utility.User, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`SELECT user_code, display_name,
		picture, email, phone, user_authority, file_path, status, insert_datetime,
		update_datetime FROM user WHERE email = ?`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Query(email)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer result.Close()
	var user utility.User
	if result.Next() {
		err = result.Scan(&(user.UserCode), &(user.DisplayName), &(user.Picture),
			&(user.Email), &(user.Phone), &(user.UserAuthority), &(user.FilePath), &(user.Status),
			&(user.InsertDatetime), &(user.UpdateDatetime))
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
	}
	return &user, err

}

func UpdateUserByEmail(user *utility.User) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`UPDATE user SET display_name = ?,
	phone = ?, user_authority = ?, file_path = ?, update_datetime = now() WHERE email = ?`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.DisplayName, user.Phone, user.UserAuthority,
		user.FilePath, user.Email)
	if err != nil {
		return false, err
	} else {
		return true, err
	}
}

func QueryUserCode(email *string) (*string, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`SELECT user_code FROM user
		WHERE email = ?`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Query(email)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer result.Close()
	var user_code string

	if result.Next() {
		err := result.Scan(&user_code)
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
	}
	return &user_code, err
}

//查询外包商列表
type vendor struct {
	VendorUser string
	UserName   string
}

func QueryVendorUserList() (*[]vendor, error) {
	result, err := mysqlUtility.DBConn.Query(`SELECT user_code, display_name FROM user WHERE status = 0 AND user_authority = '分包商'`)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	var vendorList []vendor
	for result.Next() {
		var v vendor
		err = result.Scan(&(v.VendorUser), &(v.UserName))
		if err != nil {
			return nil, err
		}
		vendorList = append(vendorList, v)
	}

	return &vendorList, nil
}
