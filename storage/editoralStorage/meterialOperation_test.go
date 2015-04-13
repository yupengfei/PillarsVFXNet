package editoralStorage

import (
	"PillarsPhenomVFXWeb/utility"
	"fmt"
	"testing"
)

func Test_InsertMaterial(t *testing.T) {
	temp := "testing"
	materialCode := "0bf4d1a1eec52d15ef55efcc4418d984"
	materialPath := utility.GenerateCode(&temp)
	materialName := utility.GenerateCode(&temp)
	materialType := "type"
	encodedPath := "/home/src/path"

	material := utility.Material{
		MaterialCode: materialCode,
		MaterialPath: *materialPath,
		MaterialName: *materialName,
		MaterialType: materialType,
		EncodedPath:  encodedPath,
	}

	result, err := InsertMaterial(&material)
	if result == false {
		fmt.Println(err.Error())
		t.Error("Insert into material failed")
	} else {
		fmt.Println(*utility.ObjectToJsonString(result))
	}
}

func Test_QueryMaterialByMaterialCode(t *testing.T) {
	materialCode := "0bf4d1a1eec52d15ef55efcc4418d984"
	_, err := QueryMaterialByMaterialCode(&materialCode)
	if err != nil {
		t.Error("query material by material_code failed")
	}
}

func Test_DeleteMaterialByMaterialCode(t *testing.T) {
	materialCode := "0bf4d1a1eec52d15ef55efcc4418d984"
	result, err := DeleteMaterialByMaterialCode(&materialCode)
	if err != nil {
		t.Error("delete material failed")
	} else {
		fmt.Println(*utility.ObjectToJsonString(result))
	}
}
