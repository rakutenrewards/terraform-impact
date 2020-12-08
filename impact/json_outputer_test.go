package impact

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonOutputer(t *testing.T) {
	outputFile := "test_resources/json-test-output.json"
	outputer := JsonOutputer{outputFile}
	want := []string{"one", "two", "three"}

	err := outputer.Output(want)

	assert := assert.New(t)
	assert.Nil(err, "Error should be nil")

	bytes, _ := ioutil.ReadFile(outputFile)
	var data jsonOutputerData
	json.Unmarshal(bytes, &data)
	assert.ElementsMatch(data.States, want, "Json unserialized file should match content")
}
