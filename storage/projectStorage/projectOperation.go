package projectStorage

import (
	"PillarsPhenomVFXWeb/mysqlUtility"
	"PillarsPhenomVFXWeb/pillarsLog"
	"PillarsPhenomVFXWeb/utility"
	"html/template"
)

func InsertProject(p *utility.Project) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`INSERT INTO project (project_code,
		project_name, picture, project_leader, project_type, start_datetime,
		end_datetime, project_detail, status, insert_datetime, update_datetime)
		VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(p.ProjectCode, p.ProjectName, p.Picture, p.ProjectLeader,
		p.ProjectType, p.StartDatetime, p.EndDatetime, p.ProjectDetail, p.Status)
	if err != nil {
		return false, err
	}

	return true, err
}

func DeleteProjectByProjectCode(code *string) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`UPDATE project SET status = 1,
		update_datetime = now() WHERE project_code = ?`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(code)
	if err != nil {
		return false, err
	}

	return true, err
}

func UpdateProjectByProjectCode(p *utility.Project) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`UPDATE project SET project_name = ?, picture = ?, project_leader = ?, project_type = ?, start_datetime = ?, end_datetime = ?, project_detail = ?, update_datetime = now() WHERE project_code = ?`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(p.ProjectName, p.Picture, p.ProjectLeader, p.ProjectType, p.StartDatetime, p.EndDatetime, p.ProjectDetail, p.ProjectCode)
	if err != nil {
		return false, err
	}

	return true, err
}

func QueryProjectByProjectCode(projectCode *string) (*utility.Project, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`SELECT project_code, project_name, picture, project_leader, project_type, start_datetime, end_datetime, project_detail FROM project WHERE project_code = ? and status = 0`)
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
	var p utility.Project
	if result.Next() {
		err = result.Scan(&(p.ProjectCode), &(p.ProjectName), &(p.Picture), &(p.ProjectLeader), &(p.ProjectType), &(p.StartDatetime), &(p.EndDatetime), &(p.ProjectDetail))
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
	}
	return &p, err
}

func QueryProjectList(start int64, end int64) (*[]utility.Project, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`SELECT * FROM (SELECT project_code, project_name, picture, project_leader, project_type, start_datetime, end_datetime, project_detail FROM project WHERE status = 0 ORDER BY update_datetime DESC) T LIMIT ?, ?`)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Query(start, end)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer result.Close()
	var projectList []utility.Project
	for result.Next() {
		var p utility.Project
		err = result.Scan(&(p.ProjectCode), &(p.ProjectName), &(p.Picture), &(p.ProjectLeader), &(p.ProjectType), &(p.StartDatetime), &(p.EndDatetime), &(p.ProjectDetail))
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
		projectList = append(projectList, p)
	}
	return &projectList, err
}

func FindProjectList(args string) (*[]utility.Project, error) {
	args = template.HTMLEscapeString(args)
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT project_code, project_name, picture, project_leader, project_type, start_datetime, end_datetime, project_detail FROM project WHERE status = 0 AND (project_name LIKE '%" + args + "%' OR project_leader LIKE '%" + args + "%' OR project_type LIKE '%" + args + "%') ORDER BY update_datetime DESC")
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Query()
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		return nil, err
	}
	defer result.Close()
	var projectList []utility.Project
	for result.Next() {
		var p utility.Project
		err = result.Scan(&(p.ProjectCode), &(p.ProjectName), &(p.Picture), &(p.ProjectLeader), &(p.ProjectType), &(p.StartDatetime), &(p.EndDatetime), &(p.ProjectDetail))
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
		projectList = append(projectList, p)
	}
	return &projectList, err
}
