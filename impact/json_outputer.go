package impact

import (
	"encoding/json"
	"io/ioutil"
)

type JsonOutputer struct {
	FilePath string
}

type jsonOutputerData struct {
	States []string `json:"states"`
}

func NewJsonOutputer(filePath string) JsonOutputer {
	return JsonOutputer{filePath}
}

func (out JsonOutputer) Output(results []string) error {
	data := jsonOutputerData{results}
	asJson, marshalErr := json.Marshal(data)
	if marshalErr != nil {
		return marshalErr
	}

	return ioutil.WriteFile(out.FilePath, asJson, 0644)
}
