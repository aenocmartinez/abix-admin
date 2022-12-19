package shared

import "encoding/json"

func ResponseJSON(body []byte) interface{} {
	var jsonParsed interface{}
	json.Unmarshal(body, &jsonParsed)
	return jsonParsed
}
