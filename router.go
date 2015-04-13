package main

import (
	"PillarsPhenomVFXWeb/controller/downloadAction"
	"PillarsPhenomVFXWeb/controller/editoralAction"
	"PillarsPhenomVFXWeb/controller/loginAction"
	"PillarsPhenomVFXWeb/controller/postAction"
	"PillarsPhenomVFXWeb/controller/projectAction"
	"PillarsPhenomVFXWeb/controller/userAction"
	"PillarsPhenomVFXWeb/controller/vendorAction"
	"net/http"
)

func RouterBinding() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./pages"))))

	http.HandleFunc("/login_action", loginAction.Login)

	http.HandleFunc("/user_add", userAction.AddUser)
	http.HandleFunc("/user_del", userAction.DeleteUser)
	http.HandleFunc("/user_upd", userAction.UpdateUser)
	http.HandleFunc("/user_sel", userAction.QueryUser)
	http.HandleFunc("/user_list", userAction.UserList)

	http.HandleFunc("/project_add", projectAction.AddProject)
	http.HandleFunc("/project_del", projectAction.DeleteProject)
	http.HandleFunc("/project_upd", projectAction.UpdateProject)
	http.HandleFunc("/project_sel", projectAction.QueryProject)
	http.HandleFunc("/project_list", projectAction.ProjectList)
	http.HandleFunc("/project_load", projectAction.LoadProject)
	http.HandleFunc("/project_find", projectAction.FindProjects)

	http.HandleFunc("/editoral_library", editoralAction.GetLibrarys)
	http.HandleFunc("/editoral_library_add", editoralAction.AddLibrary)
	http.HandleFunc("/editoral_library_materials", editoralAction.GetLibraryFileList)
	http.HandleFunc("/editoral_find_materials", editoralAction.FindMaterials)
	http.HandleFunc("/editoral_filetype", editoralAction.GetFiletypes)
	http.HandleFunc("/editoral_material", editoralAction.GetMaterialInfo)
	http.HandleFunc("/editoral_folder", editoralAction.GetFolders)
	http.HandleFunc("/editoral_folder_materials", editoralAction.QueryFolderMaterials)
	http.HandleFunc("/editoral_folder_add", editoralAction.AddFolder)
	http.HandleFunc("/editoral_folder_del", editoralAction.DeleteFolder)
	http.HandleFunc("/editoral_folder_que", editoralAction.QueryFolder)
	http.HandleFunc("/editoral_folder_upd", editoralAction.UpdateFolder)
	http.HandleFunc("/editoral_folder_addfiles", editoralAction.AddFolderFiles)
	http.HandleFunc("/editoral_folder_delfiles", editoralAction.DeleteFolderFiles)
	http.HandleFunc("/editoral_folder_countfiles", editoralAction.CountFolderFiles)
	http.HandleFunc("/editoral_download_file", downloadAction.DownloadFile)

	// 上传edl文件
	http.HandleFunc("/post_upload_edl", postAction.LoadEdlFile)
	// 镜头信息的增删改查
	http.HandleFunc("/post_shot_add", postAction.AddShot)
	http.HandleFunc("/post_shot_del", postAction.DeleteShot)
	http.HandleFunc("/post_shot_upd", postAction.UpdateShot)
	http.HandleFunc("/post_shot_updshotname", postAction.ModifyShotName)
	http.HandleFunc("/post_shot_list", postAction.QueryShots)
	http.HandleFunc("/post_shot_que", postAction.QueryShotByShotCode)
	// 镜头制作需求增删改查
	http.HandleFunc("/post_shot_demand_add", postAction.AddDemand)
	http.HandleFunc("/post_shot_demand_del", postAction.DeleteDemand)
	http.HandleFunc("/post_shot_demand_upd", postAction.UpdateDemand)
	http.HandleFunc("/post_shot_demand_que", postAction.QueryDemands)
	// 镜头参考素材增删改查
	http.HandleFunc("/post_shot_material_add", postAction.AddShotMaterial)
	http.HandleFunc("/post_shot_material_del", postAction.DeleteShotMaterial)
	//http.HandleFunc("/post_shot_material_upd", postAction.UpdateShotMaterial) --No USE
	http.HandleFunc("/post_shot_material_que", postAction.QueryShotMaterials)
	http.HandleFunc("/post_shot_material_dow", postAction.DownloadShotMaterials)
	// 镜头NOTE增查
	http.HandleFunc("/post_shot_note_add", postAction.AddNote)
	http.HandleFunc("/post_shot_note_que", postAction.QueryNotes)
	// 镜头版本查,下载
	http.HandleFunc("/post_shot_demo_version", postAction.QueryShotVersion)
	http.HandleFunc("/post_shot_product_dow", postAction.DownloadShotProduct)
	// 镜头自定义分组增删改查,添加删除镜头
	http.HandleFunc("/post_shot_folder", postAction.GetShotList)
	http.HandleFunc("/post_shot_folder_shots", postAction.QueryFolderShots)
	http.HandleFunc("/post_shot_folder_add", postAction.AddFolder)
	http.HandleFunc("/post_shot_folder_del", postAction.DeleteFolder)
	http.HandleFunc("/post_shot_folder_que", postAction.QueryFolder)
	http.HandleFunc("/post_shot_folder_upd", postAction.UpdateFolder)
	http.HandleFunc("/post_shot_folder_addfiles", postAction.AddFolderFiles)
	http.HandleFunc("/post_shot_folder_delfiles", postAction.DeleteFolderFiles)
	http.HandleFunc("/post_shot_folder_countfiles", postAction.CountFolderFiles)
	// 镜头外包商列表增删改查
	http.HandleFunc("/post_shot_vendor_add", postAction.AddShotVendor)
	http.HandleFunc("/post_shot_vendor_del", postAction.DeleteShotVendor)
	http.HandleFunc("/post_shot_vendor_specify", postAction.SpecifyShotVendorUser)
	http.HandleFunc("/post_shot_vendor_detail", postAction.ModifyVendorDetail)
	http.HandleFunc("/post_shot_vendor_auth", postAction.ModifyShotVendorAuth)
	http.HandleFunc("/post_shot_vendor_que", postAction.QueryShotVendor)
	http.HandleFunc("/post_shot_vendor_list", postAction.QueryShotVendorList)
	http.HandleFunc("/user_vendor_list", userAction.GetVendorList)
	// 镜头外包商列表镜头增删查
	http.HandleFunc("/post_shot_vendor_addShots", postAction.AddShotVendorShots)
	http.HandleFunc("/post_shot_vendor_delShots", postAction.DeleteShotVendorShots)
	http.HandleFunc("/post_shot_vendor_queShots", postAction.QueryShotVendorShots)
	// 外包商项目列表
	http.HandleFunc("/vendor_project_list", vendorAction.GetVendorProjectList)
	// 外包商项目镜头及需求列表
	http.HandleFunc("/vendor_project_shots", vendorAction.QueryVendorProjectShots)
	// 外包商小样,成品上传
	http.HandleFunc("/vendor_demo_upload", vendorAction.UploadDemo)
	http.HandleFunc("/vendor_product_upload", vendorAction.UploadProduct)

	// ---------------------------- 尚未测试 ------------------------------

	// ---------------------------- 待实现 ------------------------------

	// 下载验证,暂未使用
	http.HandleFunc("/editoral_download_file_check", downloadAction.DownloadFileCheck)
	//http.HandleFunc("/.*", NotFound) // TODO 想实现未知路由地址访问的404页面跳转
}
