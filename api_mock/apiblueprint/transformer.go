package apiblueprint

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/bukalapak/snowboard/api"
)

type Transformer struct {
	Resources            []*api.Resource
	Path                 string
	ApiBlueprintFileName string
	JSONFileName         string
}

func (t *Transformer) TransformAPIBlueprintFile() error {
	if !strings.HasSuffix(t.Path, "/") {
		t.Path = t.Path + "/"
	}

	json, err := t.throughDrafter()
	if err != nil {
		return err
	}
	elm, err := api.ParseJSON(strings.NewReader(json))
	if err != nil {
		return err
	}
	apiBlueprint, err := api.NewAPI(elm)
	t.Resources = apiBlueprint.ResourceGroups[1].Resources
	return nil
}
func (t *Transformer) throughDrafter() (string, error) {
	apibPath := t.Path + t.ApiBlueprintFileName
	jsonPath := t.Path + t.JSONFileName
	cmd := exec.Command("drafter", "-f", "json", apibPath)
	// open the out file for writing
	outfile, err := os.Create(jsonPath)
	if err != nil {
		return "", err
	}
	defer outfile.Close()
	cmd.Stdout = outfile

	err = cmd.Run()
	if err != nil {
		return "", err
	}

	json, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		return "", err
	}
	return string(json), nil
}
