package main

import (
	"errors"
	"fmt"

	"github.com/Jeffail/gabs/v2"
)

func customMappingFuncOBJSET(keyValuePairs ...interface{}) (result map[string]interface{}, err error) {
	g := gabs.New()

	result = make(map[string]interface{}, len(keyValuePairs)/2)
	// Iterate over keyValuePairs
	for j := 0; j < len(keyValuePairs)-1; j += 2 {
		key, ok := keyValuePairs[j].(string)
		if !ok {
			return nil, errors.New("key is not a string")
		}
		result[key] = keyValuePairs[j+1]
	}

	return
}
func main() {
	fmt.Println(customMappingFuncOBJSET("id", "123", "resourcePath.x", "devtests", "test", true, "x", 1))
}
