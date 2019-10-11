package gtool

import (
	"encoding/json"
	"fmt"
)

// return json encode string
func JsonEncode(i interface{}) string {
	b, err := json.Marshal(i)

	if err != nil {
		fmt.Println("util.getJsonStr.error", err)
		return ""
	}

	return string(b)
}
