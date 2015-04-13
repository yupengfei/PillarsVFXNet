package vendorAction

import (
	s "PillarsPhenomVFXWeb/session"
	ps "PillarsPhenomVFXWeb/storage/postStorage"
	u "PillarsPhenomVFXWeb/utility"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

//查询外包商列表
func GetVendorProjectList(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "分包商")
	if !flag {
		u.OutputJsonLog(w, 404, "session error!", nil, "")
		return
	}

	result, err := ps.GetVendorProject(&userCode)
	if err != nil || result == nil {
		u.OutputJsonLog(w, 1, "Query failed!", nil, "vendorAction.GetProjectList: postStorage.GetVendorProject(&userCode) failed!")
		return
	}

	u.OutputJsonLog(w, 0, "Query success.", result, "")
}

//分包商项目的镜头列表及需求列表信息
func QueryVendorProjectShots(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "分包商")
	if !flag {
		u.OutputJsonLog(w, 404, "session error!", nil, "")
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.OutputJsonLog(w, 1, "Read body failed!", nil, "vendorAction.QueryVendorProjectShots: ioutil.ReadAll(r.Body) failed!")
		return
	}

	type interim struct {
		VendorCode string
	}

	var i interim
	err = json.Unmarshal(data, &i)
	if err != nil {
		u.OutputJsonLog(w, 12, err.Error(), nil, "vendorAction.QueryShotVendorShots: json.Unmarshal(data, &shots) failed!")
		return
	}
	if len(i.VendorCode) == 0 {
		u.OutputJsonLog(w, 13, "Parameters Checked failed!", nil, "vendorAction.QueryShotVendorShots: Parameters Checked failed!")
		return
	}
	//镜头List
	result, err := ps.QueryVendorProjectShots(i.VendorCode, userCode)
	if err != nil || result == nil {
		u.OutputJsonLog(w, 14, "Query shots failed!", nil, "vendorAction.QueryVendorProjectShots: postStorage.QueryVendorProjectShots(VendorCode, userCode) failed!")
		return
	}

	type vendorShots struct {
		Shot    u.Shot
		Demands *[]u.ShotDemand
	}
	var vss []vendorShots
	//镜头的需求
	for _, s := range *result {
		var vs vendorShots
		vs.Shot = s
		ds, err := ps.QueryVendorDemands(&s.ShotCode)
		if err != nil || ds == nil {
			u.OutputJsonLog(w, 15, "Query shotDemands failed!", nil, "vendorAction.QueryDemands: postStorage.QueryVendorDemands(&ShotCode) failed!")
			return
		}
		vs.Demands = ds
		vss = append(vss, vs)
	}

	u.OutputJsonLog(w, 0, "Query success.", vss, "")
}

// 判断文件是否存在: 存在返回true,不存在返回false
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// 上传小样
func UploadDemo(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "分包商")
	if !flag {
		u.OutputJsonLog(w, 404, "session error!", nil, "")
		return
	}

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		u.OutputJsonLog(w, 1, "parse upload error!", nil, "vendorAction.UploadDemo: r.ParseMultipartForm(32 << 20) failed!")
		return
	}
	formData := r.MultipartForm
	if len(formData.Value["ShotCode"]) == 0 || len(formData.Value["ProjectCode"]) == 0 || len(formData.Value["DemoType"]) == 0 || len(formData.Value["DemoDetail"]) == 0 {
		u.OutputJsonLog(w, 12, "Parameter not find!", nil, "vendorAction.UploadDemo: Parameter not find!")
		return
	}
	var sv u.ShotVersion
	sv.VendorUser = userCode
	sv.ShotCode = formData.Value["ShotCode"][0]
	sv.ProjectCode = formData.Value["ProjectCode"][0]
	sv.DemoType = formData.Value["DemoType"][0]
	sv.DemoDetail = formData.Value["DemoDetail"][0]
	if len(sv.ShotCode) == 0 || len(sv.ProjectCode) == 0 || len(sv.DemoType) == 0 || len(sv.DemoDetail) == 0 {
		u.OutputJsonLog(w, 13, "Parameter Checked failed!", nil, "vendorAction.UploadDemo: Parameter Checked failed!")
		return
	}

	files := formData.File["files"]
	if len(files) > 0 {
		sv.DemoName = files[0].Filename
		file, err := files[0].Open()
		defer file.Close()
		if err != nil {
			u.OutputJsonLog(w, 14, "Open upload file failed!", nil, "vendorAction.UploadDemo: Open upload file failed!")
			return
		}
		var path = "/home/pillars/upload/demo/" + sv.ProjectCode
		err = os.MkdirAll(path, 0777)
		if err != nil {
			u.OutputJsonLog(w, 15, "Create file path failed!", nil, "vendorAction.UploadDemo: Create file path failed!")
			return
		}
		createFile := path + "/" + sv.DemoName
		if checkFileIsExist(createFile) { //如果文件存在
			u.OutputJsonLog(w, 202, "File Exist!", nil, "vendorAction.UploadDemo: File Exist!")
			return
		}
		out, err := os.OpenFile(createFile, os.O_CREATE|os.O_RDWR, 0777)
		if err != nil {
			u.OutputJsonLog(w, 16, "Create file failed!", nil, "vendorAction.UploadDemo: Create file failed!")
			return
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			u.OutputJsonLog(w, 17, "Copy file failed!", nil, "vendorAction.UploadDemo: Copy file failed!")
			return
		}

		// TODO 文件上传,保存成功,是否需要调用C++对素材抓图及其他信息
		//获取版本号
		num, err := ps.GetShotVersionNum(&sv)
		if num == nil || err != nil {
			u.OutputJsonLog(w, 18, "Get shot version number failed!", nil, "postAction.AddShotDemo: postStorage.GetShotVersionNum(&version) failed!")
			return
		}
		sv.VersionNum = *num + 1
		sv.VersionCode = *u.GenerateCode(&userCode)
		sv.DemoPath = out.Name()
		err = ps.AddShotDemo(&sv)
		if err != nil {
			u.OutputJsonLog(w, 19, err.Error(), nil, "vendorAction.UploadDemo: postStorage.AddShotDemo(&ShotVersion) failed!")
			return
		}

		u.OutputJsonLog(w, 0, "Upload success.", nil, "")
		return
	}

	//请求没有文件,返回错误信息
	u.OutputJson(w, 204, "Not find upload file!", nil)
}

// 上传成品,追加到最后一条小样的记录上
func UploadProduct(w http.ResponseWriter, r *http.Request) {
	flag, userCode := s.GetAuthorityCode(w, r, "分包商")
	if !flag {
		u.OutputJsonLog(w, 404, "session error!", nil, "")
		return
	}

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		u.OutputJsonLog(w, 1, "parse upload error!", nil, "vendorAction.UploadProduct: r.ParseMultipartForm(32 << 20) failed!")
		return
	}
	formData := r.MultipartForm
	if len(formData.Value["ShotCode"]) == 0 || len(formData.Value["ProjectCode"]) == 0 || len(formData.Value["ProductType"]) == 0 {
		u.OutputJsonLog(w, 12, "Parameter not find!", nil, "vendorAction.UploadProduct: Parameter not find!")
		return
	}
	var sv u.ShotVersion
	sv.ShotCode = formData.Value["ShotCode"][0]
	sv.ProjectCode = formData.Value["ProjectCode"][0]
	sv.ProductType = formData.Value["ProductType"][0]
	//sv.DemoDetail = formData.Value["ProductDetail"][0] //
	if len(sv.ShotCode) == 0 || len(sv.ProjectCode) == 0 || len(sv.ProductType) == 0 {
		u.OutputJsonLog(w, 13, "Parameter Checked failed!", nil, "vendorAction.UploadProduct: Parameter Checked failed!")
		return
	}

	files := formData.File["files"]
	if len(files) > 0 {
		sv.ProductName = files[0].Filename
		file, err := files[0].Open()
		defer file.Close()
		if err != nil {
			u.OutputJsonLog(w, 14, "Open upload file failed!", nil, "vendorAction.UploadProduct: Open upload file failed!")
			return
		}
		var path = "/home/pillars/upload/product/" + sv.ProjectCode
		err = os.MkdirAll(path, 0777)
		if err != nil {
			u.OutputJsonLog(w, 15, "Create file path failed!", nil, "vendorAction.UploadProduct: Create file path failed!")
			return
		}
		createFile := path + "/" + sv.ProductName
		if checkFileIsExist(createFile) { //如果文件存在
			u.OutputJsonLog(w, 202, "File Exist!", nil, "vendorAction.UploadProduct: File Exist!")
			return
		}
		out, err := os.OpenFile(createFile, os.O_CREATE|os.O_RDWR, 0777)
		if err != nil {
			u.OutputJsonLog(w, 16, "Create file failed!", nil, "vendorAction.UploadProduct: Create file failed!")
			return
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			u.OutputJsonLog(w, 17, "Copy file failed!", nil, "vendorAction.UploadProduct: Copy file failed!")
			return
		}

		// TODO 文件上传,保存成功,是否需要调用C++对素材抓图及其他信息
		sv.ProductPath = out.Name()
		sv.VendorUser = userCode
		err = ps.AddShotProduct(&sv)
		if err != nil {
			u.OutputJsonLog(w, 18, err.Error(), nil, "vendorAction.UploadProduct: postStorage.AddShotProduct(&ShotVersion) failed!")
			return
		}

		u.OutputJsonLog(w, 0, "Upload success.", nil, "")
		return
	}

	//请求没有文件,返回错误信息
	u.OutputJson(w, 204, "Not find upload file!", nil)
}
