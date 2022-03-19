package scamdb

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadServers(filename string) ([]ServerEntry, error) {
	var servers []ServerEntry
	file, err := os.Open(filename)
	byteValue, _ := ioutil.ReadAll(file)
	defer file.Close()
	json.Unmarshal(byteValue, &servers)
	return servers, err
}

func WriteServers(data *interface{}, filename string) error {
	dataBytes, _ := json.Marshal(*data)
	return ioutil.WriteFile(filename, dataBytes, 0644)
}
