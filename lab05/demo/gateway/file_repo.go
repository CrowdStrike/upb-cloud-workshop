package gateway

import (
	"encoding/json"
	"os"

	log "github.com/sirupsen/logrus"
)

const fileName = "/tmp/dat"

type FileStorage struct {
}

func NewFileRepo() *FileStorage {
	return &FileStorage{}
}

func (f *FileStorage) GetContent(id string) (string, error) {
	return "", nil
}

func (f *FileStorage) WriteContent(id string, content string) error {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.WithError(err).Errorf("Failed to read file")
		return err
	}
	var jsonData map[string]string
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		log.WithError(err).Errorf("Failed to unmashal file")
		return err
	}
	jsonData[id] = content
	fileContent, err := json.Marshal(jsonData)
	if err != nil {
		log.WithError(err).Errorf("Failed to marshal file")
		return err
	}
	err = os.WriteFile(fileName, fileContent, 777)
	if err != nil {
		log.WithError(err).Errorf("Failed to write file")
		return err
	}
	return nil
}
