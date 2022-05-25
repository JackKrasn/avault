package action

import (
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Decrypt struct {
	cfg *Configuration

	Password string
}

func NewDecrypt(cfg *Configuration) *Decrypt {
	return &Decrypt{
		cfg: cfg,
	}
}

func (d *Decrypt) Run(encFileName string) error {
	fmt.Println("Starting decryption")
	path, err := os.Getwd()
	if err != nil {
		return err
	}

	filePath := filepath.Join(path, encFileName)

	d.cfg.Log("performing decryption for file %s", filePath)
	d.cfg.Log("password phrase %s", d.Password)
	yamlFile, err := ioutil.ReadFile(encFileName)
	if err != nil {
		return err
	}

	j2, err := yaml.YAMLToJSON(yamlFile)
	if err != nil {
		//fmt.Printf("err: %v\n", err)
		return err
	}
	d.cfg.Log(string(j2))

	return nil
}
