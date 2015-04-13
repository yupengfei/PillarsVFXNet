package editoralAction

import (
	u "PillarsPhenomVFXWeb/utility"
	"os"
	"path/filepath"
	"strings"
)

/**
TODO 素材类型后缀格式需要添加在行25
*/
func LoadMaterials(fullPath string) (error, []*u.Material) {
	var materials []*u.Material
	var m *u.Material
	err := filepath.Walk(fullPath, func(path string, fi os.FileInfo, err error) error {
		if nil == fi {
			return err
		}
		if fi.IsDir() {
			return nil
		}
		name := fi.Name()
		if strings.Contains(name, ".R3D") || strings.Contains(name, "ARRIRAW") {
			sname := strings.Split(name, ".")
			if len(sname) != 2 {
				return nil
			}
			basePath := strings.Replace(strings.Replace(path, fullPath, "", 1), sname[1], "", -1)
			m = &u.Material{}
			m.MaterialName = sname[0]
			m.MaterialPath = basePath
			m.MaterialType = sname[1]
			materials = append(materials, m)
		}
		return nil
	})
	if err != nil {
		return err, nil
	}

	return err, materials
}
