package downloadAction

import (
	"PillarsPhenomVFXWeb/mysqlUtility"
	s "PillarsPhenomVFXWeb/session"
	es "PillarsPhenomVFXWeb/storage/editoralStorage"
	u "PillarsPhenomVFXWeb/utility"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type fileStruct struct {
	MaterialCode string
	SourceType   string
}

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	flag, _ := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		return
	}

	r.ParseForm()
	olen := len(r.Form["MaterialCode"]) + len(r.Form["SourceType"])
	if olen != 2 {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		//u.OutputJsonLog(w, 1, "Error parameter format", nil, "Error parameter format")
		return
	}
	var MaterialCode = r.Form["MaterialCode"][0]
	var SourceType = r.Form["SourceType"][0]
	if len(MaterialCode) == 0 {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		//u.OutputJsonLog(w, 12, "Error parameter MaterialCode", nil, "Error parameter MaterialCode")
		return
	}
	if len(SourceType) == 0 {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		//u.OutputJsonLog(w, 13, "Error parameter SourceType", nil, "Error parameter SourceType")
		return
	}

	var libraryCode string
	var materialName string
	var materialPath string
	var materialType string
	result := mysqlUtility.DBConn.QueryRow("SELECT library_code, material_name, material_type, material_path FROM material WHERE material_code= ? AND status = 0", MaterialCode)
	err := result.Scan(&(libraryCode), &(materialName), &(materialType), &(materialPath))
	if err != nil {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		//u.OutputJsonLog(w, 14, "File Not Found", nil, "File Not Found")
		return
	}
	library, err := es.QueryLibraryByLibraryCode(&libraryCode)
	if err != nil {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		//u.OutputJsonLog(w, 15, "Query Material failed!", nil, "editoralStorage.QueryLibraryByLibraryCode(&libraryCode) failed!")
		return
	}

	var filePath string
	var fileType string
	if SourceType == "Source" {
		if len(library.LibraryPath) == 0 {
			http.Redirect(w, r, "/404.html", http.StatusFound)
			//u.OutputJsonLog(w, 16, "LibraryPath failed!", nil, "SourceType failed!")
			return
		}
		filePath = library.LibraryPath + materialPath + materialType
		fileType = materialType
	} else if SourceType == "DPX" {
		if len(library.DpxPath) == 0 {
			http.Redirect(w, r, "/404.html", http.StatusFound)
			//u.OutputJsonLog(w, 16, "DpxPath failed!", nil, "SourceType failed!")
			return
		}
		filePath = library.DpxPath + materialPath + "dpx"
		fileType = "dpx"
	} else if SourceType == "JPG" {
		if len(library.JpgPath) == 0 {
			http.Redirect(w, r, "/404.html", http.StatusFound)
			//u.OutputJsonLog(w, 16, "JpgPath failed!", nil, "SourceType failed!")
			return
		}
		filePath = library.JpgPath + materialPath + "jpeg"
		fileType = "jpeg"
	} else if SourceType == "Mov" {
		if len(library.MovPath) == 0 {
			http.Redirect(w, r, "/404.html", http.StatusFound)
			//u.OutputJsonLog(w, 16, "MovPath failed!", nil, "SourceType failed!")
			return
		}
		filePath = library.MovPath + materialPath + "mp4"
		fileType = "mp4"
	} else {
		http.Redirect(w, r, "/404.html", http.StatusFound)
		//u.OutputJsonLog(w, 16, "SourceType failed!", nil, "SourceType failed!")
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+materialName+fileType)
	w.Header().Set("Content-Type", "application/"+fileType)
	file, _ := os.Open(filePath)
	defer file.Close()
	io.Copy(w, file)
}

func DownloadFileCheck(w http.ResponseWriter, r *http.Request) {
	flag, _ := s.GetAuthorityCode(w, r, "制片")
	if !flag {
		u.OutputJsonLog(w, 404, "Authority failed!", nil, "Authority failed!")
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "ioutil.ReadAll(r.Body) failed!")
		return
	}

	fs := fileStruct{}
	err = json.Unmarshal(data, &fs)
	if err != nil {
		u.OutputJsonLog(w, 12, "Pramaters failed!", nil, "json.Unmarshal(data, &fs) failed!")
		return
	}

	if len(fs.MaterialCode) == 0 || len(fs.SourceType) == 0 {
		u.OutputJsonLog(w, 13, "Pramaters failed!", nil, "Pramaters failed!")
		return
	}

	var libraryCode string
	var materialName string
	var materialPath string
	var materialType string
	result := mysqlUtility.DBConn.QueryRow("SELECT library_code, material_name, material_type, material_path FROM material WHERE material_code= ? AND status = 0", fs.MaterialCode)
	err = result.Scan(&(libraryCode), &(materialName), &(materialType), &(materialPath))
	if err != nil {
		u.OutputJsonLog(w, 14, "File Not Found", nil, "File Not Found")
		return
	}
	library, err := es.QueryLibraryByLibraryCode(&libraryCode)
	if err != nil {
		u.OutputJsonLog(w, 15, "Query Material failed!", nil, "editoralStorage.QueryLibraryByLibraryCode(&libraryCode) failed!")
		return
	}

	if fs.SourceType == "Source" {
		if len(library.LibraryPath) == 0 {
			u.OutputJsonLog(w, 16, "LibraryPath failed!", nil, "SourceType failed!")
			return
		}
	} else if fs.SourceType == "DPX" {
		if len(library.DpxPath) == 0 {
			u.OutputJsonLog(w, 16, "DpxPath failed!", nil, "SourceType failed!")
			return
		}
	} else if fs.SourceType == "JPG" {
		if len(library.JpgPath) == 0 {
			u.OutputJsonLog(w, 16, "JpgPath failed!", nil, "SourceType failed!")
			return
		}
	} else if fs.SourceType == "Mov" {
		if len(library.MovPath) == 0 {
			u.OutputJsonLog(w, 16, "MovPath failed!", nil, "SourceType failed!")
			return
		}
	} else {
		u.OutputJsonLog(w, 16, "SourceType failed!", nil, "SourceType failed!")
		return
	}

	u.OutputJsonLog(w, 0, "Download check succeed!", nil, "")
}
