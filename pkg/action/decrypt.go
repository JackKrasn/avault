package action

import (
	"encoding/json"
	"fmt"
	"github.com/ghodss/yaml"
	vault "github.com/sosedoff/ansible-vault-go"
	"io/ioutil"
	"log"
	"reflect"
	"strings"
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

func (d *Decrypt) Run(encFileName string) (string, error) {
	decryptedFileName := encFileName + ".dec"
	fmt.Println("Starting decryption")
	d.cfg.Log("Performing decryption for file %s", encFileName)
	d.cfg.Log("Password phrase: %s", d.Password)
	yamlFile, err := ioutil.ReadFile(encFileName)
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
	walk(m, d.Password)
	fmt.Println("File was succesfully decrypted")
	// write encypted data to the yaml file
	data, err := yaml.Marshal(&m)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return "", err
	}
	err2 := ioutil.WriteFile(decryptedFileName, data, 0644)
	if err2 != nil {
		fmt.Printf("err: %v\n", err)
		return "", err2
	}
	return decryptedFileName, nil
}

func walk(data map[string]interface{}, passwordPhrase string) {
	for key, el := range data {
		if reflect.TypeOf(el).Kind() == reflect.Map {
			walk(el.(map[string]interface{}), passwordPhrase)
		}
		if reflect.TypeOf(el).Kind() == reflect.String {
			if isEncrypted(el.(string)) {
				// Decrypt secret data
				decryptedStr, err := vault.Decrypt(el.(string), passwordPhrase)
				if err != nil {
					log.Fatalf("Can'not decrypt key: %v\n", key)
				}
				data[key] = decryptedStr
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
