package utility

// 前端请求返回自定义状态代码和信息
type FeedbackMessage struct {
	FeedbackCode int
	FeedbackText string
	Data         interface{}
}

type MaterialsOut struct {
	MaterialCode string
	MaterialName string
	MaterialType string
	MaterialPath string
	Length       string
	DpxPath      string
	JpgPath      string
	MovPath      string
}

type MaterialOut struct {
	Material
	Size   string
	Length string
}

type VendorAuth struct {
	OpenDetail   bool
	OpenDemo     bool
	DownMaterial bool
	UpDemo       bool
	UpProduct    bool
}
type ReferenceDetail struct {
	MatSource string
	MatDpx    string
	MatJPG    string
}
