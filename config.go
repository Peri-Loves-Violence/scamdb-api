package scamdb

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/Peri-Loves-Violence/scamdb-api/types"
)

func ReadServers(filename string) ([]types.ServerEntry, error) {
	var servers []types.ServerEntry
	file, err := os.Open(filename)
	byteValue, _ := ioutil.ReadAll(file)
	defer file.Close()
	json.Unmarshal(byteValue, &servers)
	return servers, err
}

func WriteServers(data []types.ServerEntry, filename string) error {
	dataBytes, _ := json.Marshal(data)
	return ioutil.WriteFile(filename, dataBytes, 0644)
}
