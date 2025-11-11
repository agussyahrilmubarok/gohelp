package xcopy

import (
	"encoding/json"
)

// ByJSON performs a json copy from source to target using JSON serialization.
func ByJSON(source interface{}, target interface{}) {
	data, _ := json.Marshal(source)
	_ = json.Unmarshal(data, target)
}
