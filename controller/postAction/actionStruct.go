package postAction

//body解析结构体
type interim struct {
	ProjectCode string
	ShotCode    string
}

// 给列表添加镜头的结构体
type addShots struct {
	ProjectCode string
	FolderCode  string
	VendorCode  string
	ShotCodes   []string
}
