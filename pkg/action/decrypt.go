package action

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/JackKrasn/avault/pkg/cli"
	"github.com/ghodss/yaml"
	vault "github.com/sosedoff/ansible-vault-go"
)

type Decrypt struct {
	cfg       *Configuration
	Settings  *cli.EnvSettings
	OutputDir bool
}

func NewDecrypt(cfg *Configuration) *Decrypt {
	return &Decrypt{
		cfg: cfg,
	}
}

func (d *Decrypt) Run(encFileName string) (string, error) {
	decryptedFileName := encFileName + ".dec"
	yamlFile, err := os.ReadFile(encFileName)
	if err != nil {
		return "", err
	}
	j2, err := yaml.YAMLToJSON(yamlFile)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return "", err
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(j2, &m)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return "", err
	}
	walk(m, d.Settings.Password, d.Settings.Dry)
	// write encypted data to the yaml file
	data, err := yaml.Marshal(&m)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return "", err
	}

	if !d.OutputDir {
		fmt.Println(string(data))
		return decryptedFileName, nil
	}

	err2 := os.WriteFile(decryptedFileName, data, 0644)
	if err2 != nil {
		fmt.Printf("err: %v\n", err)
		return "", err2
	}
	return decryptedFileName, nil
}

func walk(data map[string]interface{}, passwordPhrase string, dryRun bool) {
	for key, el := range data {
		if el == nil {
			continue
		}
		if reflect.TypeOf(el).Kind() == reflect.Map {
			walk(el.(map[string]interface{}), passwordPhrase, dryRun)
		}
		if reflect.TypeOf(el).Kind() == reflect.String {
			if isEncrypted(el.(string)) {
				if dryRun {
					data[key] = "\"*****\""
				} else {
					// Decrypt secret data
					decryptedStr, err := vault.Decrypt(el.(string), passwordPhrase)
					if err != nil {
						log.Fatalf("Can'not decrypt key: %v\n", key)
					}
					data[key] = decryptedStr
				}
			}
		}
		if reflect.TypeOf(el).Kind() == reflect.Slice {
			for _, value := range el.([]interface{}) {
				if reflect.TypeOf(value).Kind() == reflect.Map {
					walk(value.(map[string]interface{}), passwordPhrase, dryRun)
				}
			}
		}
	}
}

func isEncrypted(val string) bool {
	if strings.HasPrefix(val, "$ANSIBLE_VAULT;1.1;AES256") {
		return true
	}
	return false
}
