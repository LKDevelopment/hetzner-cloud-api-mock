package apiblueprint

import (
	"testing"
)

func TestTransformer_TransformPIBlueprintFile(t *testing.T) {
	t.Run("valid apib", func(t *testing.T) {
		transformer := Transformer{
			Path:                 "../../assets/",
			ApiBlueprintFileName: "test.apib",
			JSONFileName:         "drafterOutput.json",
		}
		err := transformer.TransformAPIBlueprintFile()
		if err != nil {
			t.Fatal(err)
		}
		if len(transformer.Resources) == 0 {
			t.Error("Transforming did not work")
		}
	})
	t.Run("invalid apib", func(t *testing.T) {
		transformer := Transformer{
			Path:                 "../../assets/",
			ApiBlueprintFileName: "test_wrong.apib",
			JSONFileName:         "drafterOutputWrong.json",
		}
		err := transformer.TransformAPIBlueprintFile()
		if err == nil || err.Error() != "exit status 4" {
			t.Fatal("Expected an error")
		}
	})
}
