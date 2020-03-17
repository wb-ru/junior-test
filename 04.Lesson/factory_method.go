package factory_method

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"gopkg.in/yaml.v2"
	"reflect"
	"runtime"
)

// This type stores the author and th text of the song
type Song struct {
	Author string
	Text   string
}

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

// Suppose we wanna serialize song for some purpose
// We may use dozens of formats such .json, .xml, .yaml
// eventually it would be to implement all of those in sepatate methods
// instead, we attempt to create to a neat single factory method
func (p *Song) serialize(format string) string {
	var method func(interface{}) ([]byte, error)
	switch format {
	case "json":
		method = json.Marshal
	case "xml":
		method = xml.Marshal
	case "yaml":
		method = yaml.Marshal
	}
	product, err := method(*p)
	if err != nil {
		panic(fmt.Sprintf("Error converting with method %s", getFunctionName(method)))
	}
	return string(product)
}
